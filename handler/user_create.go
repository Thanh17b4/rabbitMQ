package handler

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"io/ioutil"
	"net/http"

	"Thanh17b4/practice/handler/responses"

	model "Thanh17b4/practice/model"
)

type UserService interface {
	GetListUser(page int64, limit int64) ([]*model.User, error)
	GetDetailUser(userID int64) (*model.User, error)
	UpdateUserService(user *model.User) (*model.User, error)
	DeleteUser(userID int64) (int64, error)
	CreateUser(user *model.User) (*model.User, error)
	GetDetailUserByEmail(email string) (*model.User, error)
}
type UserHandle struct {
	userService UserService
}

func NewUserHandle(userService UserService) *UserHandle {
	return &UserHandle{userService: userService}
}

func (h UserHandle) CreatUserHandle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user *model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		responses.Error(w, r, 400, err, "could not marshal your rq")
		return
	}
	userRes, err := h.userService.CreateUser(user)
	if err != nil {
		responses.Error(w, r, http.StatusInternalServerError, err, "could not creat user")
		return
	}

	err = h.SendCreateMessage(user)
	if err != nil {
		responses.Error(w, r, http.StatusInternalServerError, err, "failed to send create massage")
	}
	responses.Success(w, r, http.StatusCreated, userRes)
}

func (h UserHandle) SendCreateMessage(user *model.User) error {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	message, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"topic",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
