package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoHOST := os.Getenv("MONGO_HOST")
	mongoPORT := os.Getenv("MONGO_PORT")
	mongoUSER := os.Getenv("MONGO_USER")
	mongoPASS := os.Getenv("MONGO_PASS")
	mongoURL := fmt.Sprintf(`mongodb://%v:%v/?authSource=admin&readPreference=primary&directConnection=true&ssl=false`, mongoHOST, mongoPORT)

	fmt.Println(mongoURL)

	credential := options.Credential{
		Username: mongoUSER,
		Password: mongoPASS,
	}

	ctxMongo, cancel := context.WithTimeout(context.Background(), time.Second*10)
	clientOptions := options.Client().ApplyURI(mongoURL).SetAuth(credential).SetDirect(true)

	c, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error al crear cliente %v", err)
	}
	err = c.Connect(ctxMongo)
	if err != nil {
		log.Fatalf("Error al realizar conexion %v", err)
	}

	err = c.Ping(ctxMongo, nil)
	if err != nil {
		log.Fatalf("Error al conectar %v", err)
	}

	defer cancel()
	return c

}

// Crear instancia del cliente Mongo
var DB *mongo.Client = ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("sopes2").Collection(collectionName)
	return collection
}
