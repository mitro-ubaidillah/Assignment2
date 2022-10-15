package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Tugas2/structs"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []structs.Orders
		result gin.H
	)

	idb.DB.Preload("Item").Find(&orders)
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		order  structs.Orders
		item   structs.Items
		result gin.H
	)

	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")

	order.Ordered_At = time.Now()
	order.Customer_Name = customer_name
	item.Item_Code = item_code
	item.Description = description
	item.Quantity, _ = strconv.Atoi(quantity)

	idb.DB.Create(&order)
	item.Order_Id = order.Order_Id
	idb.DB.Create(&item)

	// orders := structs.Orders{}

	err := idb.DB.Preload("Item").Find(&order).Error
	if err != nil {
		fmt.Println("gagal mendapatkan data")
		return
	}

	result = gin.H{
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")

	var (
		order    structs.Orders
		newOrder structs.Orders
		item     structs.Items
		newItem  structs.Items
		result   gin.H
	)

	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newOrder.Customer_Name = customer_name
	newItem.Item_Code = item_code
	newItem.Description = description
	newItem.Quantity, _ = strconv.Atoi(quantity)

	err = idb.DB.Model(&order).Updates(newOrder).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		errItem := idb.DB.First(&item, id).Error
		if errItem != nil {
			result = gin.H{
				"result": "data not found",
			}
		}
		err = idb.DB.Model(&item).Updates(newItem).Error
		if err != nil {
			result = gin.H{
				"result": "update Item failed",
			}
		} else {
			err = idb.DB.Preload("Item").Find(&order).Error
			result = gin.H{
				"result": order,
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		order  structs.Orders
		item   structs.Items
		result gin.H
	)

	id := c.Param("id")
	// errItem := idb.DB.First(&item, id).Error
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		errItem := idb.DB.First(&item, id).Error
		if errItem != nil {
			result = gin.H{
				"result": "data not found",
			}
		}
		err = idb.DB.Delete(&item).Error
		if err != nil {
			result = gin.H{
				"result": "delete failed",
			}
		} else {
			result = gin.H{
				"result": "data deleted successfully",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
