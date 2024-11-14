package data

import (
	"context"
	"fmt"
	models2 "github.com/DenisKhanov/ResumeGame/internal/server/models"
	repodata "github.com/DenisKhanov/ResumeGame/internal/server/repositories/data"
	protodata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

var _ RepoData = (*repodata.RepositoryData)(nil)

// ServiceData ...
type ServiceData struct {
	ownerID    int
	repository RepoData
	dbPool     *pgxpool.Pool
}

// NewServiceData .....
func NewServiceData(repository RepoData, dbPool *pgxpool.Pool) *ServiceData {
	return &ServiceData{
		ownerID:    1,
		repository: repository,
		dbPool:     dbPool,
	}
}

type RepoData interface {
	GetAboutOwner(ctx context.Context, ownerId int) (owner models2.Owner, err error)
	GetProjectList(ctx context.Context, ownerId int) (data []models2.Project, err error)
	GetSkills(ctx context.Context, ownerId int) (data []models2.Skill, err error)
	GetPreviousJobs(ctx context.Context, ownerId int) (data []models2.Experience, err error)
	GetContacts(ctx context.Context, ownerId int) (contacts models2.Contacts, err error)
}

var (
	buildVersion string
	buildDate    string
	description  string
	buildCommit  string
)

// getValueOrDefault returns the value, and if it is empty,it returns the default value.
func getValueOrDefault(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// PrintProjectInfo print info (version,date,commit) about build.
func printProjectInfo() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Build version: %s\n", getValueOrDefault(buildVersion, "1.0")))
	sb.WriteString(fmt.Sprintf("Build date: %s\n", getValueOrDefault(buildDate, "1.2024")))
	sb.WriteString(fmt.Sprintf("About game: %s\n", getValueOrDefault(description, "Эта игра создана для интереса разработчика "+
		"и в игровой форме помогает вам познакомиться с его резюме. Получайте удовольствие :)")))
	sb.WriteString(fmt.Sprintf("Build commit: %s\n", getValueOrDefault(buildCommit, "Developer edition")))

	return sb.String()
}

func (s *ServiceData) GetGameInfo(_ context.Context) (info string, err error) {
	info = printProjectInfo()
	if info == "" {
		return "", models2.ErrGameInfo
	}
	return info, nil
}

func (s *ServiceData) GetAboutOwner(ctx context.Context) (info string, err error) {
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	owner, err := s.repository.GetAboutOwner(ctx, s.ownerID)
	if err != nil {
		return "", err
	}
	sb := strings.Builder{}
	sb.WriteString(yellow("Меня зовут " + owner.FirstName + " " + owner.LastName + "\n"))
	sb.WriteString(blue("Мне " + owner.Age + " года\n"))
	sb.WriteString(owner.About + "\n")
	return sb.String(), nil
}
func (s *ServiceData) GetProjectList(ctx context.Context) (data []*protodata.Project, err error) {
	ownerProjects, err := s.repository.GetProjectList(ctx, s.ownerID)
	if err != nil {
		return nil, err
	}
	var projectProto *protodata.Project
	for _, project := range ownerProjects {
		projectProto = &protodata.Project{
			Name:        project.Name,
			Description: project.Description,
			UsedSkills:  project.UsedSkills,
		}
		data = append(data, projectProto)
	}
	return data, nil
}
func (s *ServiceData) GetSkills(ctx context.Context) (data []*protodata.Skill, err error) {
	ownerSkills, err := s.repository.GetSkills(ctx, s.ownerID)
	if err != nil {
		return nil, err
	}
	var skillProto *protodata.Skill
	for _, skill := range ownerSkills {
		skillProto = &protodata.Skill{
			Name:        skill.Name,
			Description: skill.Description,
			Level:       skill.Level,
		}
		data = append(data, skillProto)
	}
	return data, nil
}

func (s *ServiceData) GetPreviousJobs(ctx context.Context) (data []*protodata.Experience, err error) {
	ownerExperience, err := s.repository.GetPreviousJobs(ctx, s.ownerID)
	if err != nil {
		return nil, err
	}
	var experienceProto *protodata.Experience
	for _, experience := range ownerExperience {
		experienceProto = &protodata.Experience{
			Organisation:     experience.Organisation,
			Position:         experience.Position,
			Responsibilities: experience.Responsibilities,
			DateStart:        experience.DateStart.Format("2006-01"),
			DateEnd:          experience.DateEnd.Format("2006-01"),
		}
		data = append(data, experienceProto)
	}
	return data, nil
}

func (s *ServiceData) GetContacts(ctx context.Context) (contacts *protodata.Contacts, err error) {
	ownerContacts, err := s.repository.GetContacts(ctx, s.ownerID)
	if err != nil {
		return nil, err
	}
	contacts = &protodata.Contacts{
		PhoneNumber:  ownerContacts.PhoneNumber,
		Email:        ownerContacts.Email,
		Telegram:     ownerContacts.Telegram,
		LinkedinLink: ownerContacts.Linkedin,
		Github:       ownerContacts.GitHub,
	}

	return contacts, nil
}
