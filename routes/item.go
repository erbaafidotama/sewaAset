package routes

import (
	"strconv"

	"sewaAset/config"
	"sewaAset/models"

	"github.com/gin-gonic/gin"
)

func PostItem(c *gin.Context) {
	price := c.PostForm("price")
	floatPrice, _ := strconv.ParseFloat(price, 8)

	quantity := c.PostForm("quantity")
	intQuantity, _ := strconv.Atoi(quantity)

	items := models.Item{
		ItemName: c.PostForm("item_name"),
		Quantity: intQuantity,
		Price:    floatPrice,
		InfoDesc: c.PostForm("description"),
	}

	// crete data to db
	config.DB.Create(&items)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   items,
	})
}

func GetItems(c *gin.Context) {
	items := []models.Item{}

	if err := config.DB.Find(&items).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "GET Data Items",
		"data":    items,
	})
}

func UpdateItem(c *gin.Context) {
	itemId := c.Param("id")

	price := c.PostForm("price")
	floatPrice, _ := strconv.ParseFloat(price, 8)

	quantity := c.PostForm("quantity")
	intQuantity, _ := strconv.Atoi(quantity)

	var dataItem models.Item
	if err := config.DB.Where("id = ?", itemId).First(&dataItem).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	config.DB.Model(&dataItem).Where("id = ?", itemId).Updates(models.Item{
		ItemName: c.PostForm("item_name"),
		Quantity: intQuantity,
		Price:    floatPrice,
		InfoDesc: c.PostForm("description"),
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataItem,
	})
}

func DeleteItem(c *gin.Context) {
	// get id from url
	itemId := c.Param("id")

	var dataItem models.Item
	if err := config.DB.Where("id = ?", itemId).First(&dataItem).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", itemId).Delete(&dataItem)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   dataItem,
	})
}
