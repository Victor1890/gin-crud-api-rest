package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Victor1890/gin-crud-api-rest/src/database"
	"github.com/Victor1890/gin-crud-api-rest/src/database/models"
	"github.com/Victor1890/gin-crud-api-rest/src/helpers"
	"github.com/gin-gonic/gin"
)

func GetAllProductHandler(ctx *gin.Context) {

	var products []models.Product

	database.DB.Model(&products)

	ctx.JSON(http.StatusOK, &products)
}

func GetOneProductHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	var product models.Product

	database.DB.First(&product, id).Association("files").Find(&product.Files)

	if product.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "The product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, &product)
}

func CreateProductHandler(ctx *gin.Context) {

	size, sizeError := strconv.ParseUint(ctx.PostForm("size"), 64, 64)
	if sizeError != nil {
		size = 0
	}

	price, priceError := strconv.ParseFloat(ctx.PostForm("price"), 64)

	if priceError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "The price is required",
		})
		return
	}

	var product models.Product = models.Product{
		Name:  ctx.PostForm("name"),
		Brand: ctx.PostForm("brand"),
		Price: price,
		Size:  size,
	}

	defer database.DB.Create(&product)

	image, formImgError := ctx.FormFile("image")

	if formImgError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "The image is required",
		})
		return
	}

	if err := helpers.MakeDir("./assets/upload/" + product.Name + "/"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	imageError := ctx.SaveUploadedFile(image, "./assets/upload/"+product.Name+"/"+image.Filename)

	if imageError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "unknown error",
		})
		return
	}

	fileName := strings.Split(image.Filename, ".")

	var file = models.File{
		Name:      fileName[0],
		Extension: fileName[1],
		MimeType:  image.Header.Get("Content-Type"),
		Size:      image.Size,
		Url:       "./assets/upload/" + image.Filename,
		ProductId: product.ID,
	}

	defer database.DB.Create(&file)

	ctx.JSON(http.StatusOK, &product)
}

func UpdateOneProductHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	size, sizeError := strconv.ParseUint(ctx.PostForm("size"), 64, 64)
	if sizeError != nil {
		size = 0
	}

	price, priceError := strconv.ParseFloat(ctx.PostForm("price"), 64)

	if priceError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "The price is required",
		})
		return
	}

	var product models.Product

	productUpdated := models.Product{
		Name:  ctx.PostForm("name"),
		Brand: ctx.PostForm("brand"),
		Price: price,
		Size:  size,
	}

	database.DB.First(&product, id)

	if product.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "product not found",
		})
		return
	}

	if productUpdated.Name != "" {
		database.DB.Model(&product).Update("name", productUpdated.Name)
	}
	if productUpdated.Brand != "" {
		database.DB.Model(&product).Update("brand", productUpdated.Brand)
	}
	if productUpdated.Size != 0 {
		database.DB.Model(&product).Update("size", productUpdated.Size)
	}
	if productUpdated.Price != 0 {
		database.DB.Model(&product).Update("price", productUpdated.Price)
	}

	ctx.JSON(http.StatusOK, &productUpdated)

}

func DeleteOneProductHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	var product models.Product

	database.DB.First(&product, id)

	if product.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "product not found",
		})
		return
	}

	database.DB.Delete(&product)

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "product is remove",
	})
}
