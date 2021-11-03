package common

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return "", err
	// }
	return string(hashedPassword)

}
