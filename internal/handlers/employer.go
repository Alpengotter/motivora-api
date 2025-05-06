package handlers

import (
	"database/sql"
	"fmt"
	"motivora-backend/internal/db"
	"motivora-backend/internal/models"
)

func GetUsers(offset, limit int) ([]models.User, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name, job_title, company_id, email, lemons, diamonds, user_role, is_active FROM employers OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		var jobTitle sql.NullString
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &jobTitle, &u.CompanyID, &u.Email, &u.Lemons, &u.Diamonds, &u.UserRole, &u.IsActive)
		if err != nil {
			return nil, err
		}

		if jobTitle.Valid {
			u.JobTitle = &jobTitle.String
		} else {
			u.JobTitle = nil
		}

		users = append(users, u)
	}

	return users, nil
}

func CreateUser(user models.User) (int, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var newID int
	err = db.QueryRow(`
        INSERT INTO employers (
            first_name, last_name, job_title, company_id, email, lemons, diamonds, user_role, is_active
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id`,
		user.FirstName,
		user.LastName,
		user.JobTitle,
		user.CompanyID,
		user.Email,
		user.Lemons,
		user.Diamonds,
		user.UserRole,
		user.IsActive,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func GetUserByID(id int) (*models.User, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var u models.User
	var jobTitle sql.NullString

	row := db.QueryRow("SELECT id, first_name, last_name, job_title, company_id, email, lemons, diamonds, user_role, is_active FROM employers WHERE id = $1", id)

	err = row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&jobTitle,
		&u.CompanyID,
		&u.Email,
		&u.Lemons,
		&u.Diamonds,
		&u.UserRole,
		&u.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if jobTitle.Valid {
		u.JobTitle = &jobTitle.String
	} else {
		u.JobTitle = nil
	}

	return &u, nil
}

func UpdateUser(id int, updateReq models.UpdateUserRequest) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE employers SET "
	args := []interface{}{}
	argNum := 1

	if updateReq.FirstName != nil {
		query += fmt.Sprintf("first_name = $%d, ", argNum)
		args = append(args, *updateReq.FirstName)
		argNum++
	}
	if updateReq.LastName != nil {
		query += fmt.Sprintf("last_name = $%d, ", argNum)
		args = append(args, *updateReq.LastName)
		argNum++
	}
	if updateReq.JobTitle != nil {
		query += fmt.Sprintf("job_title = $%d, ", argNum)
		args = append(args, *updateReq.JobTitle)
		argNum++
	} else if updateReq.JobTitle == nil {
		query += "job_title = NULL, "
	}
	if updateReq.CompanyID != nil {
		query += fmt.Sprintf("company_id = $%d, ", argNum)
		args = append(args, *updateReq.CompanyID)
		argNum++
	}
	if updateReq.Email != nil {
		query += fmt.Sprintf("email = $%d, ", argNum)
		args = append(args, *updateReq.Email)
		argNum++
	}
	if updateReq.Lemons != nil {
		query += fmt.Sprintf("lemons = $%d, ", argNum)
		args = append(args, *updateReq.Lemons)
		argNum++
	}
	if updateReq.Diamonds != nil {
		query += fmt.Sprintf("diamonds = $%d, ", argNum)
		args = append(args, *updateReq.Diamonds)
		argNum++
	}
	if updateReq.UserRole != nil {
		query += fmt.Sprintf("user_role = $%d, ", argNum)
		args = append(args, *updateReq.UserRole)
		argNum++
	}
	if updateReq.IsActive != nil {
		query += fmt.Sprintf("is_active = $%d ", argNum)
		args = append(args, *updateReq.IsActive)
		argNum++
	}

	if query[len(query)-2] == ',' {
		query = query[:len(query)-2]
	}

	query += fmt.Sprintf(" WHERE id = $%d", argNum)
	args = append(args, id)

	_, err = db.Exec(query, args...)
	return err
}

func GetEmployerStatistic() (*models.EmployerStatistic, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stat models.EmployerStatistic

	row := db.QueryRow(`
        SELECT 
            COUNT(*) AS users,
            SUM(lemons) AS lemons,
            SUM(diamonds) AS diamonds
        FROM employers
    `)

	err = row.Scan(&stat.Users, &stat.Lemons, &stat.Diamonds)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}
