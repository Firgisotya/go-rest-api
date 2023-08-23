package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/Firgisotya/go-rest-api/app/middlewares"
	// "github.com/Firgisotya/go-rest-api/app/middlewares"
	"github.com/Firgisotya/go-rest-api/app/routes"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/Firgisotya/go-rest-api/config/command"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [function]")
		return
	}

	functionRun := os.Args[1]

	switch functionRun {
	case "db:migrate":
		command.Migrate()
	case "db:seed":
		command.Seed()
	case "serve":
		config.ConnectDB()
		if key := os.Getenv("JWT_SECRET"); key == "" {
			fmt.Println("Your key is null, please run `go run main.go key:generate`")
		} else {
			
			fmt.Println("Starting Go API Server...")

			// Baca variabel lingkungan dari file .env
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

			// Periksa apakah variabel GIN_MODE ada di lingkungan
			ginMode := os.Getenv("GIN_MODE")
			if ginMode == "development" {
				ginMode = gin.DebugMode
			} else if ginMode == "release" {
				gin.SetMode(gin.ReleaseMode)
			} else {
				log.Fatalf("Invalid GIN_MODE value: %s. Should be 'release' or empty.", ginMode)
			}

			// Inisialisasi router Gin
			// router := gin.Default()
			router := gin.New()

			// Untuk CORS
			// router.Use(middlewares.CorsMiddleware())
			router.Use(cors.Default())

			// Setup rute untuk pengguna (user) dan produk (product)
			routes.SetupRouter(router)

			// Jalankan server
			router.Run(":9000")
		}

	}
	

}
