package data

import (
	"context"
	"fmt"
	"github.com/DenisKhanov/ResumeGame/internal/client/state"
	protodata "github.com/DenisKhanov/ResumeGame/pkg/resume_v1/data"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"strings"
)

// GRPCData is an interface that defines methods for get game data.
// Implementations of this interface should provide the actual functionality.
type GRPCData interface {
	GetGameInfo(ctx context.Context) (info string, err error)
	GetAboutOwner(ctx context.Context) (info string, err error)
	GetProjectList(ctx context.Context) (data []*protodata.Project, err error)
	GetSkills(ctx context.Context) (data []*protodata.Skill, err error)
	GetPreviousJobs(ctx context.Context) (data []*protodata.Experience, err error)
	GetContacts(ctx context.Context) (contacts *protodata.Contacts, err error)
}

// ServiceData is a struct that provides user-related functionalities.
// It contains a reference to a GRPCData and a ClientState.
type ServiceData struct {
	dataService GRPCData           // The data service implementation
	State       *state.ClientState // Client state management
}

// NewServiceDataService initializes a new ServiceData with the given user service and client state.
// It returns a pointer to the newly created ServiceData.
func NewServiceDataService(u GRPCData, state *state.ClientState) *ServiceData {
	return &ServiceData{
		dataService: u,
		State:       state,
	}
}

// GetGameInfo request info about game.
func (u *ServiceData) GetGameInfo(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	info, err := u.dataService.GetGameInfo(ctx)
	if err != nil {
		logrus.WithError(err).Error("game info")
		fmt.Println(red("It's empty here for now"))
		return
	}
	fmt.Println(green(info))
}

// GetAboutOwner request info about creator.
func (u *ServiceData) GetAboutOwner(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	info, err := u.dataService.GetAboutOwner(ctx)
	if err != nil {
		logrus.WithError(err).Error("owner info")
		fmt.Println(red("Owner don't found here  now"))
		return
	}
	fmt.Println(green(info))
}

// GetProjectList request info about creator.
func (u *ServiceData) GetProjectList(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	projects, err := u.dataService.GetProjectList(ctx)
	if err != nil {
		logrus.WithError(err).Error("projects list info")
		fmt.Println(red("Projects don't found here  now"))
		return
	}
	sb := strings.Builder{}
	sb.WriteString("Мной были реализованы следующие проекты:\n")
	for _, project := range projects {
		sb.WriteString(yellow(project.Name, "\n"))
		sb.WriteString(green(project.Description, "\n"))
		sb.WriteString("При его написании были задействованы следующие технологии и навыки:  " + blue(project.UsedSkills, "\n"))
		sb.WriteString(green("___________________________________________________\n"))
	}
	fmt.Println(sb.String())
}

// GetSkills request info about creator.
func (u *ServiceData) GetSkills(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	skills, err := u.dataService.GetSkills(ctx)
	if err != nil {
		logrus.WithError(err).Error("projects list info")
		fmt.Println(red("Projects don't found here  now"))
		return
	}
	sb := strings.Builder{}
	sb.WriteString("Мой список навыков:\n")
	for _, skill := range skills {
		sb.WriteString(yellow(skill.Name, "\n"))
		sb.WriteString(green(skill.Description, "\n"))
		sb.WriteString("Уровень владения - " + blue(skill.Level, "\n"))
		sb.WriteString(green("___________________________________________________\n"))
	}
	fmt.Println(sb.String())
}

// GetPreviousJobs request info about creator.
func (u *ServiceData) GetPreviousJobs(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	jobs, err := u.dataService.GetPreviousJobs(ctx)
	if err != nil {
		logrus.WithError(err).Error("projects list info")
		fmt.Println(red("Projects don't found here  now"))
		return
	}
	sb := strings.Builder{}
	sb.WriteString("Мой предыдущий опыт работы:\n")
	for _, experience := range jobs {
		sb.WriteString(yellow(experience.Organisation, "\n"))
		sb.WriteString("Должность: " + yellow(experience.Position, "\n"))
		sb.WriteString(blue("C "+experience.DateStart+" по "+experience.DateEnd, "\n"))
		sb.WriteString(green(experience.Responsibilities, "\n"))
		sb.WriteString(green("___________________________________________________\n"))
	}
	fmt.Println(sb.String())
}

// GetContacts request info about creator.
func (u *ServiceData) GetContacts(ctx context.Context) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	contacts, err := u.dataService.GetContacts(ctx)
	if err != nil {
		logrus.WithError(err).Error("projects list info")
		fmt.Println(red("Projects don't found here  now"))
		return
	}
	sb := strings.Builder{}
	sb.WriteString("Со мной можно связаться следующими способами:\n")
	sb.WriteString("Телефон: " + yellow(contacts.PhoneNumber, "\n"))
	sb.WriteString("Email: " + yellow(contacts.Email, "\n"))
	sb.WriteString("Telegram: " + green(contacts.Telegram, "\n"))
	sb.WriteString("Linkedin: " + blue(contacts.LinkedinLink, "\n"))
	sb.WriteString("Github repository: " + blue(contacts.Github, "\n\n"))

	fmt.Println(sb.String())
}
