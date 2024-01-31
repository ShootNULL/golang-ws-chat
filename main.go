package main

import (
	"log"

	"net/http"
)

func main() {

	//mainRouter := gin.Default()

	connectDB()
	//log.Print(DB)
	setupWsRouter()
	//mainRouter.Run("localhost:8080")
}

func setupWsRouter() {
	//router.GET("/getUsers", getUsers)

	http.HandleFunc("/ws", handleWsConnections)
	http.HandleFunc("/auth", userAuth)
	http.HandleFunc("/register", userRegister)

	go handleMessages()

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}

}
