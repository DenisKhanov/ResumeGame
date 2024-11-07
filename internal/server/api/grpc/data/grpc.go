package data

import (
	"context"
	servicedata "github.com/DenisKhanov/ResumeGame/internal/server/services/data"
	protodata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// checking interface compliance at the compiler level
var _ ServData = (*servicedata.ServiceData)(nil)

type GRPCData struct {
	protodata.UnimplementedResumeDataV1Server
	service ServData
}

func NewGRPCData(service ServData) *GRPCData {
	return &GRPCData{
		service: service,
	}
}

type ServData interface {
	GetGameInfo(ctx context.Context) (info string, err error)
	GetAboutOwner(ctx context.Context) (info string, err error)
	GetProjectList(ctx context.Context) (data []*protodata.Project, err error)
	GetSkills(ctx context.Context) (data []*protodata.Skill, err error)
	GetPreviousJobs(ctx context.Context) (data []*protodata.Experience, err error)
	GetContacts(ctx context.Context) (contacts *protodata.Contacts, err error)
}

func (g *GRPCData) GetGameInfo(ctx context.Context, _ *protodata.GetGameInfoRequest) (response *protodata.GetGameInfoResponse, err error) {
	info, err := g.service.GetGameInfo(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get game info")
		return nil, status.Error(codes.Internal, "don't get game info")
	}
	if info == "" {
		return nil, status.Error(codes.Internal, "don't have game info")
	}
	response.GameInfo = info
	return response, nil
}

func (g *GRPCData) GetAboutOwner(ctx context.Context, _ *protodata.GetAboutOwnerRequest) (response *protodata.GetAboutOwnerResponse, err error) {
	info, err := g.service.GetAboutOwner(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get owner info")
		return nil, status.Error(codes.Internal, "don't get owner info")
	}
	if info == "" {
		return nil, status.Error(codes.Internal, "don't have info about this project")
	}
	response.OwnerInfo = info
	return response, nil
}
func (g *GRPCData) GetProjectList(ctx context.Context, _ *protodata.GetProjectListRequest) (response *protodata.GetProjectListResponse, err error) {
	projects, err := g.service.GetProjectList(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get projects")
		return nil, status.Error(codes.Internal, "don't get any projects")
	}
	if len(projects) == 0 {
		return nil, status.Error(codes.Internal, "don't have any projects")
	}
	response.AllProjects = projects
	return response, nil
}

func (g *GRPCData) GetSkills(ctx context.Context, _ *protodata.GetSkillsRequest) (response *protodata.GetSkillsResponse, err error) {
	skills, err := g.service.GetSkills(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get skills")
		return nil, status.Error(codes.Internal, "don't get any skills")
	}
	if len(skills) == 0 {
		return nil, status.Error(codes.Internal, "don't have any skills")
	}
	response.AllSkills = skills
	return response, err
}

func (g *GRPCData) GetPreviousJobs(ctx context.Context, _ *protodata.GetPreviousJobsRequest) (response *protodata.GetPreviousJobsResponse, err error) {
	jobs, err := g.service.GetPreviousJobs(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get jobs history")
		return nil, status.Error(codes.Internal, "don't get any jobs history")
	}
	if len(jobs) == 0 {
		return nil, status.Error(codes.Internal, "don't have any jobs history")
	}
	response.AllExperience = jobs
	return response, err
}

func (g *GRPCData) GetContacts(ctx context.Context, _ *protodata.GetContactsRequest) (response *protodata.GetContactsResponse, err error) {
	contacts, err := g.service.GetContacts(ctx)
	if err != nil {
		logrus.WithError(err).Error("don't get owner contacts")
		return nil, status.Error(codes.Internal, "don't get owner contacts")
	}
	response.Contacts = contacts
	return response, nil
}
