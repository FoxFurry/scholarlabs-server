package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *ScholarLabs) GetPage(ctx *gin.Context) {
	var (
		pageID     = ctx.Param("pageID")
		courseUUID = ctx.Param("courseUUID")
	)

	page, err := s.courseService.GetPage(ctx, &proto.GetPageRequest{
		PageID:     pageID,
		CourseUUID: courseUUID,
	})
	if err != nil {
		s.lg.WithError(err).WithField("course uuid", courseUUID).WithField("page id", pageID).Error("failed to get page")
		httperr.Handle(ctx, httperr.New("internal error", 500))
		return
	}

	response := models.Page{}
	response.PageMetadata.ID = page.Metadata.ID
	response.PageMetadata.Title = page.Metadata.Title
	response.PageMetadata.Type = page.Metadata.Type.String()

	response.Data, err = protoToModelPageData(page.Metadata.Type.String(), page.Data)
	if err != nil {
		s.lg.WithError(err).WithField("course uuid", courseUUID).WithField("page id", pageID).Error("failed to convert page data")
		httperr.Handle(ctx, httperr.New("internal error", 500))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *ScholarLabs) CreatePage(ctx *gin.Context) {
	var (
		page models.Page
		err  error
	)

	courseUUID := ctx.Param("courseUUID")

	if err := ctx.BindJSON(&page); err != nil {
		s.lg.WithError(err).Error("failed to bind request")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	var pageData []byte
	switch page.Type {
	case "ASSIGNMENT":
		pageData, err = json.Marshal(page.Data.Assignment)
		if err != nil {
			s.lg.WithError(err).Error("failed to marshal assignment")
			httperr.Handle(ctx, httperr.New("internal error", 500))
			return
		}
	case "LESSON":
		pageData, err = json.Marshal(page.Data.Lesson)
		if err != nil {
			s.lg.WithError(err).Error("failed to marshal lesson")
			httperr.Handle(ctx, httperr.New("internal error", 500))
			return
		}
	default:
		s.lg.WithField("page type", page.Type).Error("unknown page type")
	}

	pageType, err := stringToPageType(page.Type)
	if err != nil {
		s.lg.WithError(err).WithField("page type", page.Type).Error("failed to convert page type")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	_, err = s.courseService.CreatePage(ctx, &proto.CreatePageRequest{
		CourseUUID: courseUUID,
		Title:      page.Title,
		Type:       pageType,
		Data:       pageData,
	})
	if err != nil {
		s.lg.WithError(err).WithField("req", page).Error("failed to create page")
		httperr.Handle(ctx, httperr.New("internal error", 500))
		return
	}

	ctx.Status(http.StatusCreated)
}

func stringToPageType(pageType string) (proto.PageType, error) {
	switch pageType {
	case "ASSIGNMENT":
		return proto.PageType_ASSIGNMENT, nil
	case "LESSON":
		return proto.PageType_LESSON, nil
	default:
		return proto.PageType_Unknown, fmt.Errorf("unknown page type: %s", pageType)
	}
}

func protoToModelPageData(pageType string, pageData []byte) (models.PageData, error) {
	var modelPage models.PageData

	switch pageType {
	case "ASSIGNMENT":
		if err := json.Unmarshal(pageData, &modelPage.Assignment); err != nil {
			return modelPage, fmt.Errorf("failed to unmarshal assignment page data: %w", err)
		}
		modelPage.Assignment.UUID = uuid.New().String()
	case "LESSON":
		if err := json.Unmarshal(pageData, &modelPage.Lesson); err != nil {
			return modelPage, fmt.Errorf("failed to unmarshal lesson page data: %w", err)
		}
	default:
		return modelPage, fmt.Errorf("unknown page type: %s", pageType)
	}

	return modelPage, nil
}
