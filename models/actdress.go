package models

type Actdress struct {
	ActID   int    `db:"act_id", json:"act_id"`
	ActName string `db:"act_name", json:"act_name"`
}
