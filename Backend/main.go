package main

import (
	"backend/configs"
	"backend/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Saludo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Backend - SOPES 2 :D"})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Se crea el servidor con GIN
	r := gin.Default()

	// Se aplican middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	//MongoDB conection
	//run database
	configs.ConnectDB()

	//Rutas
	r.GET("/", Saludo)
	routes.BooksRoute(r)

	// Se inicia el servidor en el puerto 3000
	r.Run(":3000")
}
