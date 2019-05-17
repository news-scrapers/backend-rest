package models

import (
	u "backend-rest/utils"
	"context"
	"os"
	"strings"

	log "log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

/*
JWT claims struct
*/
type Token struct {
	UserId string
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	Email    string `bson:"email" json:"email,omitempty"`
	Password string `bson:"password" json:"password,omitempty"`
	Token    string `bson:"token" json:"token,omitempty"`
	ID       string `bson:"_id,omitempty" json:"_id,omitempty"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//check for errors and duplicate emails
	db := GetDB()
	collection := db.Collection("users")
	foundAccount := &Account{}
	err := collection.FindOne(context.Background(), bson.M{"email": account.Email}).Decode(foundAccount)

	if err != nil {
		log.Println(err)
	}

	if foundAccount.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}
	account.ID = bson.NewObjectId().Hex()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	db := GetDB()
	collection := db.Collection("users")
	res, err := collection.InsertOne(context.Background(), account)
	log.Println(res)

	if err != nil {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (resp map[string]interface{}, code int) {

	account := &Account{}
	db := GetDB()
	collection := db.Collection("users")
	foundAccount := &Account{}
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(foundAccount)
	if err != nil {
		log.Println(err)
		return u.Message(false, "Connection error. Please retry"), 500
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again"), 401
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resps := u.Message(true, "Logged In")
	resps["account"] = account
	return resps, 200
}

func GetUser(u string) *Account {

	acc := &Account{}
	db := GetDB()
	collection := db.Collection("users")
	foundAccount := Account{}
	err := collection.FindOne(context.Background(), bson.M{"_d": foundAccount.ID}).Decode(foundAccount)
	if err != nil { //User not found!
		log.Println(err)
		return nil
	}

	acc.Password = ""
	return acc
}
