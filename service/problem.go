package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oj/define"
	"oj/models"
	"strconv"
)

// GetProblemList
// @Tags         公共方法
// @Summary      问题列表
// @Param        keyword   query   string  false  "keyword"
// @Param        page   query      int  false  "请输入当前页,默认第一页"
// @Param        size   query      int  false  "size"
// @Param        category_identity   query    string  false  "分类的唯一标识"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /problem_list [get]
func GetProblemList(context *gin.Context) {
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
	category_identity := context.Query("category_identity")
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, category_identity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Panicln("get problem list:", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})

}

// GetProblemDetail
// @Tags         公共方法
// @Summary      问题详情
// @Param        identity   query   string  false  "problem identity"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /problem_detail [get]
func GetProblemDetail(context *gin.Context) {
	identity := context.Query("identity")
	if identity == "" {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity=?", identity).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
		}
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetProblemDetail ERROR" + err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
