package model

import "time"

//// Testing instance of user ...
func TestUser() *User {
	return &User{
		Email:       "email@example.org",
		Password:    "password",
		Role:        EmployeeRole,
		Verified:    true,
		Name:        "Name",
		Surname:     "Surname",
		MiddleName:  "MiddleName",
		Sex:         SexMale,
		DateOfBirth: time.Time{}.AddDate(2000, 2, 2),
		Address:     "Minsk Pr. Nezavisimosti 22-222",
		Phone:       "+375-29-154-89-33",
		Photo:       "Photo",
	}

}

func TestPet() *Pet {
	return &Pet{
		Name:        "Name",
		Type:        PetTypeCat,
		Weight:      1,
		Diseases:    "Disease",
		Owner:       *TestUser(),
		PetPhotoURL: "/",
	}
}