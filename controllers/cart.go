package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"james.practice/ecommerce/database"
)

type Application struct {
	proCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(proCollection *mongo.Collection, userCollection *mongo.Collection) *Application {
	return &Application{
		proCollection: proCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" { 
			log.Println("Product ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is empty"))
			return 
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" { 
			log.Println("User ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return 
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err = database.AddProductToCart(ctx, app.proCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Product added to cart"})
	}
}

func RemoveItem(app *Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" { 
			log.Println("Product ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is empty"))
			return 
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" { 
			log.Println("User ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return 
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.proCollection, app.userCollection, productID, userQueryID)
		if err != nil { 
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully remove ", productID,   " item from cart"})
	}
}

func GetItemFromCart() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" { 
			log.Println("User ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("User ID is empty"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx, app.proCollection, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully bought from cart"})
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" { 
			log.Println("Product ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("Product ID is empty"))
			return 
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" { 
			log.Println("User ID is empty")

			c.AbortWithError(http.StatusBadRequest, errors.New("User ID is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return 
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.proCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Product ", productID, " added to cart"})
	}
}