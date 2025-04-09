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

	log.Println("Server starting on port 8000")
	err = router.run(":8000")

	if err!=nil{
		log.Fatal("Server failed to start ", err)
	}
}
