package interfaces

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// For Login
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	bytePlain := getBytePassword(plainPwd)
	byteHash := getBytePassword(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func GeneratePwd(pwd string) string {
	bytePassword := getBytePassword(pwd)
	hashed := hashAndSalt(bytePassword)
	return string(hashed)

}

func getBytePassword(strPwd string) []byte{
	return []byte(strPwd)
}

func hashAndSalt(pwd []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil{
		log.Println(err)

	}
	return string(hashed)
}

