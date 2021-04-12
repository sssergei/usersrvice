package main

import (
	"context"
	"time"

	"usersrvice/proto/user/v1"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx, _ := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	clientCert, err := credentials.NewClientTLSFromFile("../../server.crt", "")
	if err != nil {
		log.Fatalln("failed to create cert", err)
	}

	reminderConn, err := grpc.DialContext(ctx, "localhost:8080",
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		log.Fatalln("Failed to dial server: ", err)
	}

	reminderClient := user.NewUserServiceClient(reminderConn)
	fiveSeconds, _ := ptypes.TimestampProto(time.Now().Add(5 * time.Second))
	resp, err := reminderClient.ScheduleReminder(ctx,
		&user.ScheduleReminderRequest{
			When: fiveSeconds,
		})
	if err != nil {
		log.Fatalln("Failed to schedule a user: ", err)
	}
	log.Infof("User have been successfully scheduled. New  user id is %s", resp.GetId())
}
