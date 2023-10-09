package driver

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	service_server2 "github.com/ritscc/Snack/internal/adaptor/service_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbAuth "github.com/ritscc/Snack/pb/authentication/v1"
	pbUser "github.com/ritscc/Snack/pb/user/v1"
	pbUserGroup "github.com/ritscc/Snack/pb/user_group/v1"

	// pbStamp "github.com/ritscc/Snack/pb/stamp/v1"
	pbChannel "github.com/ritscc/Snack/pb/channel/v1"
	pbEvent "github.com/ritscc/Snack/pb/event/v1"
	pbMessage "github.com/ritscc/Snack/pb/message/v1"
)

type Server struct {
	serverConfig ServerConfig
}

func NewServer(serverConfig ServerConfig) *Server {
	return &Server{
		serverConfig: serverConfig,
	}
}

// StartServer start the grpc server.
func (s *Server) StartServer() error {
	address := fmt.Sprintf("%s:%d", s.serverConfig.Host, s.serverConfig.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)

	auth := service_server2.InitAuthenticationService()
	user := service_server2.InitUserService()
	userGroup := service_server2.InitUserGroupService()
	member := service_server2.InitMemberService()
	// stamp := stamppb.
	message := service_server2.InitMessageService()
	// pinMessage := pbMessage.UnimplementedPinMessageServiceServer{}
	event := service_server2.InitEventService()
	channel := service_server2.InitChannelService()
	messageChannel := service_server2.InitMessageChannelService()

	pbAuth.RegisterAuthenticationServer(server, auth)
	pbUser.RegisterUserServiceServer(server, user)
	pbUserGroup.RegisterUserGroupServiceServer(server, userGroup)
	pbUserGroup.RegisterMemberServiceServer(server, member)
	pbMessage.RegisterMessageServiceServer(server, message)
	// pbMessage.RegisterPinMessageServiceServer(server, pinMessage)
	pbEvent.RegisterEventServiceServer(server, event)
	pbChannel.RegisterChannelServiceServer(server, channel)
	pbChannel.RegisterMessageChannelServiceServer(server, messageChannel)

	log.Println("start gRPC server ...")
	server.Serve(listener)

	waitSIGINT()

	log.Println("stopping gRPC server...")
	server.GracefulStop()

	return nil
}

/*
 * Get the signal issued when stopping Server with "CTRL+C".
 * This is used for Graceful Shutdown.
 */
func waitSIGINT() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)
}
