package store

import (
	"errors"
	"log"
	"user/domain/model"
)

// PetRepository ...
type PetRepository struct {
	Store *Store
}

// Create pet and save it to DB
func (r *PetRepository) Create(p *model.Pet) (*model.Pet, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO pet (name, type, weight, diseases, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		p.Name,
		string(p.Type),
		p.Weight,
		p.Diseases,
		p.Owner.UserID,
	).Scan(&p.PetID); err != nil {
		log.Print(err)
		return nil, err
	}
	return p, nil
}

// GetAll returns all pets
func (r *PetRepository) GetAll() (*[]model.PetDTO, error) {
	rows, err := r.Store.Db.Query("SELECT * FROM pet")
	if err != nil {
		log.Print(err)
	}
	pets := []model.PetDTO{}

	for rows.Next() {
		pet := model.PetDTO{}
		err := rows.Scan(
			&pet.PetID,
			&pet.Name,
			&pet.Type,
			&pet.Weight,
			&pet.Diseases,
			&pet.OwnerID,
		)
		if err != nil {
			log.Print(err)
			continue
		}
		pets = append(pets, pet)
	}
	return &pets, nil
}

//FindByID searchs and returns petDTO by ID
func (r *PetRepository) FindByID(id int) (*model.PetDTO, error) {
	pet := &model.PetDTO{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE id = $1",
		id).Scan(
		&pet.PetID,
		&pet.Name,
		&pet.Type,
		&pet.Weight,
		&pet.Diseases,
		&pet.OwnerID,
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return pet, nil
}

//FindByID searchs and returns pet by ID
func (r *PetRepository) FindPetID(id int) (*model.Pet, error) {
	pet := &model.Pet{}
	if err := r.Store.Db.QueryRow("SELECT * FROM pet WHERE id = $1",
		id).Scan(
		&pet.PetID,
		&pet.Name,
		&pet.Type,
		&pet.Weight,
		&pet.Diseases,
		&pet.Owner.UserID,
	); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return pet, nil
}

// Delete pet from DB by ID
func (r *PetRepository) Delete(id int) error {
	result, err := r.Store.Db.Exec("DELETE FROM pet WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("no rows affected")
	}
	log.Printf("Pet deleted, rows affectet: %d", result)
	return nil
}

// Update pet from DB
func (r *PetRepository) Update(p *model.Pet) error {

	result, err := r.Store.Db.Exec(
		"UPDATE pet SET name = $1, type = $2, weight = $3, diseases = $4, user_id = $5 WHERE id = $6",
		p.Name,
		string(p.Type),
		p.Weight,
		p.Diseases,
		p.Owner.UserID,
		p.PetID,
	)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	log.Printf("Pet updated, rows affectet: %d", result)
	return nil
}