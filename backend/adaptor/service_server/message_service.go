package service_server

import pbMessage "github.com/ritscc/Snack/pb/message/v1"

type MessageService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbMessage.UnimplementedMessageServiceServer
}

func InitMessageService() *MessageService {
	return &MessageService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}
