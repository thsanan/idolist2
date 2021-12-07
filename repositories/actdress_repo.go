package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/idolist/models"
)

type actdressDb struct {
	db *sqlx.DB
}

type actdressRepo interface {
	AddAct(act models.Actdress) (*models.Actdress, error)
	//Update(actId int, actRequest models.Actdress) (*models.Actdress, error)
	//Delete(actId int) (*models.Actdress, error)
	//ActAll() ([]models.Actdress, error)
	//ActById(actId int) (*models.Actdress, error)
}

func NewActDb(db *sqlx.DB) actdressRepo {
	return actdressDb{db}
}

func (actDb actdressDb) AddAct(act models.Actdress) (*models.Actdress, error) {
	query := "INSERT actdress (act_name_en, act_name_jp, birth, tall, cup, waist, hip, display) VALUES (?,?,?,?,?,?,?,?)"
	result, err := actDb.db.Exec(query, act.ActNameEn, act.ActNameJp, act.Birth, act.Tall, act.Cup, act.Waist, act.Hip, act.Display)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	actRes := models.Actdress{
		ActId:     int(id),
		ActNameEn: act.ActNameEn,
		ActNameJp: act.ActNameJp,
		Birth:     act.Birth,
		Tall:      act.Tall,
		Cup:       act.Cup,
		Waist:     act.Waist,
		Hip:       act.Hip,
		Display:   act.Display,
	}

	return &actRes, nil
}
