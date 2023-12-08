package testools

import (
	"time"

	"github.com/cdvelop/model"
)

type AuthTest struct{}

func (AuthTest) BackendCheckUser(params any) (u *model.User, err string) {

	user := model.User{
		Token:          "123",
		Id:             "123456789101112",
		Ip:             "172.0.0.1",
		Name:           "don Juanito dev test",
		Area:           "s",
		AccessLevel:    "1",
		LastConnection: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &user, ""
}

func (AuthTest) UserAuthNumber() (number string, err string) {
	return "1", ""
}
func (AuthTest) NameOfAuthHandler() (name string) {
	return "login"
}
