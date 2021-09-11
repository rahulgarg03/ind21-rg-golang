package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulgarg03/ind21-rg-golang/auth"
	"github.com/rahulgarg03/ind21-rg-golang/database"
	"github.com/rahulgarg03/ind21-rg-golang/src"

	"gorm.io/gorm"
)

func initializeRouters() {
	db := database.ConnectDataBase()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/login", auth.Login)
	r.GET("/home", auth.Home)

	r.GET("/getusers", func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		var users []src.User
		db.Find(&users)
		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	r.POST("/createuser", func(c *gin.Context) {
		DB := c.MustGet("db").(*gorm.DB)
		var input src.User
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(input)
		user := src.User{Email: input.Email, Password: input.Password,
			Active: input.Active, FirstName: input.FirstName,
			LastName: input.LastName, Age: input.Age, Gender: input.Gender,
			MartialStatus: input.MartialStatus, ResidentialAddress: input.ResidentialAddress,
			ResidentialCity: input.ResidentialCity, ResidentialState: input.ResidentialState,
			ResidentialCountry: input.ResidentialCountry, ResidentialContactNo1: input.ResidentialContactNo1,
			ResidentialContactNo2: input.ResidentialContactNo2, OfficialDetailsAddress: input.OfficialDetailsAddress,
			OfficialDetailsState: input.OfficialDetailsState, OfficialDetailsCity: input.OfficialDetailsCity,
			OfficialDetailsCountry: input.OfficialDetailsCountry, OfficialDetailsCompanyContactNo: input.OfficialDetailsCompanyContactNo,
			OfficialDetailsCompanyEmail: input.OfficialDetailsCompanyEmail, OfficialDetailsCompanyName: input.OfficialDetailsCompanyName}
		fmt.Println(user)
		DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"data": user})

	})

	r.PUT("/updateuser/:id", func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		var user src.User
		err := db.Where("id = ?", c.Param("id")).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		// Validate input
		var input src.User
		err1 := c.ShouldBindJSON(&input)
		if err1 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
			return
		}
		db.Model(&user).Updates(input)
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.DELETE("/deleteuser/:id", func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		var user src.User
		err := db.Where("id = ?", c.Param("id")).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"data": true})
	})

	r.GET("/getuser/:id", func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		var user src.User
		err := db.Where("id = ?", c.Param("id")).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	r.Run()
}

func main() {

	initializeRouters()
}
