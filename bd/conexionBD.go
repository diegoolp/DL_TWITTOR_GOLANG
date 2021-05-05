package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto conexión a la BD*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:m5FdZlMmg7R6@dltwittorgolang.waolh.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD es la función que me permite conectar la BD*/
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
