package controller

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"

	"github.com/ecommerce/db"
	u "github.com/ecommerce/models"
	q "github.com/ecommerce/utility"
)

type login struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type loginFacebook struct {
	Email		string	`json:"email"`
	Token		string	`json:"token"`
}

type claims struct {
	UserId	string	`json:"userId"`
	jwt.StandardClaims
}


//SignUpUser create user in db
func SignUpUser(w *fiber.Ctx)  {
	id, err := uuid.NewV4()
	if err != nil {
		w.Status(500).JSON("Error in uuid generate")
		return
	}

	var aux u.User
	aux.ID = id
	aux.Name = w.FormValue("name")
	aux.Password = w.FormValue("password")
	aux.Email = w.FormValue("email")
	aux.Phone = w.FormValue("phone")
	aux.Address = w.FormValue("address")
	aux.FacebookID = w.FormValue("facebookID")
	aux.Birthday = w.FormValue("birthday")
	aux.Gender = w.FormValue("gender")
	aux.Roles = pq.StringArray{"user"}

	var aux1 u.User
	db.DBConn.Where("email = ?", aux.Email).First(&aux1)
	if len(aux1.Email) != 0 {
		w.Status(400).JSON("Already has a user with this email")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(aux.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Status(500).JSON("Error in bcrypt")
		return
	}
	aux.Password = string(hashPassword)

	avatarFile, _ := w.FormFile("avatar")
	if avatarFile != nil {
		file, err := avatarFile.Open()
		key, url := q.SendImageToAWS(file, avatarFile.Filename, avatarFile.Size, "user")
		if key ==  "" || err != nil {
			w.Status(500).JSON("Error upload image")
			return
		}
		defer file.Close()

		aux.ImageKey = key
		aux.ImageURL = url
	} else {
		aux.ImageKey = ""
		aux.ImageURL = ""
	}

	result := db.DBConn.Create(&aux)
	if result.Error != nil {
		w.Status(500).JSON("Error creating user")
		q.DeleteImageInAWS(aux.ImageKey)
		return
	}

	fileEmail, err := ioutil.ReadFile("template/welcome.html")
	if err != nil {
		w.Status(201).JSON("User created, but failed sending email")
	}

	q.SendEmailUtility(aux.Email, string(fileEmail), "Welcome to Cash And Grab")

	w.Status(201).JSON("User created")
}

//Login user in application
func Login(w *fiber.Ctx) {
	login := new(login)
	if err := w.BodyParser(login); err != nil {
		w.Status(500).JSON("Missing fields")
		return
	}

	var user u.User
	result := db.DBConn.Where("email = ?", login.Email).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("Error listing user")
		return
	}

	if user.Email == "" {
		w.Status(500).JSON("No user with this email")
		return
	}

	hashPass := []byte(user.Password)
	bodyPass := []byte(login.Password)
	errorHash := bcrypt.CompareHashAndPassword(hashPass, bodyPass)
	if errorHash != nil {
		w.Status(500).JSON("Wrong password")
		return
	}

	expTime := time.Now().Add(4000 * time.Minute)
	claimsJwt := &claims{
		UserId: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	tokenMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)
	jwtKey := []byte(q.GetDotEnv("JWT_KEY"))
	token, err := tokenMethod.SignedString(jwtKey)
	if err != nil {
		w.Status(500).JSON("Error in jwt token")
		return
	}

	user.Password = ""

	w.Status(200).JSON(&fiber.Map{
		"user": user,
		"token": token,
	})
}

