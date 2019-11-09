package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"theletter/models/user"
)

func GetAllUsers(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data" : user.GetAllUsers(db),
		})
	}
}
