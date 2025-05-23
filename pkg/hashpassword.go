package pkg

import(
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPassword(hash,password string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}