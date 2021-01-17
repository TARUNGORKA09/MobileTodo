package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Username string = "TARUN"
var password string = "12345	"

func DecodeCredentials(r *http.Request) (string, string) {
	password := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if password[0] != "Basic" {
		fmt.Errorf("Not a valid Authorization", password)
	}

	pass, err := base64.StdEncoding.DecodeString(password[1])
	if err != nil {
		fmt.Errorf("unable to convert from base64 to string", password)
	}
	pass1 := strings.SplitN(string(pass), ":", 2)
	//fmt.Fprint(rw, string(pass1[]))
	return string(pass1[0]), string(pass1[1])
}

type UserClaims struct {
	jwt.StandardClaims
	Username string
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Inavlid token")
	}
	if u.Username != Username  {
		return fmt.Errorf("Not valid Username")
	}
	return nil
}

var claims *UserClaims = &UserClaims{}
var Key = []byte{} 
for var i=1 ;i<=64;i++ {
	Key = append(key,byte(i))
}

func CreateToken(c *UserClaims) (string,error){

	t := jwt.NewWithClaims(jwt.SigningMethodES512,claims)


}

func ParseToken(r *http.Request) (UserClaims,error) {

	jwt.
}

func main()
