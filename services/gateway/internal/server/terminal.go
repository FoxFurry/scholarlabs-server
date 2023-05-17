package server

import (
	"context"
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const TERMINAL_INIT = "TERMINAL_INIT"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *ScholarLabs) BidirectionalTerminal(ctx *gin.Context) {
	EnvironmentUUID := ctx.Query("environment_uuid")
	if EnvironmentUUID == "" {
		s.lg.Error("environment uuid is empty")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	terminal, err := s.environmentService.BidirectionalTerminal(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get terminal")
		httperr.Handle(ctx, httperr.New("something went wrong", 500))
		return
	}

	webS, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		s.lg.WithError(err).Error("failed to upgrade connection to web socket")
		httperr.Handle(ctx, httperr.New("something went wrong", 500))
		return
	}

	defer webS.Close()

	if err := terminal.Send(&proto.BidirectionalTerminalRequest{
		EnvironmentUUID: EnvironmentUUID,
		Command:         TERMINAL_INIT,
	}); err != nil {
		s.lg.WithError(err).Error("failed to initialize terminal")
		httperr.Handle(ctx, httperr.New("something went wrong", 500))
		return
	}

	var (
		terminalResponses = make(chan []byte, 25)
		wsMessages        = make(chan []byte, 25)
	)

	go s.listenTerminal(ctx, terminal, terminalResponses)
	go s.listenWebSocket(ctx, webS, wsMessages)

	for {
		select {
		case <-ctx.Done():
			s.lg.Info("context closed. Stopping listening terminal")
			return
		case msg := <-terminalResponses:
			if err := webS.WriteMessage(websocket.TextMessage, msg); err != nil {
				s.lg.WithError(err).Error("failed to write message to web socket")
				return
			}
		case msg := <-wsMessages:
			if err := terminal.Send(&proto.BidirectionalTerminalRequest{Command: string(msg)}); err != nil {
				s.lg.WithError(err).Error("failed to send message to terminal")
				return
			}
		}
	}
}

func (s *ScholarLabs) listenTerminal(ctx context.Context, terminal proto.Environment_BidirectionalTerminalClient, resultChan chan<- []byte) {
	for {
		select {
		case <-ctx.Done():
			s.lg.Info("context closed. Stopping listening terminal")
			return
		default:
		}

		message, err := terminal.Recv()
		if err != nil {
			s.lg.WithError(err).Error("failed to receive message from terminal")
			return
		}

		resultChan <- []byte(message.GetCommand())
	}
}

func (s *ScholarLabs) listenWebSocket(ctx context.Context, conn *websocket.Conn, resultChan chan<- []byte) {
	for {
		select {
		case <-ctx.Done():
			s.lg.Info("context closed. Stopping listening terminal")
			return
		default:
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			s.lg.WithError(err).Error("failed to receive message from web socket")
			return
		}

		resultChan <- message
	}
}
