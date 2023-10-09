package model

const (
	STATUS_ONLINE  Status = 0
	STATUS_OFFLINE Status = 1
)

type User struct {
	UserID    int64
	UserName  string
	Nick      string
	RealName  string
	Password  [64]byte
	Avatar    string
	Role      string
	Locale    string
	RitsEmail string
	OwnEmail  string
	Comment   string
	Status    Status
	Services  *Service
}

type Status int32

type Service struct {
	Twitter string
	Github  string
	Discord string
}
