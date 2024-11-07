package data

import (
	"context"
	"github.com/DenisKhanov/ResumeGame/internal/server/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RepositoryData ...
type RepositoryData struct {
	dbPool *pgxpool.Pool //opened in main func dbPool pool connections
}

func NewPostgresData(dbPool *pgxpool.Pool) *RepositoryData {
	return &RepositoryData{
		dbPool: dbPool,
	}
}

func (r *RepositoryData) GetAboutOwner(ctx context.Context, ownerID int) (owner models.Owner, err error) {
	const sqlQuery = `
		SELECT 
    			first_name, 
   				last_name, 
    			birth_date,
    			about_me 
		FROM
				owners
		WHERE 
		    	id = $1
`
	if err = r.dbPool.QueryRow(ctx, sqlQuery, ownerID).Scan(
		&owner.FirstName,
		&owner.LastName,
		&owner.BirthDate,
		&owner.About,
	); err != nil {
		return models.Owner{}, err
	}
	return owner, nil
}

func (r *RepositoryData) GetProjectList(ctx context.Context, ownerID int) (data []models.Project, err error) {
	const sqlQuery = `
		SELECT 
    			p.name AS project_name,
    			p.description AS project_description,
    			GROUP_CONCAT(s.name ORDER BY s.name SEPARATOR ', ') AS used_skills
		FROM 
    			projects p
		JOIN 
   				owner_projects op ON p.id = op.project_id
		JOIN 
    			project_skills ps ON p.id = ps.project_id
		JOIN 
    			skills s ON ps.skill_id = s.id
		WHERE 
    			op.owner_id = $1 
		ORDER BY 
    			p.name;`
	rows, err := r.dbPool.Query(ctx, sqlQuery, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var project models.Project
	for rows.Next() {
		if err = rows.Scan(
			&project.Name,
			&project.Description,
			&project.UsedSkills,
		); err != nil {
			return nil, err
		}
		data = append(data, project)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
func (r *RepositoryData) GetSkills(ctx context.Context, ownerID int) (data []models.Skill, err error) {
	const sqlQuery = `
		SELECT 
    			s.name,
    			s.description,
    			s.level
		FROM 
    			skills s
		JOIN 
   				owner_skills os ON s.id = os.skill_id
		WHERE 
    			os.owner_id = $1 
		ORDER BY 
    			s.name;`
	rows, err := r.dbPool.Query(ctx, sqlQuery, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var skill models.Skill
	for rows.Next() {
		if err = rows.Scan(
			&skill.Name,
			&skill.Description,
			&skill.Level,
		); err != nil {
			return nil, err
		}
		data = append(data, skill)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
func (r *RepositoryData) GetPreviousJobs(ctx context.Context, ownerID int) (data []models.Experience, err error) {
	const sqlQuery = `
		SELECT 
    			e.organisation,
    			e.responsibilities,
    			e.date_start,
    			e.date_end
		FROM 
    			experience e
		JOIN 
   				owner_experience oe ON e.id = oe.experience_id
		WHERE 
    			oe.owner_id = $1 
		ORDER BY 
    			e.date_end DESC;`
	rows, err := r.dbPool.Query(ctx, sqlQuery, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var experience models.Experience
	for rows.Next() {
		if err = rows.Scan(
			&experience.Organisation,
			&experience.Responsibilities,
			&experience.DateStart,
			&experience.DateEnd,
		); err != nil {
			return nil, err
		}
		data = append(data, experience)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *RepositoryData) GetContacts(ctx context.Context, ownerID int) (contacts models.Contacts, err error) {
	const sqlQuery = `
		SELECT 
    			c.phone_number, 
   				c.email, 
    			c.telegram,
    			c.linkedin,
				c.github
		FROM
				contacts c
		JOIN owner_contacts oc ON c.id = oc.contact_id
		WHERE 
		    	oc.owner_id = $1
`
	if err = r.dbPool.QueryRow(ctx, sqlQuery, ownerID).Scan(
		&contacts.PhoneNumber,
		&contacts.Email,
		&contacts.Telegram,
		&contacts.Linkedin,
		&contacts.GitHub,
	); err != nil {
		return models.Contacts{}, err
	}
	return contacts, nil
}
