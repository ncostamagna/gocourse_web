package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {
	user.ID = uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("user created with id: ", user.ID)
	return nil
}

func (repo *repo) GetAll() ([]User, error) {
	var u []User

	if err := repo.db.Model(&u).Order("created_at desc").Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil

}

func (repo *repo) Get(id string) (*User, error) {
	user := User{ID: id}

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *repo) Delete(id string) error {
	user := User{ID: id}

	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {

	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}
	if phone != nil {
		values["phone"] = *phone
	}

	if err := repo.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}
