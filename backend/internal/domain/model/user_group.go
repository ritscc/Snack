package model

type UserGroup struct {
	GroupID     int64
	Name        string
	Description string
	Type        string
	Icon        string
	Role        string
	Member      []*Member
}

type Member struct {
	GroupId int64
	UserId  int64
	Grant   string
}
