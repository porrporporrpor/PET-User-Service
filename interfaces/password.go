package interfaces

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

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

// for login
func ComparePassword(hashedPwd string, plainPwd []byte) bool{
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil{
		//passwordMissMatch
		log.Println(err)
		return false
	}
	return true

}
