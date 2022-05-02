package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbclient *mongo.Client

type User struct {
	Username string `json:"username"`
	Email    string `email:"email"`
	Password string `json:"password"`
}

func postUser(c *gin.Context) {
	dbclient = connect()
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")
	collection := dbclient.Database("react-go-app").Collection("users")

	userDoc := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	res, err := collection.InsertOne(context.TODO(), userDoc)
	if err != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to add user"})
	}
	fmt.Printf("Sucessfully inserted with id: %v\n", res.InsertedID)
	disconnect(dbclient)
	c.IndentedJSON(http.StatusOK, userDoc)
}

func getUsers(c *gin.Context) {
	//collection := dbclient.Database("react-go-app").Collection("users")
}

func getUserByUsername(c *gin.Context) {
	dbclient = connect()
	username := c.Param("username")
	collection := dbclient.Database("react-go-app").Collection("users")
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}, nil).Decode(&result)
	if err != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found with username: " + username})
		return
	}
	_, res := json.Marshal(result)
	if res != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error converting bson"})
		return
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, result)
	}
}

func connect() *mongo.Client {
	uri := getDBURI("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func getDBURI(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
	return os.Getenv(key)
}

func disconnect(c *mongo.Client) {
	if c == nil {
		return
	}
	err := c.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:username", getUserByUsername)
	router.POST("/users", postUser)
	router.Run("localhost:8080")
}