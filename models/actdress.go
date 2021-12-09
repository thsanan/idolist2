package models

type Actdress struct {
	ActId     int    `db:"act_id" json:"actid"`
	ActNameEn string `db:"act_name_en" json:"actnameen"`
	ActNameJp string `db:"act_name_jp" json:"actnamejp"`
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

type ActdressEvo struct {
	ActName string `json:"name"`
	Age     int    `json:"age"`
	Cup     string `json:"cup"`
	Tall    int    `json:"tall"`
	Waist   int    `json:"waist"`
	Hip     int    `json:"hip"`
	Display string `json:"display"`
}
