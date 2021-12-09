package services

import (
	"github.com/thsanan/idolist/models"
	"github.com/thsanan/idolist/repositories"
)

type ActressService interface {
	//GetActEvoById(actId int) (*ActdressEvo, error)
	//GetActAll() ([]models.ActdressEvo, error)
	InsertAct(act models.Actdress) (*models.Actdress, error)
}

type actressEvo struct {
	actRepo repositories.ActdressRepo
}

func NewActressEvo(repo repositories.ActdressRepo) ActressService {
	return actressEvo{repo}
}

func (actRepo actressEvo) InsertAct(act models.Actdress) (*models.Actdress, error) {

	actEvo, err := actRepo.actRepo.AddAct(act)
	if err != nil {
		return nil, err
	}

	return actEvo, nil
}
