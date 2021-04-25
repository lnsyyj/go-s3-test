package s3_adminosd_test

import (
	"fmt"
	radosAPI "github.com/QuentinPerez/go-radosgw/pkg/api"
	"testing"
)

func TestCreateUser(t *testing.T) {
	api, _ := radosAPI.New("http://object.yujiang.com:7480", "3DNM7Z57L4UXGDPMG3FU", "rWsCjB0u7GF4uPXTTaE0BU4rfNP33OE2WufJBJEt")
	// create a new user named JohnDoe
	user, err := api.CreateUser(radosAPI.UserConfig{
		UID:         "JohnDoe",
		DisplayName: "John Doe",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}

func TestDeleteUser(t *testing.T) {
	api, _ := radosAPI.New("http://object.yujiang.com:7480", "3DNM7Z57L4UXGDPMG3FU", "rWsCjB0u7GF4uPXTTaE0BU4rfNP33OE2WufJBJEt")
	// ...
	// remove JohnDoe
	err := api.RemoveUser(radosAPI.UserConfig{
		UID: "JohnDoe",
	})
	if err != nil {
		fmt.Println(err)
	}
}