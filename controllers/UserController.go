package controller

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/gofrs/uuid"
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
	email	string
	jwt.StandardClaims
}

// swagger:operation POST /signUp user response
// Adiciona novo usuário
// ---
// responses:
//     '201':
//         description: User created
//     '500':
//         description: Error

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

// swagger:operation POST /login user response
// Login na aplicação
// ---
// responses:
//     '200':
//         description: Login
//     '500':
//         description: Error

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
		email: login.Email,
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
	
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + loginFacebook.Token)
	if err != nil {
		w.Status(401).JSON("Invalid token")
		return
	}
	
	defer resp.Body.Close()

	var user u.User
	result := db.DBConn.Where("email = ?", loginFacebook.Email).Find(&user)
	if result.Error != nil {
		w.Status(500).JSON("No user with this email")
		return
	}

	expTime := time.Now().Add(4000 * time.Minute)
	claimsJwt := &claims{
		email: user.Email,
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
	email := w.Params("email")

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

	q.SendEmailUtility(user.Email, string(fileEmail), "Reset Password")

	w.Status(200).JSON("Email sended")
}