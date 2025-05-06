package handlers

import (
	"motivora-backend/internal/db"
	"motivora-backend/internal/models"
)

func GetCompanies(offset, limit int) ([]models.Company, error) {
	db, err := db.ConnectDB()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, title, lemons, diamonds, is_active FROM companies OFFSET $1 LIMIT $2", offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []models.Company

	for rows.Next() {
		var c models.Company
		err := rows.Scan(&c.ID, &c.Title, &c.Lemons, &c.Diamonds, &c.IsActive)
		if err != nil {
			return nil, err
		}

		companies = append(companies, c)
	}

	return companies, nil
}

func GetCompaniesStatistic() (*models.CompaniesStatistic, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stat models.CompaniesStatistic

	row := db.QueryRow(`
        SELECT 
            COUNT(*) AS companies,
            SUM(lemons) AS lemons,
            SUM(diamonds) AS diamonds
        FROM companies
    `)

	err = row.Scan(&stat.Companies, &stat.Lemons, &stat.Diamonds)
	if err != nil {
		return nil, err
	}

	return &stat, nil
}
