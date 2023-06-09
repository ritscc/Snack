package driver

import (
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


func StartServer(port net.Listener) {
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
	server.Serve(port)

	waitSIGINT()

	log.Println("stopping gRPC server...")
	server.GracefulStop()
}

func waitSIGINT() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)
}
