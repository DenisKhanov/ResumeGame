// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: resumeDataV1.proto

package resumeDataV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ResumeDataV1_GetGameInfo_FullMethodName     = "/resume_data_v1.ResumeDataV1/GetGameInfo"
	ResumeDataV1_GetAboutOwner_FullMethodName   = "/resume_data_v1.ResumeDataV1/GetAboutOwner"
	ResumeDataV1_GetProjectList_FullMethodName  = "/resume_data_v1.ResumeDataV1/GetProjectList"
	ResumeDataV1_GetSkills_FullMethodName       = "/resume_data_v1.ResumeDataV1/GetSkills"
	ResumeDataV1_GetPreviousJobs_FullMethodName = "/resume_data_v1.ResumeDataV1/GetPreviousJobs"
	ResumeDataV1_GetContacts_FullMethodName     = "/resume_data_v1.ResumeDataV1/GetContacts"
)

// ResumeDataV1Client is the client API for ResumeDataV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeDataV1Client interface {
	GetGameInfo(ctx context.Context, in *GetGameInfoRequest, opts ...grpc.CallOption) (*GetGameInfoResponse, error)
	GetAboutOwner(ctx context.Context, in *GetAboutOwnerRequest, opts ...grpc.CallOption) (*GetAboutOwnerResponse, error)
	GetProjectList(ctx context.Context, in *GetProjectListRequest, opts ...grpc.CallOption) (*GetProjectListResponse, error)
	GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*GetSkillsResponse, error)
	GetPreviousJobs(ctx context.Context, in *GetPreviousJobsRequest, opts ...grpc.CallOption) (*GetPreviousJobsResponse, error)
	GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error)
}

type resumeDataV1Client struct {
	cc grpc.ClientConnInterface
}

func NewResumeDataV1Client(cc grpc.ClientConnInterface) ResumeDataV1Client {
	return &resumeDataV1Client{cc}
}

func (c *resumeDataV1Client) GetGameInfo(ctx context.Context, in *GetGameInfoRequest, opts ...grpc.CallOption) (*GetGameInfoResponse, error) {
	out := new(GetGameInfoResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetGameInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeDataV1Client) GetAboutOwner(ctx context.Context, in *GetAboutOwnerRequest, opts ...grpc.CallOption) (*GetAboutOwnerResponse, error) {
	out := new(GetAboutOwnerResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetAboutOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeDataV1Client) GetProjectList(ctx context.Context, in *GetProjectListRequest, opts ...grpc.CallOption) (*GetProjectListResponse, error) {
	out := new(GetProjectListResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetProjectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeDataV1Client) GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*GetSkillsResponse, error) {
	out := new(GetSkillsResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetSkills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeDataV1Client) GetPreviousJobs(ctx context.Context, in *GetPreviousJobsRequest, opts ...grpc.CallOption) (*GetPreviousJobsResponse, error) {
	out := new(GetPreviousJobsResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetPreviousJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeDataV1Client) GetContacts(ctx context.Context, in *GetContactsRequest, opts ...grpc.CallOption) (*GetContactsResponse, error) {
	out := new(GetContactsResponse)
	err := c.cc.Invoke(ctx, ResumeDataV1_GetContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeDataV1Server is the server API for ResumeDataV1 service.
// All implementations must embed UnimplementedResumeDataV1Server
// for forward compatibility
type ResumeDataV1Server interface {
	GetGameInfo(context.Context, *GetGameInfoRequest) (*GetGameInfoResponse, error)
	GetAboutOwner(context.Context, *GetAboutOwnerRequest) (*GetAboutOwnerResponse, error)
	GetProjectList(context.Context, *GetProjectListRequest) (*GetProjectListResponse, error)
	GetSkills(context.Context, *GetSkillsRequest) (*GetSkillsResponse, error)
	GetPreviousJobs(context.Context, *GetPreviousJobsRequest) (*GetPreviousJobsResponse, error)
	GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error)
	mustEmbedUnimplementedResumeDataV1Server()
}

// UnimplementedResumeDataV1Server must be embedded to have forward compatible implementations.
type UnimplementedResumeDataV1Server struct {
}

func (UnimplementedResumeDataV1Server) GetGameInfo(context.Context, *GetGameInfoRequest) (*GetGameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameInfo not implemented")
}
func (UnimplementedResumeDataV1Server) GetAboutOwner(context.Context, *GetAboutOwnerRequest) (*GetAboutOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAboutOwner not implemented")
}
func (UnimplementedResumeDataV1Server) GetProjectList(context.Context, *GetProjectListRequest) (*GetProjectListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectList not implemented")
}
func (UnimplementedResumeDataV1Server) GetSkills(context.Context, *GetSkillsRequest) (*GetSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkills not implemented")
}
func (UnimplementedResumeDataV1Server) GetPreviousJobs(context.Context, *GetPreviousJobsRequest) (*GetPreviousJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviousJobs not implemented")
}
func (UnimplementedResumeDataV1Server) GetContacts(context.Context, *GetContactsRequest) (*GetContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContacts not implemented")
}
func (UnimplementedResumeDataV1Server) mustEmbedUnimplementedResumeDataV1Server() {}

// UnsafeResumeDataV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeDataV1Server will
// result in compilation errors.
type UnsafeResumeDataV1Server interface {
	mustEmbedUnimplementedResumeDataV1Server()
}

func RegisterResumeDataV1Server(s grpc.ServiceRegistrar, srv ResumeDataV1Server) {
	s.RegisterService(&ResumeDataV1_ServiceDesc, srv)
}

func _ResumeDataV1_GetGameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetGameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetGameInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetGameInfo(ctx, req.(*GetGameInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeDataV1_GetAboutOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAboutOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetAboutOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetAboutOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetAboutOwner(ctx, req.(*GetAboutOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeDataV1_GetProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetProjectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetProjectList(ctx, req.(*GetProjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeDataV1_GetSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetSkills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetSkills(ctx, req.(*GetSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeDataV1_GetPreviousJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviousJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetPreviousJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetPreviousJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetPreviousJobs(ctx, req.(*GetPreviousJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeDataV1_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeDataV1Server).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeDataV1_GetContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeDataV1Server).GetContacts(ctx, req.(*GetContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResumeDataV1_ServiceDesc is the grpc.ServiceDesc for ResumeDataV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResumeDataV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resume_data_v1.ResumeDataV1",
	HandlerType: (*ResumeDataV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGameInfo",
			Handler:    _ResumeDataV1_GetGameInfo_Handler,
		},
		{
			MethodName: "GetAboutOwner",
			Handler:    _ResumeDataV1_GetAboutOwner_Handler,
		},
		{
			MethodName: "GetProjectList",
			Handler:    _ResumeDataV1_GetProjectList_Handler,
		},
		{
			MethodName: "GetSkills",
			Handler:    _ResumeDataV1_GetSkills_Handler,
		},
		{
			MethodName: "GetPreviousJobs",
			Handler:    _ResumeDataV1_GetPreviousJobs_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _ResumeDataV1_GetContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resumeDataV1.proto",
}