package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server/models"
	"go-server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthsController struct {
	db             *mongo.Client
	ctx            context.Context
	authCollection *mongo.Collection
}

func NewAuthsController(db *mongo.Client, ctx context.Context) *AuthsController {
	authCollection := models.GetCollection(db, models.AuthCollectionName)
	return &AuthsController{db, ctx, authCollection}
}

// TODO: imcomplete: get user from token
func (uc AuthsController) GetUser(c *fiber.Ctx) error {
	user := models.Auths{}
	err := uc.authCollection.FindOne(context.Background(), bson.D{}).Decode(&user)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(back)
}

// TODO: imcomplete: get admin from token
func (uc AuthsController) GetAdmin(c *fiber.Ctx) error {
	user := []models.Auths{}
	err := uc.authCollection.FindOne(context.Background(), bson.D{{}}).Decode(&user)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(user)
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
		UserID:    primitive.NewObjectID(),
		UserName:  req.UserName,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		PhNumber:  req.PhNumber,
		Gender:    req.Gender,
		Password:  string(hash),
		AvatarID:  primitive.NewObjectID(),
	}

	// create a new avatar also with default config

	res, err := uc.authCollection.InsertOne(context.Background(), bson.M{
		"UserID":    user.UserID,
		"UserName":  user.UserName,
		"FirstName": user.FirstName,
		"LastName":  user.LastName,
		"Email":     user.Email,
		"PhNumber":  user.PhNumber,
		"Gender":    user.Gender,
		"Password":  user.Password,
		"AvatarID":  user.AvatarID,
	})

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	fmt.Println(res)

	token, exp, err := createJWTToken(user)
	if err != nil {
		return utils.HandleControllerError(c, err)
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
	collection := models.GetCollection(uc.db, models.AuthCollectionName)
	err := collection.FindOne(context.Background(), bson.M{"UserName": req.UserName}).Decode(&user)

	if err != nil {
		return utils.HandleControllerError(c, err)
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
	collection := models.GetCollection(uc.db, models.AuthCollectionName)
	err := collection.FindOne(context.Background(), bson.M{"UserName": req.UserName}).Decode(&user)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid login credentials")
	}

	token, exp, err := createJWTToken(user)
	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	return c.JSON(fiber.Map{
		"token":  token,
		"expiry": exp,
		"user":   user,
	})
}

func (uc AuthsController) GetOneOtherUser(c *fiber.Ctx) error {
	user := models.Auths{}
	collection := models.GetCollection(uc.db, models.AuthCollectionName)
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&user)

	if err != nil {
		return utils.HandleControllerError(c, err)
	}

	back, err := json.Marshal(user)
	if err != nil {
		return utils.HandleControllerError(c, err)
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
