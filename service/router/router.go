package router

import (
	"log"
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

		// Log the received username and password
		log.Printf("Received login request: username=%s, password=%s", user.Username, user.Password)

		var dbUser entities.User
		err := r.app.Mysql.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error
		if err != nil {
			log.Printf("Login failed: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, dbUser)
	})

	// Create user handler
	r.gin.POST("/user", func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		r.app.Mysql.Create(&user)
		c.JSON(http.StatusCreated, user)
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

	// Delete user handler
	r.gin.DELETE("/user/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var user entities.User
		r.app.Mysql.First(&user, id)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		r.app.Mysql.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"msg": "User deleted"})
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
		c.JSON(http.StatusCreated, intake)
	})

	// Get water intake history handler
	r.gin.GET("/water/:userId", func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.Param("userId"))
		var intakes []entities.WaterIntake
		r.app.Mysql.Where("user_id = ?", userId).Find(&intakes)
		c.JSON(http.StatusOK, intakes)
	})

	// Update water intake handler
	r.gin.PUT("/water/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var intake entities.WaterIntake
		r.app.Mysql.First(&intake, id)
		if intake.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Water intake not found"})
			return
		}

		var updatedIntake entities.WaterIntake
		if err := c.ShouldBindJSON(&updatedIntake); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		intake.Amount = updatedIntake.Amount
		r.app.Mysql.Save(&intake)
		c.JSON(http.StatusOK, intake)
	})

	// Delete water intake handler
	r.gin.DELETE("/water/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var intake entities.WaterIntake
		r.app.Mysql.First(&intake, id)
		if intake.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Water intake not found"})
			return
		}

		r.app.Mysql.Delete(&intake)
		c.JSON(http.StatusOK, gin.H{"msg": "Water intake deleted"})
	})

	return r.gin
}
