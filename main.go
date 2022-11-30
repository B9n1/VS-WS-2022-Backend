package main

import (
    "context"
    "fmt"
    "log"
    "time"
    "encoding/json"
    "net/http"


    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

    /*
            Connect to my cluster
    */
    client, err := mongo.NewClient(options.Client().ApplyURI(
        "mongodb://localhost:27017/ToDo"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)
    
      /*
            List databases
    */
    databases, err := client.ListDatabaseNames(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(databases)

}
