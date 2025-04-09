package main

import (
    "log"
    "login-signup-gin/db"     // Import our db package
    "login-signup-gin/routes" // Import our routes package
)

func main() {
	err := db.InitDB()

	if err != nil {
		log.Fatal("DB Connection failed ", err)
	}

	router := routes.SetupRouter()

	log.Println("Server starting on port 8080")
	err = router.Run(":8080")

	if err!=nil{
		log.Fatal("Server failed to start ", err)
	}
}
