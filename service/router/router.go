package router

import (
	"net/http"
	"strconv"
	"time"

	"uas-api/core"
	"uas-api/service/entities"

	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
	app *core.Application
}

type RouterContract interface {
	NewRouter() http.Handler
}

func RouterConstructor(gin *gin.Engine, app *core.Application) RouterContract {
	return &Router{
		gin: gin,
		app: app,
	}
}

func (r *Router) NewRouter() http.Handler {
	// Root route
	r.gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"msg":    "Hello world",
		})
	})

	// Login handler
	r.gin.POST("/login", func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dbUser entities.User
		r.app.Mysql.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser)
		if dbUser.ID != 0 {
			c.JSON(http.StatusOK, dbUser)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	// Get user data handler
	r.gin.GET("/user/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var user entities.User
		r.app.Mysql.First(&user, id)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// Update user data handler
	r.gin.PUT("/user/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var user entities.User
		r.app.Mysql.First(&user, id)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		var updatedUser entities.User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Weight = updatedUser.Weight
		user.Height = updatedUser.Height
		user.Age = updatedUser.Age

		r.app.Mysql.Save(&user)
		c.JSON(http.StatusOK, user)
	})

	// Record water intake handler
	r.gin.POST("/water", func(c *gin.Context) {
		var intake entities.WaterIntake
		if err := c.ShouldBindJSON(&intake); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		intake.Date = time.Now()

		r.app.Mysql.Create(&intake)
		c.JSON(http.StatusOK, intake)
	})

	// Get water intake history handler
	r.gin.GET("/water/:userId", func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.Param("userId"))
		var intakes []entities.WaterIntake
		r.app.Mysql.Where("user_id = ?", userId).Find(&intakes)
		c.JSON(http.StatusOK, intakes)
	})

	return r.gin
}
