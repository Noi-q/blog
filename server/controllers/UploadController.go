package controllers

import (
	"blog-admin/initialization"
	"blog-admin/models"
	"blog-admin/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

type UploadController struct {
}

func (receiver UploadController) Img(ctx *gin.Context) {
	// 获取上传文件
	file, err := ctx.FormFile("file")
	if err != nil {
		tools.Fail(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}
	// 获取文件后缀名 判断类型是否正确
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		tools.Response(ctx, http.StatusOK, 200, nil, "上传的文件类型不合法")
		return
	}
	// 创建图片保存目录 static/upload/
	day := time.Now().Format("20060102")
	dir := "./static/avatar/" + day
	errMkdir := os.MkdirAll(dir, 0666)
	if errMkdir != nil {
		tools.Response(ctx, http.StatusInternalServerError, 500, nil, "创建目录失败")
		return
	}
	// 生成文件名称和文件保存的目录
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + extName
	dst := path.Join(dir, fileName)
	// 执行上传
	errUpload := ctx.SaveUploadedFile(file, dst)
	if errUpload != nil {
		tools.Response(ctx, http.StatusOK, 200, nil, "上传失败: "+errUpload.Error())
	}
	// 更新数据库
	db := initialization.GetDB()
	var admin models.Admin
	db.Find(&admin).Updates(models.Admin{
		Avatar: dst,
	})
	tools.Response(ctx, http.StatusOK, 200, gin.H{
		"src": dst,
	}, "更新成功")
}
