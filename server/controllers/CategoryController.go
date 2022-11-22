package controllers

import (
	"blog-admin/dto"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
}

// 添加栏目
func (receiver CategoryController) Insert(ctx *gin.Context) {
	db := initialization.GetDB()
	var category models.Category
	// 获取参数
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	// 判断title是否为空
	if title == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "title不能为空")
		return
	}
	// 查询数据库
	db.Where("title = ?", title).First(&category)
	// 判断title是否存在
	if category.ID != 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "title已存在")
		return
	}
	// 创建数据
	newData := models.Category{
		Title:       title,
		Description: description,
	}
	db.Create(&newData)
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"data": newData,
	}, "添加成功")
}

// 更新栏目
func (receiver CategoryController) Update(ctx *gin.Context) {
	db := initialization.GetDB()
	var category models.Category
	// 获取参数
	id := ctx.PostForm("id")
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	// 判断title是否为空
	if title == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "title不能为空")
		return
	}
	// 查询数据库
	tx := db.Where("id = ?", id).First(&category)
	// 判断id是否为空
	if category.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目id不存在")
		return
	}
	// 判断数据库有没有同名栏目
	var categoryTitle models.Category
	db.Where("title = ?", title).First(&categoryTitle)
	if categoryTitle.Title == title {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目title已存在")
		return
	}
	// 更新数据
	tx.Updates(&models.Category{
		Title:       title,
		Description: description,
	})

	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"data": dto.ToCategoryDto(category),
	}, "更新成功")
}

// 删除栏目
func (receiver CategoryController) Delete(ctx *gin.Context) {
	db := initialization.GetDB()
	var category models.Category
	// 获取参数
	id := ctx.PostForm("id")
	// 判断id不能为空
	if id == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目id不能为空")
		return
	}
	// 查询id是否存在
	tx := db.Where("id = ?", id).First(&category)
	if category.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目id不存在")
		return
	}
	tx.Unscoped().Delete(&category)
	tools.Response(ctx, http.StatusOK, 200, nil, "删除成功")
}

// 查询栏目
func (receiver CategoryController) Query(ctx *gin.Context) {
	db := initialization.GetDB()
	var categorySlice []models.Category
	var category models.Category
	// 获取参数
	id, b := ctx.GetQuery("id")
	// 判断是否有id
	if b == false {
		db.Where("1 = 1").Find(&categorySlice)
		tools.Response(ctx, http.StatusOK, 200, gin.H{
			"list": categorySlice,
		}, "获取成功")
	} else {
		db.Where("id = ?", id).First(&category)
		tools.Response(ctx, http.StatusOK, 200, gin.H{
			"data": dto.ToCategoryDto(category),
		}, "获取成功")
	}
}
