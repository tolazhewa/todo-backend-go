package models

type Todo struct {
	Id        string `db:"id" json:"id"`
	UserId    string `db:"user_id" json:"userId"`
	Completed bool   `db:"completed" json:"completed"`
	Title     string `db:"title" json:"title"`
	Created   string `db:"creation_datetime" json:"createdDatetime"`
	Updated   string `db:"updated_datetime" json:"updatedDatetime"`
}
