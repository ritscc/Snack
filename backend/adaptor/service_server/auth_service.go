package service_server

import pbAuth "github.com/ritscc/Snack/pb/authentication/v1"

type AuthenticationService struct {
	// converter  *converter.UserGroupConverter
	// interactor *interactor.UserGroupInteractor
	pbAuth.UnimplementedAuthenticationServer
}

func InitAuthenticationService() *AuthenticationService {
	return &AuthenticationService{
		// converter:  converter.NewUserConverter(),
		// interactor: interactor.NewUserInteractor(),
	}
}
