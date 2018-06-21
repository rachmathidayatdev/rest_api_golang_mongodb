package repository

import (
	"github.com/rachmathidayatdev/go_crud_mongodb/src/modules/profile/model"
)

//ProfileRepository interface
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}