package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabs) Enroll(ctx context.Context, req *proto.EnrollRequest) (*emptypb.Empty, error) {
	err := p.service.Enroll(ctx, req.UserUUID, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to enroll")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) Unenroll(ctx context.Context, req *proto.UnenrollRequest) (*emptypb.Empty, error) {
	err := p.service.Unenroll(ctx, req.UserUUID, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to unenroll")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
