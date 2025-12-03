package auth

import "fmt"

// loginWithCredentials - only for internal usage
// LoginWithCredentials - for external usage
func LoginWithCredentials(username string, password string) {
	fmt.Println("login user using", username, "and", password)
}
