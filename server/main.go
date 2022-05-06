package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbclient *mongo.Client

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func postUser(c *gin.Context) {
	dbclient = connect()

	body, _ := ioutil.ReadAll(c.Request.Body)
	var newUser User
	json.Unmarshal(body, &newUser)

	collection := dbclient.Database("react-go-app").Collection("users")

	_, erro := collection.InsertOne(context.TODO(), newUser)
	if erro != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add user"})
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully added new user"})
	}
}

func updateUserEmail(c *gin.Context) {
	username := c.Param("username")
	email := c.Query("email")
	dbclient = connect()

	filter := bson.M{
		"$set": bson.M{"username": username},
	}

	update := bson.M{
		"$set": bson.M{"email": email},
	}

	collection := dbclient.Database("react-go-app").Collection("users")
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, nil)

	if err != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error updating user: " + username})
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated user: " + username})
	}
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

func checkUserExists(c *gin.Context) {
	dbclient = connect()
	username := c.Query("username")
	collection := dbclient.Database("react-go-app").Collection("users")
	exists := bson.D{{Key: "username", Value: username}}
	var result bson.M
	err := collection.FindOne(context.TODO(), exists, nil).Decode(&result)
	if err != nil {
		disconnect(dbclient)
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "No user with username: " + username, "exists": false})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error accessing database"})
		}
	} else {
		_, erro := json.Marshal(result)
		disconnect(dbclient)
		if erro != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error converting bson"})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User found with username: " + username, "exists": result != nil})
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.PUT("/users/:username", updateUserEmail)
	router.GET("/users/find/:username", getUserByUsername)
	router.GET("/users/exists", checkUserExists)
	router.POST("/users", postUser)
	router.Run("localhost:8080")
}
