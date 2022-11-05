package controllers

import (
	"blog-admin/dto"
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleController struct {
}

// 添加文章
func (receiver ArticleController) Insert(ctx *gin.Context) {
	db := initialization.GetDB()
	var category models.Category
	// 获取参数
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	categoryId := ctx.PostForm("category_id")
	// 判断title是否为空
	if title == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "title不能为空")
		return
	}
	// 类型转换
	intCategoryId, _ := strconv.Atoi(categoryId)
	uintCategoryId := uint(intCategoryId)
	// 查询数据库category_id是否存在
	db.Where("id = ?", categoryId).First(&category)
	if category.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目id不存在")
		return
	}
	// 创建数据
	newData := models.Article{
		Title:      title,
		Content:    content,
		CategoryId: uintCategoryId,
	}
	db.Create(&newData)
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"data": newData,
	}, "添加成功")
}

// 更新文章
func (receiver ArticleController) Update(ctx *gin.Context) {
	db := initialization.GetDB()
	var article models.Article
	var category models.Category
	// 获取参数
	id := ctx.PostForm("id")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	categoryId := ctx.PostForm("category_id")

	// 判断id是否为空
	if id == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "id不能为空")
		return
	}
	if categoryId == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "category_id不能为空")
		return
	}
	// 判断文章id是否存在
	tx := db.Where("id = ?", id).First(&article)
	if article.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "文章id不存在")
		return
	}
	// 判断栏目id是否存在
	db.Where("id = ?", categoryId).First(&category)
	if category.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "栏目id不存在")
		return
	}
	// 类型转换
	intCategoryId, _ := strconv.Atoi(categoryId)
	uintCategoryId := uint(intCategoryId)
	// 更新数据
	tx.Updates(&models.Article{
		Title:      title,
		Content:    content,
		CategoryId: uintCategoryId,
	})

	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"data": dto.ToArticleDto(article),
	}, "更新成功")
}

// 删除文章
func (receiver ArticleController) Delete(ctx *gin.Context) {
	db := initialization.GetDB()
	var article models.Article
	// 获取参数
	id := ctx.PostForm("id")
	// 判断id不能为空
	if id == "" {
		tools.Response(ctx, http.StatusOK, 200, nil, "文章id不能为空")
		return
	}
	// 查询id是否存在
	tx := db.Where("id = ?", id).First(&article)
	if article.ID == 0 {
		tools.Response(ctx, http.StatusOK, 200, nil, "文章id不存在")
		return
	}
	tx.Unscoped().Delete(&article)
	tools.Response(ctx, http.StatusOK, 200, nil, "删除成功")
}

// 查询文章
func (receiver ArticleController) Query(ctx *gin.Context) {
	db := initialization.GetDB()
	var article []models.Article
	var articleId models.Article
	// 获取参数
	categoryId, cb := ctx.GetQuery("category_id")
	id, ib := ctx.GetQuery("id")
	// 判断id和category_id是否为空
	if cb == true {
		db.Where("category_id = ?", categoryId).Find(&article)
		tools.Response(ctx, http.StatusOK, 200, gin.H{
			"list": article,
		}, "获取成功")
		return
	}
	if ib == true {
		db.Where("id = ?", id).Find(&articleId)
		tools.Response(ctx, http.StatusOK, 200, gin.H{
			"data": articleId,
		}, "获取成功")
		return
	}
	// 为空id或category_id的返回
	db.Where("1=1").Find(&article)
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"list": article,
	}, "获取成功")
}
