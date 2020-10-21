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
