package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type RequestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var mongoClient *mongo.Client


func connectMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	log.Println("Connected to MongoDB")
	return client
}

func authenticateUser(username, password string) bool {
	collection := mongoClient.Database("Account").Collection("VPN")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		return false
	}
	return true
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 1 {
		http.Error(w, "Invalid path format", http.StatusBadRequest)
		return
	}

	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	authSuccess := authenticateUser(requestData.Username, requestData.Password)

	response := map[string]interface{}{
		"path":     pathParts,
		"username": requestData.Username,
	}
	if authSuccess {
		response["message"] = "Login success"
	} else {
		response["message"] = "Invalid credentials"
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	mongoClient = connectMongoDB()
	defer mongoClient.Disconnect(context.Background())

	port := 2234
	addr := fmt.Sprintf("127.0.0.1:%d", port)

	http.HandleFunc("/", requestHandler)

	log.Printf("Server is running on %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
