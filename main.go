package main

import (
    "fmt"
    "log"
    "os"
    "encoding/json"
)

var config struct {
    ConnectionString string `json:"db_connection"`
    Mapping map[string]interface{}
}

func init() {
    configFile, err := os.Open("config.json")

    if err != nil {
        fmt.Printf("[Error]: %s \n", err.Error())
        os.Exit(1)
    }

    err = json.NewDecoder(configFile).Decode(&config)

    if err != nil {
        fmt.Printf("[Error]: %s \n", err.Error())
        os.Exit(1)
    }
}

func main() {
    connection := Repository{}
    err := connection.NewConnection(config.ConnectionString)

    if err != nil {
        log.Fatal(err)
    }

    defer connection.Close()

    fmt.Println("Working?")
}