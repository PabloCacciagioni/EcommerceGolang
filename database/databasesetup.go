package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Definición de la función: DBSet es una función que no recibe parámetros y retorna un puntero a un objeto mongo.Client, que representa una conexión al servidor MongoDB.
func DBSet() *mongo.Client {
	// Crear un nuevo cliente de MongoDB utilizando la URI de conexión
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// Si hay un error al crear el cliente, se registra el error y se finaliza la ejecución del programa
	if err != nil {
		log.Fatal(err)
	}
	// Crear un contexto con un tiempo de espera de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Asegurarse de cancelar el contexto al final de la función
	defer cancel()

	// Conectar el cliente al servidor MongoDB utilizando el contexto creado
	err = client.Connect(ctx)
	// Si hay un error al conectar, se registra el error y se finaliza la ejecución del programa
	if err != nil {
		log.Fatal(err)
	}

	// Verificar la conexión al servidor MongoDB
	err = client.Ping(context.TODO(), nil)
	// Si hay un error al hacer ping, se registra el error y se finaliza la ejecución del programa
	if err != nil {
		log.Println("failed to connect to mongoDB")
		return nil
	}
	// Si la conexión es exitosa, se imprime un mensaje de éxito
	fmt.Println("Connected to MongoDB!")
	// Devolver el cliente de MongoDB
	return client
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