//Login user in application
func LoginAdmin(w *fiber.Ctx) {
	login := new(login)
	if err := w.BodyParser(login); err != nil {
		w.Status(500).JSON("Missing fields")
		return
	}

	var user u.User
	result := db.DBConn.Where("email = ?", login.Email).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("Error listing user")
		return
	}

	if user.Email == "" {
		w.Status(500).JSON("No user with this email")
		return
	}

	if len(user.Roles) == 1 {
		w.Status(401).JSON("User doenst have admin permission")
		return
	}

	hashPass := []byte(user.Password)
	bodyPass := []byte(login.Password)
	errorHash := bcrypt.CompareHashAndPassword(hashPass, bodyPass)
	if errorHash != nil {
		w.Status(500).JSON("Wrong password")
		return
	}

	expTime := time.Now().Add(4000 * time.Minute)
	claimsJwt := &claims{
		UserId: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	tokenMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)
	jwtKey := []byte(q.GetDotEnv("JWT_KEY"))
	token, err := tokenMethod.SignedString(jwtKey)
	if err != nil {
		w.Status(500).JSON("Error in jwt token")
		return
	}

	user.Password = ""

	w.Status(200).JSON(&fiber.Map{
		"user": user,
		"token": token,
	})
}

//LoginWithFacebook verify if user is logged with facebook
func LoginWithFacebook(w *fiber.Ctx)  {
	loginFacebook := new(loginFacebook)
	if err := w.BodyParser(loginFacebook); err != nil {
		w.Status(500).JSON("Missing fields")
		return
	}

	var user u.User
	result := db.DBConn.Where("email = ?", loginFacebook.Email).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("No user with this email")
		return
	}
	
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + loginFacebook.Token)
	if err != nil {
		w.Status(401).JSON("Invalid token")
		return
	}
	
	defer resp.Body.Close()

	expTime := time.Now().Add(4000 * time.Minute)
	claimsJwt := &claims{
		UserId: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	tokenMethod := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsJwt)
	jwtKey := []byte(q.GetDotEnv("JWT_KEY"))
	tokenJWT, err := tokenMethod.SignedString(jwtKey)
	if err != nil {
		w.Status(500).JSON("Error in jwt token")
		return
	}

	w.Status(200).JSON(&fiber.Map{
		"user": user,
		"token": tokenJWT,
	})
}

//ResetPassword change password
func ResetPassword(w *fiber.Ctx)  {
	email := w.FormValue("email")
	password := w.FormValue("password")
	repeatPassword := w.FormValue("reset")
	if password != repeatPassword {
		w.Status(500).JSON("Passwords don't match")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.Status(500).JSON("bcrypt error")
		return
	}

	result := db.DBConn.Model(&u.User{}).Where("email = ?", email).Update("password", hashPassword)
	if result.Error != nil {
		w.Status(500).JSON("Error reseting password")
		return
	}

	w.Status(200).JSON("Password changed")
}

//SendEmailToResetPassword send link to email to reset password
func SendEmailToResetPassword(w *fiber.Ctx)  {
	email := w.Query("email")

	if email == "" {
		w.Status(500).JSON("Error: Missing email in query")
		return
	}

	var user u.User
	result := db.DBConn.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("Error listing user")
		return
	}

	if user.Name == "" {
		w.Status(500).JSON("No user with this email")
		return
	}

	fileEmail, err := ioutil.ReadFile("template/resetPassword.html")
	if err != nil {
		w.Status(500).JSON("Server error")
	}

	rand.Seed(time.Now().UnixNano())
	hash := strconv.Itoa(rand.Intn(999999 - 100000 + 1) + 100000)

	q.SendEmailUtility(user.Email, string(fileEmail), "Reset Password - Código de Verificação: " + hash)

	w.Status(200).JSON("Email sended")
}

//Change roles in user field
func ToggleRolesUser(w *fiber.Ctx) {
	userId := w.Params("id")

	var user u.User
	result := db.DBConn.Where("id = ?", userId).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("User doenst exist in database")
		return
	}

	if len(user.Roles) > 1 {
		resultUpdate := db.DBConn.Model(&user).Where("id = ?", userId).Update("roles", pq.StringArray{"user"})
		if resultUpdate.Error != nil {
			w.Status(500).JSON("Error update in user roles")
			return
		}
	} else {
		resultUpdate := db.DBConn.Model(&user).Where("id = ?", userId).Update("roles", pq.StringArray{"user", "admin"})
		if resultUpdate.Error != nil {
			w.Status(500).JSON("Error update in user roles")
			return
		}
	}

	w.Status(200).JSON(user)
} 