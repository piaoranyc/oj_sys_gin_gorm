package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oj/define"
	"oj/help"
	"oj/models"
	"strconv"
)

// GetCategoryList
// @Tags         管理员私有方法
// @Summary      分类列表
// @Param        authorization   header   string  true  "authorization"
// @Param        keyword   query   string  false  "keyword"
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /admin/category_list [get]
func GetCategoryList(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println(err)
		return
	}
	size, err := strconv.Atoi(context.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println(err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := context.Query("keyword")
	categoryList := make([]*models.CategoryBasic, 0)
	err = models.DB.Model(new(models.CategoryBasic)).Where("name like ?", "%"+keyword+"%").
		Count(&count).Limit(size).Offset(page).Find(&categoryList).Error
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取分类列表失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  200,
		"list":  categoryList,
		"count": count,
	})
}

// CategoryCreate
// @Tags         管理员私有方法
// @Summary      问题创建
// @Param        authorization   header   string  true  "authorization"
// @Param        name   formData   string  true  "name"
// @Param        parentId   formData   string  false  "parentId"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /admin/category_create [post]
func GetCategoryCreate(context *gin.Context) {
	name := context.PostForm("name")
	parentId, _ := strconv.Atoi(context.PostForm("parentId"))
	category := &models.CategoryBasic{
		Identity: help.GetUUID(),
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Create(&category).Error
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error() + "创建分类失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "创建成功",
	})
}

// CategoryUpdate
// @Tags         管理员私有方法
// @Summary      问题修改
// @Param        authorization   header   string  true  "authorization"
// @Param        identity  formData   string  true  "identity"
// @Param        name   formData   string  true  "name"
// @Param        parentId   formData   string  false  "parentId"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /admin/category_modify [put]
func GetCategoryModify(context *gin.Context) {
	name := context.PostForm("name")
	parentId, _ := strconv.Atoi(context.PostForm("parentId"))
	identity := context.PostForm("identity")
	if name == "" || parentId == 0 || identity == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "不能为空",
		})
		return
	}
	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("id=?", category.Identity).Updates(&category).Error
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "修改失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// CategoryDelete
// @Tags         管理员私有方法
// @Summary      分类删除
// @Param        authorization   header   string  true  "authorization"
// @Param        identity  query string true  "identity"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /admin/category_delete [delete]
func GetCategoryDelete(context *gin.Context) {
	println(context)
	identity := context.Query("identity")
	if identity == "" {
		context.JSON(http.StatusOK, gin.H{
			"code":     -1,
			"msg":      "参数不正确3",
			"identity": identity,
		})
		return
	}
	var cnt int64

	err := models.DB.Model(new(models.ProblemCategory)).Where("category_id = (SELECT id FROM category_basic WHERE identity = ? LIMIT 1)", identity).Count(&cnt).Error

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确2",
		})
		return
	}
	if cnt > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该分类下面存在问题",
		})
		return
	}
	err = models.DB.Where("identity = ?", identity).Delete(new(models.CategoryBasic)).Error
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
