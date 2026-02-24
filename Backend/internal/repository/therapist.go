package repository

import (
	"database/sql"
	"sispa-backend/internal/domain"
)

type sqlTherapistRespository struct {
	DB *sql.DB
}

func NewTherapistRepository(db *sql.DB) domain.TherapistRepository {
	return &sqlTherapistRespository{DB: db}
}

func (r *sqlTherapistRespository) Create(therapist *domain.Therapist) error {
	query := `
		INSERT INTO therapists (name, joined_date, phone, address, status)
		VALUES (?,?,?,?,?)
	`
	_, err := r.DB.Exec(query, &therapist.Name, &therapist.JoinedDate, &therapist.Phone, &therapist.Address, &therapist.Status)

	return err
}

func (r *sqlTherapistRespository) GetByID(id int) (*domain.Therapist, error) {
	query := `
        SELECT id, name, phone, address, commission_rate, joined_date, exit_date, status, created_at, updated_at 
        FROM therapists 
        WHERE id = ?
    `

	var t domain.Therapist

	err := r.DB.QueryRow(query, id).Scan(&t.ID, &t.Name, &t.Phone, &t.Address, &t.CommissionRate, &t.JoinedDate, &t.ExitDate, &t.Status, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Or a custom "Not Found" error
		}
		return nil, err
	}

	return &t, nil
}

func (r *sqlTherapistRespository) GetAll() ([]domain.Therapist, error) {
	query := `
        SELECT id, name, phone, address, commission_rate, joined_date, exit_date, status, created_at, updated_at 
        FROM therapists
    `

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	therapist := []domain.Therapist{}

	for rows.Next() {
		var t domain.Therapist
		err := rows.Scan(&t.ID, &t.Name, &t.Phone, &t.Address, &t.CommissionRate, &t.JoinedDate, &t.ExitDate, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		therapist = append(therapist, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return therapist, nil
}

func (r *sqlTherapistRespository) Update(therapist *domain.Therapist) (*domain.Therapist, error) {
	query := `
        UPDATE therapists 
        SET name = ?, phone = ?, address = ?, Commission_rate = ?, Joined_date = ?, Exit_date = ?, status = ?
        WHERE id = ?
    `

	_, err := r.DB.Exec(query,
		therapist.Name,
		therapist.Phone,
		therapist.Address,
		therapist.CommissionRate,
		therapist.JoinedDate,
		therapist.ExitDate,
		therapist.Status,
		therapist.ID,
	)

	if err != nil {
		return nil, err
	}

	return therapist, nil
}

func (r *sqlTherapistRespository) Delete(id int) error {
	query := `
		DELETE FROM therapists WHERE id=?
    `

	_, err := r.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
