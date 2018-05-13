package common

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"encoding/json"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2/bson"
)

//using asymmetric keys
var (
	privKeyPath, _ = filepath.Abs("keys/app.rsa")
	pubKeyPath, _  = filepath.Abs("keys/app.rsa.pub")
)

// private key for signing and public key for verification
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//Read the key before starting the http handlers
func initKeys() {
	var err error
	signKeyByte, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)

	verifyKeyByte, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
}

//Generate JWT token
func GenerateJWT(email string, id bson.ObjectId, role string) (string, error) {
	//create a signer for rsa256
	t := jwt.New(jwt.SigningMethodRS256)

	//set claims for the JWT token
	claims := t.Claims.(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["UserInfo"] = struct {
		Email string
		Id	bson.ObjectId
		Role string
	}{email, id, role}

	//set the expire time for the jwt token
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//Middleware for validating the JWT tokens
func Authorize(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//validate the token
		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			//verify the token with public key, which is the counter part of private key
			return verifyKey, nil
		})

		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError: // JWT validation error
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired: // JWT expired
					DisplayAppError(
						w, err, "Access Token is expired, get a new Token",
						401,
					)
					return

				default:
					DisplayAppError(
						w, err, "Error while parsing the Access Token!",
						500,
					)
					return
				}
			default:
				DisplayAppError(
					w, err, "Error while parsing the Access Token!",
					500,
				)
				return
			}
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userInfo,err := json.Marshal(claims["UserInfo"])
			if err != nil {
				LogAppError(err)
			}
			var userObject map[string]string
			err = json.Unmarshal(userInfo,&userObject)
			if err != nil {
				LogAppError(err)
			}
			context.Set(r, "userEmail", userObject["Email"])
			context.Set(r, "userId", userObject["Id"])
			LogAppInfo(userObject["Email"] + " Authenticated Successfully")
			next.ServeHTTP(w, r)

		} else {
			DisplayAppError(w, err, "Invalid Access Token!",
				401,
			)
		}
	}
	return http.HandlerFunc(fn)
}
