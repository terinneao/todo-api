package db

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	Code string
	// Name  string
	Email string
	Todo1 string
	Todo2 string
	Todo3 string
}

func ReadAll() (*[]UserDB, error) {
	db := DB()

	var userDB []UserDB
	// userDB := []UserDB{}
	if err := db.Find(&userDB).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Read(email string) (*UserDB, error) {
	db := DB()

	var userDB UserDB

	// if err := db.First(&userDB, "Code = ?", code).Error; err != nil {
	if err := db.First(&userDB, "Email = ?", email).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}

func Create(user UserDB) error {
	db := DB()
	userData := UserDB{
		Code: user.Code,
		// Name:  user.Name,
		Email: user.Email,
		Todo1: user.Todo1,
		Todo2: user.Todo2,
		Todo3: user.Todo3,
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}
	return nil
}

func Update(code string, user UserDB) error {
	db := DB()

	var userDB UserDB
	if err := db.Find(&userDB, "Code = ?", code).Error; err != nil {
		// if err := db.Find(&userDB, "Email = ?", email).Error; err != nil {
		return err
	}

	// userDB.Name = user.Name
	userDB.Email = user.Email
	if err := db.Save(&userDB).Error; err != nil {
		return err
	}

	return nil
}

func Delete(code string) (*UserDB, error) {
	db := DB()

	var userDB UserDB

	if err := db.Unscoped().Delete(&userDB, "Code = ?", code).Error; err != nil {
		return nil, err
	}

	return &userDB, nil
}
