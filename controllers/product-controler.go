package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	myerror "main/error"
	"main/models"
	"net/http"
	"strconv"
)

func GetAllProduct(ctx *gin.Context) {
	var product []models.Product
	db.Find(&product)
	ctx.JSON(200, gin.H{
		"message": product,
	})
}

func GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		myerror.Errors(ctx, err, "Invalid product ID", http.StatusBadRequest)
		return
	}
	fmt.Println(idStr)
	var product models.Product
	err = db.Where("id = ?", id).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			myerror.Errors(ctx, err, "Product not found", http.StatusNotFound)
		} else {
			myerror.Errors(ctx, err, "Can't find product ID", http.StatusBadRequest)
		}
		return
	}

	ctx.JSON(200, gin.H{
		"message": product,
	})
}

func SearchProduct(ctx *gin.Context) {
	searchItem := ctx.Query("product")
	var products []models.Product
	err := db.Where("name ILIKE ? OR description ILIKE ? OR category ILIKE ? OR brand ILIKE ?",
		"%"+searchItem+"%", "%"+searchItem+"%", "%"+searchItem+"%", "%"+searchItem+"%").Find(&products).Error
	if err != nil {
		myerror.Errors(ctx, err, "cant find item", 401)
		return
	}

	ctx.JSON(200, gin.H{
		"meesage": products,
	})
}

func FilterProduct(ctx *gin.Context) {
	var filter models.Filter
	err := ctx.BindQuery(&filter)
	if err != nil {
		myerror.Errors(ctx, err, "cant find product", http.StatusBadRequest)
		return
	}
	var products []models.Product
	query := db.Model(&models.Product{})

	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if filter.IsAvailable != nil {
		query = query.Where("is_available = ?", *filter.IsAvailable)
	}
	if filter.Category != "" {
		query = query.Where("category ILIKE ?", "%"+filter.Category+"%")
	}
	if filter.Brand != "" {
		query = query.Where("brand ILIKE ?", "%"+filter.Brand+"%")
	}
	if err := query.Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// When you initialize your DB connection
