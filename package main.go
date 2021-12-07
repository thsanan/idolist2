package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/idolist/models"
	"github.com/thsanan/idolist/repositories"
)

func main() {

	db, err := sqlx.Open("mysql", "sandland:IntelliP24.X@tcp(203.159.94.79:3306)/idolist")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	repoAct := repositories.NewActDb(db)

	actx := models.Actdress{

		ActNameEn: "xxx",
		ActNameJp: "xxxx",
		Birth:     "2000-01-01",
		Tall:      170,
		Cup:       "D",
		Waist:     100,
		Hip:       100,
		Display:   "dis",
	}
	act, err := repoAct.AddAct(actx)
	if err != nil {
		panic(err.Error())
	}

	log.Println(act)

}
