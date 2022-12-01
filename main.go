package main

import "kagiri/GolangNotesApi/api/controllers"

func main() {
	server := controllers.Server{}
	server.Initialize()
	server.Run(":8080")
}
