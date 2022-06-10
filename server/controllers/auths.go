package controllers

import (
	"encoding/json"
	"go-server/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthsController struct {
	session *mgo.Session
}

func NewAuthsController(s *mgo.Session) *AuthsController {
	return &AuthsController{s}
}

func (uc AuthsController) GetUser(c *fiber.Ctx) error {
	users := []models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(nil).All(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func (uc AuthsController) GetAdmin(c *fiber.Ctx) error {
	users := []models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(nil).All(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(users)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

type SignupRequest struct {
	UserName  string
	FirstName string
	LastName  string
	Email     string
	PhNumber  string
	Gender    string
	Password  string
}

func (uc AuthsController) Signup(c *fiber.Ctx) error {
	req := new(SignupRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if req.UserName == "" || req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" || req.Gender == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid Credentials for signup")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.Auths{
		UserID:    bson.NewObjectId(),
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		PhNumber:  req.PhNumber,
		Gender:    req.Gender,
		Password:  string(hash),
		AvatarID:  bson.NewObjectId(),
	}

	// create a new avatar also with default config
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Insert(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	token, exp, err := createJWTToken(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token":  token,
		"expiry": exp,
		"user":   user,
	})
}

type LoginRequest struct {
	UserName string
	Password string
}

func (uc AuthsController) Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if req.UserName == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	user := models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(user).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	token, exp, err := createJWTToken(user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"token":  token,
		"expiry": exp,
		"user":   user,
	})
}

func (uc AuthsController) AdminLogin(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if req.UserName == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	user := models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(user).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if user.Password != req.Password {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	token, exp, err := createJWTToken(user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"token":  token,
		"expiry": exp,
		"user":   user,
	})
}

func (uc AuthsController) GetOneOtherUser(c *fiber.Ctx) error {
	user := models.Auths{}
	if err := uc.session.DB(models.DatabaseName).C(models.AuthCollectionName).Find(nil).One(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

func createJWTToken(user models.Auths) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserID
	claims["exp"] = exp
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
