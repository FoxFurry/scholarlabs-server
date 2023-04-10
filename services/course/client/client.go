package client

import (
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCourseClient(address string) (proto.CoursesClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return proto.NewCoursesClient(conn), nil
}
