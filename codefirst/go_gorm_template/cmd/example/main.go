package main

import (
	"fmt"
	"gorm_template/internal/configs"
	"gorm_template/internal/core/domain"
	"gorm_template/internal/core/ports"
	"gorm_template/internal/infrastructure/db"
	repo "gorm_template/internal/infrastructure/repository"
	"gorm_template/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	cfg, _ := configs.LoadConfig()
	fmt.Printf("config loaded %s \n", cfg.ExamplePort)
	fmt.Printf("config loaded %s \n", cfg.DBDriver)
	fmt.Printf("config loaded %s \n", cfg.DNS)

	var dbConn *gorm.DB
	var err error
	var exampleRepo ports.ExamplePort

	switch cfg.DBDriver {
	case "postgres":
		dbConn, err = db.ConnectPostgres(cfg.DNS)
		exampleRepo = repo.NewExampleRepositoryPostgres(dbConn)
	case "sqlserver":
		dbConn, err = db.ConnectPostgres(cfg.DNS)
		exampleRepo = repo.NewExampleRepositorySQLServer(dbConn)
	}

	fmt.Printf("using %s \n", cfg.DBDriver)

	if err != nil {
		log.Fatal("cannot init db connection")
	}

	fmt.Println("prepare to migrate")
	dbConn.AutoMigrate(&domain.Example{})
	fmt.Println("done migrate")

	exampleService := services.NewExampleService(exampleRepo)

	r := gin.Default()

	r.POST("/example", func(c *gin.Context) {
		var x domain.Example
		if err := c.ShouldBindJSON(&x); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := exampleService.CreateExample(&x); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, x)
	})

	r.PUT("/example/:id", func(c *gin.Context) {
		var x domain.Example
		id := c.Param("id")
		if err := c.ShouldBindJSON(&x); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var uintID uint
		if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		x.ID = uintID
		if err := exampleService.UpdateExample(&x); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, x)
	})

	r.DELETE("/example/:id", func(c *gin.Context) {

		id := c.Param("id")
		var uintID uint
		if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if err := exampleService.DeleteExample(uintID); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})

	r.GET("/example/:id", func(c *gin.Context) {
		id := c.Param("id")
		var uintID uint
		if _, err := fmt.Sscanf(id, "%d", &uintID); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		x, err := exampleService.GetExampleByID(uintID)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "example not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, x)
	})

	r.GET("/example", func(c *gin.Context) {
		x, err := exampleService.GetExamples()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, x)
	})

	r.Run(":" + cfg.ExamplePort)
}
