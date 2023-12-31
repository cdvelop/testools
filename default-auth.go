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

func (AuthTest) UserSessionNumber() (number string, err string) {
	return "1", ""
}
func (AuthTest) SessionHandlerName() (name string) {
	return "login"
}
func (AuthTest) BackendLoadBootData(u *model.User) (out model.BootPageData) {
	return model.BootPageData{
		JsonBootActions: "none",
	}
}
