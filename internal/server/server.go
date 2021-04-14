package server

import (
	"context"
	"strconv"
	"time"

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"usersrvice/proto/user/v1"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyServer struct {
}

func (s *MyServer) ScheduleReminder(ctx context.Context, req *user.ScheduleReminderRequest) (*user.ScheduleReminderResponse, error) {
	if req.When == nil {
		return nil, status.Error(codes.InvalidArgument, "when cant be nil")
	}

	when, err := ptypes.Timestamp(req.GetWhen())
	if err != nil {
		return nil, status.Error(codes.Internal, "cant convert timestamp into time")
	}

	if when.Before(time.Now()) {
		return nil, status.Error(codes.InvalidArgument, "when should be in the future")
	}

	newTimerID := uuid.New().String()
	go func(id string, dur time.Duration) {
		timer := time.NewTimer(dur)
		<-timer.C
		log.Infof("Timer %s time!", newTimerID)
	}(newTimerID, when.Sub(time.Now()))

	return &user.ScheduleReminderResponse{
		Id: newTimerID,
	}, nil
}

func (s *MyServer) GetUsers(ctx context.Context, req *user.GetUsersRequest) (*user.UsersResponse, error) {
	fmt.Println("Go MySQL database")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM mydb.users")

	if err != nil {
		panic(err.Error())
	}
	var users []*user.User
	for rows.Next() {
		var user user.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Othername); err != nil {
			log.Println(err.Error())
		}

		users = append(users, &user)
	}
	defer rows.Close()
	return &user.UsersResponse{
		User: users,
	}, nil
}
func (s *MyServer) InsertUser(ctx context.Context, req *user.InsertUserRequest) (*user.InsertUserResponse, error) {
	fmt.Println("Go MySQL database")
	if len(req.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "name cant be nil")
	}
	if len(req.Surname) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Surname cant be nil")
	}

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	result, err := db.Exec("insert into mydb.users(name,surname,othername) values('" + req.GetName() + "','" + req.GetSurname() + "', '" + req.GetOthername() + "');")

	if err != nil {
		panic(err.Error())
	}
	if result == nil {
		return &user.InsertUserResponse{
			Message: "The user was inserted",
		}, nil
	}

	res, err := result.LastInsertId()

	return &user.InsertUserResponse{
		Message: strconv.FormatInt(res, 10),
	}, nil
}
