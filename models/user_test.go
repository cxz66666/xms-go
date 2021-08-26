package models

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T)  {
	user:=User{
		ID: 67,
	}
	fmt.Println(user.GetAvatarURL())
}