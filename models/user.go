package models

type User struct {
	Id       string `db:"id" json:"id"`
	UserName string `db:"user_name" json:"userName"`
	Created  string `db:"creation_datetime" json:"createdDatetime"`
	Updated  string `db:"updated_datetime" json:"updatedDatetime"`
}
