package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/cake-gorm/model"
	"github.com/sidz111/cake-gorm/service"
)

type CakeController struct {
	serv service.CakeService
}

func NewCakeController(s service.CakeService) *CakeController {
	return &CakeController{serv: s}
}

func (c *CakeController) Create(ctx *gin.Context) {
	var cake model.Cake
	if err := ctx.ShouldBindJSON(&cake); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data",
		})
		return
	}
	if err := c.serv.Create(ctx.Request.Context(), &cake); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"cake": cake,
	})
}
func (c *CakeController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	cake, err := c.serv.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"cake": cake,
	})
}
func (c *CakeController) GetAll(ctx *gin.Context) {
	cakes, err := c.serv.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"cakes": cakes,
	})
}
func (c *CakeController) Update(ctx *gin.Context) {
	var cake model.Cake
	if err := ctx.ShouldBindJSON(&cake); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.serv.Update(ctx.Request.Context(), &cake); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "updation done",
		"new_cake": cake,
	})
}
func (c *CakeController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	founded_cake, err := c.serv.GetById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	if founded_cake.ID != id {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "cake Not found",
		})
		return
	}
	if err := c.serv.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deletion done",
	})
}
