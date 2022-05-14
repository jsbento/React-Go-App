package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func postUser(c *gin.Context) {
	dbclient = connect()

	body, _ := ioutil.ReadAll(c.Request.Body)
	var newUser User
	json.Unmarshal(body, &newUser)
	newUser.Token = randSeq(15)

	collection := dbclient.Database("react-go-app").Collection("users")

	_, erro := collection.InsertOne(context.TODO(), newUser)
	if erro != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add user"})
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully added new user", "token": newUser.Token})
	}
}

func deleteUser(c *gin.Context) {
	username := c.Query("username")

	dbclient = connect()
	collection := dbclient.Database("react-go-app").Collection("users")
	filter := bson.D{
		{Key: "username", Value: username},
	}

	err := collection.FindOneAndDelete(context.TODO(), filter, options.FindOneAndDelete())
	if err != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Deletion of user failed"})
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully deleted " + username})
	}
}

func checkUserExists(c *gin.Context) {
	dbclient = connect()
	username := c.Query("username")
	collection := dbclient.Database("react-go-app").Collection("users")
	exists := bson.D{{Key: "username", Value: username}}
	var result bson.M
	err := collection.FindOne(context.TODO(), exists, options.FindOne()).Decode(&result)
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

func login(c *gin.Context) {
	dbclient = connect()

	body, _ := ioutil.ReadAll(c.Request.Body)
	var user User
	json.Unmarshal(body, &user)

	find := bson.D{{Key: "username", Value: user.Username}}

	var result User
	collection := dbclient.Database("react-go-app").Collection("users")
	findErr := collection.FindOne(context.TODO(), find, options.FindOne()).Decode(&result)
	fmt.Printf("%v\n", result)

	if findErr != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error finding user information"})
	} else {
		if result.Username == user.Username && result.Password == user.Password {
			token := randSeq(15)
			filter := bson.D{{Key: "username", Value: user.Username}}
			update := bson.M{"$set": bson.M{"token": token}}
			var updated User
			updateErr := collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate()).Decode(&updated)
			if updateErr != nil {
				disconnect(dbclient)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
			} else {
				disconnect(dbclient)
				c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully logged in", "token": token})
			}
		}
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
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate())

	if err != nil {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error updating user: " + username})
	} else {
		disconnect(dbclient)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated user: " + username})
	}
}
