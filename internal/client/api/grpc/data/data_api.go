package data

import (
	"context"
	protodata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
)

// GRPCDataClient is a client for interacting with the user service via gRPC.
// It holds a reference to the user service client interface.
type GRPCDataClient struct {
	pbDataClient protodata.ResumeDataV1Client
}

// NewDataPBClient initializes a new GRPCDataClient with the provided user service client.
// It returns a pointer to the newly created GRPCDataClient.
func NewDataPBClient(pbDataClient protodata.ResumeDataV1Client) *GRPCDataClient {
	return &GRPCDataClient{
		pbDataClient: pbDataClient,
	}
}

// GetGameInfo .
func (u *GRPCDataClient) GetGameInfo(ctx context.Context) (string, error) {
	req := &protodata.GetGameInfoRequest{}

	res, err := u.pbDataClient.GetGameInfo(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GameInfo, nil
}

// GetAboutOwner .
func (u *GRPCDataClient) GetAboutOwner(ctx context.Context) (string, error) {
	req := &protodata.GetAboutOwnerRequest{}

	res, err := u.pbDataClient.GetAboutOwner(ctx, req)
	if err != nil {
		return "", err
	}

	return res.OwnerInfo, nil
}

// GetProjectList  .
func (u *GRPCDataClient) GetProjectList(ctx context.Context) ([]*protodata.Project, error) {
	req := &protodata.GetProjectListRequest{}

	res, err := u.pbDataClient.GetProjectList(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.AllProjects, nil
}

// GetSkills  .
func (u *GRPCDataClient) GetSkills(ctx context.Context) ([]*protodata.Skill, error) {
	req := &protodata.GetSkillsRequest{}

	res, err := u.pbDataClient.GetSkills(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.AllSkills, nil
}

// GetPreviousJobs  .
func (u *GRPCDataClient) GetPreviousJobs(ctx context.Context) ([]*protodata.Experience, error) {
	req := &protodata.GetPreviousJobsRequest{}

	res, err := u.pbDataClient.GetPreviousJobs(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.AllExperience, nil
}

// GetContacts  .
func (u *GRPCDataClient) GetContacts(ctx context.Context) (*protodata.Contacts, error) {
	req := &protodata.GetContactsRequest{}

	res, err := u.pbDataClient.GetContacts(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.Contacts, nil
}
