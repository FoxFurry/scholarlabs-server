package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabs) GetPage(ctx context.Context, req *proto.GetPageRequest) (*proto.GetPageResponse, error) {
	page, err := p.service.GetPageByID(ctx, req.CourseUUID, req.PageID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get page")
		return nil, err
	}

	return &proto.GetPageResponse{
		Data: page.Data,
		Metadata: &proto.PageMetadata{
			ID:    page.PageID,
			Title: page.Title,
			Type:  storeToProtoPageType(page.Type),
		},
	}, nil
}

func (p *ScholarLabs) CreatePage(ctx context.Context, req *proto.CreatePageRequest) (*emptypb.Empty, error) {
	if err := p.service.CreatePage(ctx, store.Page{
		CourseUUID: req.CourseUUID,
		Data:       req.Data,
		PageIdentifier: store.PageIdentifier{
			Title: req.Title,
			Type:  protoToStorePageType(req.Type),
		},
	}); err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to create page")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) DeletePage(ctx context.Context, req *proto.DeletePageRequest) (*emptypb.Empty, error) {
	if err := p.service.DeletePage(ctx, req.CourseUUID, req.PageID); err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to delete page")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) GetCourseToC(ctx context.Context, req *proto.GetCourseToCRequest) (*proto.GetCourseToCResponse, error) {
	pages, err := p.service.GetCourseToc(ctx, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get course toc")
		return nil, err
	}

	var result []*proto.PageMetadata
	for _, page := range pages {
		result = append(result, &proto.PageMetadata{
			ID:    page.PageID,
			Title: page.Title,
			Type:  storeToProtoPageType(page.Type),
		})
	}

	return &proto.GetCourseToCResponse{
		Pages: result,
	}, nil
}

func protoToStorePageType(page proto.PageType) store.PageType {
	switch page {
	case proto.PageType_ASSIGNMENT:
		return store.PageTypeAssignment
	case proto.PageType_LESSON:
		return store.PageTypeLesson
	default: // TODO: handle unknown
		return store.PageTypeLesson
	}
}

func storeToProtoPageType(page store.PageType) proto.PageType {
	switch page {
	case store.PageTypeAssignment:
		return proto.PageType_ASSIGNMENT
	case store.PageTypeLesson:
		return proto.PageType_LESSON
	default:
		return proto.PageType_Unknown
	}
}
