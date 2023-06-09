package driver

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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

/*
 * Configuration the grpc server.
 */
type Server struct {
	port int
}

/*
 * Initalization the grpc server.
 */
func InitServer(port int) *Server {
	return &Server{
		port: port,
	}
}

/*
 * Start the grpc server.
 */
func (s *Server) StartServer() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	reflection.Register(server)

	auth := pbAuth.UnimplementedAuthenticationServer{}
	user := pbUser.UnimplementedUserServiceServer{}
	userGroup := pbUserGroup.UnimplementedUserGroupServiceServer{}
	memberGroup := pbUserGroup.UnimplementedMemberServiceServer{}
	// stamp := stamppb.
	message := pbMessage.UnimplementedMessageServiceServer{}
	pinMessage := pbMessage.UnimplementedPinMessageServiceServer{}
	event := pbEvent.UnimplementedEventServiceServer{}
	channel := pbChannel.UnimplementedChannelServiceServer{}
	messageChannel := pbChannel.UnimplementedMessageChannelServiceServer{}

	pbAuth.RegisterAuthenticationServer(server, auth)
	pbUser.RegisterUserServiceServer(server, user)
	pbUserGroup.RegisterUserGroupServiceServer(server, userGroup)
	pbUserGroup.RegisterMemberServiceServer(server, memberGroup)
	pbMessage.RegisterMessageServiceServer(server, message)
	pbMessage.RegisterPinMessageServiceServer(server, pinMessage)
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
