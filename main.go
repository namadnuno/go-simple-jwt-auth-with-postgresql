package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"log"
	"os"
	"theletter/auth"
	"theletter/controllers"
	"flag"
	"theletter/database"
	"theletter/models"
)

func main() {
	seed := flag.Bool("seed", false, "enable database seeding")
	migrate := flag.Bool("migrate", false, "migrate the database")
	flag.Parse()

	db := pg.Connect(&pg.Options{
		User: "nunoalexandre",
		Password: "",
		Database: "theletter",
	})

	defer db.Close()

	if *migrate && *seed {
		err := models.CreateSchema(db)
		if err != nil {
				panic(err)
		}

		database.Seed(db)

		fmt.Println("Run make dev next.\n\n\n\n")
		os.Exit(1)
	}


	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := auth.AuthMiddleware(db);

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"userID": "oi",
		})
	})

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", controllers.GetAllUsers(db))
	}


	r.Run(":8080")
}
