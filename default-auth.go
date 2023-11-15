package testools

import (
	"time"

	"github.com/cdvelop/model"
)

type AuthTest struct{}

func (AuthTest) GetLoginUser(params any) (*model.User, error) {

	user := model.User{
		Token:          "123",
		Id:             "123456789101112",
		Ip:             "172.0.0.1",
		Name:           "don Juanito dev test",
		Area:           "s",
		AccessLevel:    "1",
		LastConnection: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &user, nil
}

func (AuthTest) UserAuthNumber() (string, error) {
	return "1", nil
}
