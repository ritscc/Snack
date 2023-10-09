package service_server

import pbEvent "github.com/ritscc/Snack/pb/event/v1"

type EventService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbEvent.UnimplementedEventServiceServer
}

func InitEventService() *EventService {
	return &EventService{}
}
