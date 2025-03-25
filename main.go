package main

import (
	config "app/src/config"
	route "app/src/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  /* Connect Database */
  config.InitDB()
  defer config.CloseDB()
  /* Setup routers */
  port := ":" + os.Getenv("PORT")
  log.Println("Port: ", port)
  if port == ":" {
    port = ":3000"
  }

  router := gin.Default()
  route.InitRouter(router, port)

}
