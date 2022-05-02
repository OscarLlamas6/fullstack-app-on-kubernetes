package controllers

import (
	"backend/configs"
	"backend/models"
	"backend/responses"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var booksCollection *mongo.Collection = configs.GetCollection(configs.DB, "Books")
var validate = validator.New()

func CreateBook() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newBook models.BookStruct
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&newBook); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := booksCollection.InsertOne(ctx, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		fmt.Println("Libro creado satisfactoriamente")
		newID := result.InsertedID.(primitive.ObjectID)
		data := map[string]interface{}{"id": newID.Hex()}
		c.JSON(http.StatusCreated, responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: data})

	}
}

func GetBooks() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		myCursor, err := booksCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var myBooks []bson.M

		if err = myCursor.All(ctx, &myBooks); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var myResult []map[string]interface{}
		resultJson, _ := json.Marshal(myBooks)
		json.Unmarshal(resultJson, &myResult)

		c.JSON(http.StatusOK, responses.ArrResponse{Status: http.StatusOK, Message: "success", Data: myResult})

	}
}

func GetBook() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookId := c.Param("bookID")

		var bookAux models.BookStruct

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookId)

		err := booksCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&bookAux)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var data map[string]interface{}
		inrec, _ := json.Marshal(&bookAux)
		json.Unmarshal(inrec, &data)

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: data})

	}
}

func UpdateBook() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bookAux models.BookUpdateStruct
		defer cancel()

		if err := c.BindJSON(&bookAux); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		id, _ := primitive.ObjectIDFromHex(bookAux.Id)

		filter := bson.M{"_id": bson.M{"$eq": id}}
		update := bson.M{"$set": bson.M{"title": bookAux.Title,
			"author": bookAux.Author,
			"year":   bookAux.Year,
			"pages":  bookAux.Pages,
			"genre":  bookAux.Genre}}

		respuesta, error := booksCollection.UpdateOne(ctx, filter, update)
		if error != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": error.Error()}})
			return
		}

		var data map[string]interface{}
		inrec, _ := json.Marshal(&respuesta)
		json.Unmarshal(inrec, &data)

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: data})
	}
}

func DeleteBook() gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookID")
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(bookID)

		filter := bson.M{"_id": bson.M{"$eq": id}}

		respuesta, error := booksCollection.DeleteOne(ctx, filter)
		if error != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": error.Error()}})
			return
		}

		var data map[string]interface{}
		inrec, _ := json.Marshal(&respuesta)
		json.Unmarshal(inrec, &data)

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: data})

	}
}
