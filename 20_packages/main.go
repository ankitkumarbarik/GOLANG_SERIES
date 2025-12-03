package main

// go mod init <project_name>
// go add <package_name>
// go mod tidy

import (
	"fmt"

	"github.com/ankitkumarbarik/podcast/auth"
	"github.com/ankitkumarbarik/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("admin", "123")
	sess := auth.GetSession()
	fmt.Println(sess)

	user := user.User{
		Email: "admin@gmail.com",
		Name:  "Admin",
	}

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	color.Red("Name: %s", user.Name)
	color.Green("Email: %s", user.Email)
}
