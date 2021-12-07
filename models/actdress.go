package models

type Actdress struct {
	ActId     int    `db:"act_id" json:"act_id"`
	ActNameEn string `db:"act_name_en" json:"act_name_en"`
	ActNameJp string `db:"act_name_jp" json:"act_name_jp"`
	Birth     string `db:"birth" json:"birth"`
	Tall      int    `db:"tall" json:"tall"`
	Cup       string `db:"cup" json:"cup"`
	Waist     int    `db:"waist" json:"waist"`
	Hip       int    `db:"hip" json:"hip"`
	Display   string `db:"display" json:"display"`
}

type ActTag struct {
	Title string `db:"tile" json:"title"`
}
