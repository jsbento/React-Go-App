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

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	Username string `json:"username"`
	Email    string `email:"email"`
	Password string `json:"password"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postUser(c *gin.Context) {
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to add user"})
	}
	fmt.Printf("Sucessfully inserted with id: %v\n", res.InsertedID)
	c.IndentedJSON(http.StatusOK, userDoc)
}

func getUsers(c *gin.Context) {
	//collection := dbclient.Database("react-go-app").Collection("users")
}

func getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	collection := dbclient.Database("react-go-app").Collection("users")
	fmt.Printf("%s\n", username)
	//opts := options.FindOne().SetSort(bson.D{{Key: "username", Value: username}})
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}, nil).Decode(&result)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found with username: " + username})
		return
	}
	_, res := json.Marshal(result)
	if res != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error converting bson"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
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

var dbclient *mongo.Client

func main() {
	dbclient = connect()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/users", getUsers)
	router.GET("/users/:username", getUserByUsername)
	router.POST("/users", postUser)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}
