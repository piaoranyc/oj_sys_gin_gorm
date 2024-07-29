package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"oj/define"
	"oj/help"
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

// ProblemCreate
// @Tags         管理员私有方法
// @Summary      问题创建
// @Param        authorization   header   string  true  "authorization"
// @Param        title   formData   string  true  "title"
// @Param        content   formData   string  true  "content"
// @Param        max_runtime   formData   int  false  "max_runtime"
// @Param        max_mem   formData   int  false  "max_mem"
// @Param        category_ids   formData   array   false  "category_ids"
// @Param        test_cases   formData   array  true  "test_cases"
// @Success      200  string json "{"code":"200","msg":,"",data:""}"
// @Router       /problem-create [post]
func ProblemCreate(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	max_runtime, _ := strconv.Atoi(context.PostForm("max_runtime"))
	max_mem, _ := strconv.Atoi(context.PostForm("max_mem"))
	categoryIds := context.PostFormArray("category_ids")
	test_cases := context.PostFormArray("test_cases")

	//err := context.ShouldBind(&in)
	//if err != nil {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  err.Error() + "参数错误",
	//	})
	//	return
	//}
	//log.Println(in)
	if title == "" || content == "" || len(categoryIds) == 0 || len(test_cases) == 0 || max_runtime == 0 || max_mem == 0 {

		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := help.GetUUID()
	data := &models.ProblemBasic{
		Identity:   identity,
		Title:      title,
		Content:    content,
		MaxRuntime: max_runtime,
		MaxMem:     max_mem,
	}

	categoryBasics := make([]*models.ProblemCategory, 0)

	for _, id := range categoryIds {
		categoryId, _ := strconv.Atoi(id)
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(categoryId),
		})

	}
	data.ProblemCategories = categoryBasics
	testCaseBasics := make([]*models.TestCase, 0)
	for _, v := range test_cases {
		caseMap := make(map[string]string)
		err := json.Unmarshal([]byte(v), &caseMap)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "json Unmarshal ERROR" + err.Error(),
			})
		}
		testCaseBasic := &models.TestCase{
			Identity:        help.GetUUID(),
			ProblemIdentity: identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics

	err := models.DB.Create(data).Error

	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Create ERROR" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})

}
