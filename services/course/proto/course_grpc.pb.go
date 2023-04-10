// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/course.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoursesClient is the client API for Courses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursesClient interface {
	// For users
	Enroll(ctx context.Context, in *EnrollRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unenroll(ctx context.Context, in *UnenrollRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetActiveCoursesForUser(ctx context.Context, in *GetActiveCoursesForUserRequest, opts ...grpc.CallOption) (*GetActiveCoursesForUserResponse, error)
	// For teachers
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// General
	GetCourseInfo(ctx context.Context, in *GetCourseInfoRequest, opts ...grpc.CallOption) (*GetCourseInfoResponse, error)
	GetAllPublicCourses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllPublicCoursesResponse, error)
}

type coursesClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursesClient(cc grpc.ClientConnInterface) CoursesClient {
	return &coursesClient{cc}
}

func (c *coursesClient) Enroll(ctx context.Context, in *EnrollRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/Enroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) Unenroll(ctx context.Context, in *UnenrollRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/Unenroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) GetActiveCoursesForUser(ctx context.Context, in *GetActiveCoursesForUserRequest, opts ...grpc.CallOption) (*GetActiveCoursesForUserResponse, error) {
	out := new(GetActiveCoursesForUserResponse)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/GetActiveCoursesForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) GetCourseInfo(ctx context.Context, in *GetCourseInfoRequest, opts ...grpc.CallOption) (*GetCourseInfoResponse, error) {
	out := new(GetCourseInfoResponse)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/GetCourseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesClient) GetAllPublicCourses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllPublicCoursesResponse, error) {
	out := new(GetAllPublicCoursesResponse)
	err := c.cc.Invoke(ctx, "/scholarlabs.services.course.Courses/GetAllPublicCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursesServer is the server API for Courses service.
// All implementations must embed UnimplementedCoursesServer
// for forward compatibility
type CoursesServer interface {
	// For users
	Enroll(context.Context, *EnrollRequest) (*emptypb.Empty, error)
	Unenroll(context.Context, *UnenrollRequest) (*emptypb.Empty, error)
	GetActiveCoursesForUser(context.Context, *GetActiveCoursesForUserRequest) (*GetActiveCoursesForUserResponse, error)
	// For teachers
	CreateCourse(context.Context, *CreateCourseRequest) (*emptypb.Empty, error)
	// General
	GetCourseInfo(context.Context, *GetCourseInfoRequest) (*GetCourseInfoResponse, error)
	GetAllPublicCourses(context.Context, *emptypb.Empty) (*GetAllPublicCoursesResponse, error)
	mustEmbedUnimplementedCoursesServer()
}

// UnimplementedCoursesServer must be embedded to have forward compatible implementations.
type UnimplementedCoursesServer struct {
}

func (UnimplementedCoursesServer) Enroll(context.Context, *EnrollRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enroll not implemented")
}
func (UnimplementedCoursesServer) Unenroll(context.Context, *UnenrollRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unenroll not implemented")
}
func (UnimplementedCoursesServer) GetActiveCoursesForUser(context.Context, *GetActiveCoursesForUserRequest) (*GetActiveCoursesForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveCoursesForUser not implemented")
}
func (UnimplementedCoursesServer) CreateCourse(context.Context, *CreateCourseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCoursesServer) GetCourseInfo(context.Context, *GetCourseInfoRequest) (*GetCourseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseInfo not implemented")
}
func (UnimplementedCoursesServer) GetAllPublicCourses(context.Context, *emptypb.Empty) (*GetAllPublicCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicCourses not implemented")
}
func (UnimplementedCoursesServer) mustEmbedUnimplementedCoursesServer() {}

// UnsafeCoursesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursesServer will
// result in compilation errors.
type UnsafeCoursesServer interface {
	mustEmbedUnimplementedCoursesServer()
}

func RegisterCoursesServer(s grpc.ServiceRegistrar, srv CoursesServer) {
	s.RegisterService(&Courses_ServiceDesc, srv)
}

func _Courses_Enroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).Enroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/Enroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).Enroll(ctx, req.(*EnrollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courses_Unenroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnenrollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).Unenroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/Unenroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).Unenroll(ctx, req.(*UnenrollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courses_GetActiveCoursesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveCoursesForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).GetActiveCoursesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/GetActiveCoursesForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).GetActiveCoursesForUser(ctx, req.(*GetActiveCoursesForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courses_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courses_GetCourseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).GetCourseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/GetCourseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).GetCourseInfo(ctx, req.(*GetCourseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courses_GetAllPublicCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServer).GetAllPublicCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scholarlabs.services.course.Courses/GetAllPublicCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServer).GetAllPublicCourses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Courses_ServiceDesc is the grpc.ServiceDesc for Courses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scholarlabs.services.course.Courses",
	HandlerType: (*CoursesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enroll",
			Handler:    _Courses_Enroll_Handler,
		},
		{
			MethodName: "Unenroll",
			Handler:    _Courses_Unenroll_Handler,
		},
		{
			MethodName: "GetActiveCoursesForUser",
			Handler:    _Courses_GetActiveCoursesForUser_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _Courses_CreateCourse_Handler,
		},
		{
			MethodName: "GetCourseInfo",
			Handler:    _Courses_GetCourseInfo_Handler,
		},
		{
			MethodName: "GetAllPublicCourses",
			Handler:    _Courses_GetAllPublicCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/course.proto",
}
