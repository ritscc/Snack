package service_server

import pbChannel "github.com/ritscc/Snack/pb/channel/v1"

type ChannelService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbChannel.UnimplementedChannelServiceServer
}

func InitChannelService() *ChannelService {
	return &ChannelService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}

type MessageChannelService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbChannel.UnimplementedMessageChannelServiceServer
}

func InitMessageChannelService() *MessageChannelService {
	return &MessageChannelService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}
