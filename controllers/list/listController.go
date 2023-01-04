package list

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListController struct{}

// GetBacklogList
// @Tags 简历相关
// @Summary 获取简历列表接口
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/resume [get]
func (list ListController) ResumeList(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	email := c.DefaultQuery("email", "")
	page, pageError := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageError != nil {
		fmt.Println("pageError", pageError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "page参数异常",
		})
		return
	}
	pageSize, pageSizeError := strconv.Atoi(c.DefaultQuery("pageSize", "50"))
	if pageSizeError != nil {
		fmt.Println("pageError", pageSizeError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "pageSize参数异常",
		})
		return
	}

	/*token := c.GetHeader("token")
	userinfo, _ := helper.ParseToken(c, token)*/
	userinfo, err := helper.AnalysisTokenGetUserInfo(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userinfo)
	fmt.Println("获取简历列表")
	pageNum := (page - 1) * pageSize
	var (
		resumeList []models.ResumeInterface
		errs       error
	)
	if name == "" && email == "" {
		errs = models.DB.Model(models.Resume{}).Offset(pageNum).Limit(pageSize).Find(&resumeList).Error
	} else {
		errs = models.DB.Model(models.Resume{}).Where("name= ? or email = ? ", name, email).Where("is_delete", 0).Offset(pageNum).Limit(pageSize).Find(&resumeList).Error
	}
	if errs != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "查询异常",
		})
		return
	}
	//c.String(http.StatusOK, "简历列表信息")
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    resumeList,
		"text":    "简历列表接口",
		"message": "success",
	})
}

// MainResumeList
// @Tags 简历相关
// @Summary 获取重点关注人群简历列表接口
// @Description { "page": 1, "pageSize": 10 }
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/mainResume [get]
func (list ListController) MainResumeList(c *gin.Context) {
	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageErr != nil {
		fmt.Println(pageErr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "page参数异常",
			"data":    gin.H{},
		})
		return
	}
	pageSize, pageSizeError := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	if pageSize > 50 {
		pageSize = 50
	}
	if pageSizeError != nil {
		fmt.Println(pageSizeError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "pageSize参数异常",
			"data":    gin.H{},
		})
		return
	}
	pageNumber := (page - 1) * pageSize
	var mainResumeList []models.ResumeInterface
	findErr := models.DB.Model(models.Resume{}).Where("follow", 1).Where("is_delete", 0).Offset(pageNumber).Limit(pageSize).Find(&mainResumeList).Error
	if findErr != nil {
		fmt.Println("findErr", findErr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "数据查询异常",
			"data":    gin.H{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data": gin.H{
			"data":        mainResumeList,
			"total":       len(mainResumeList),
			"currentPage": page,
			"size":        pageSize,
		},
	})
}

// ModifyMainStatus
// @Tags 简历相关
// @Summary 取消/添加重点标记 false取消 true 添加
// @Param status query string true "true/false"
// @Param id query string true "id"
// @Description {"status":0, id: 1}
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/modifyMain [post]
func (list ListController) ModifyMainStatus(c *gin.Context) {
	json := make(map[string]interface{})

	c.ShouldBindJSON(&json)
	status := json["status"]
	id := json["id"]
	if status == nil || id == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "参数不全",
			"data":    gin.H{},
		})
		return
	}
	statusBool, boolerr := strconv.ParseBool(status.(string))
	/*if _, boolerr := status.(bool); !boolerr { // 使用类型断言
		fmt.Println(boolerr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "status值不符合规范",
			"data":    gin.H{},
		})
		return
	}*/
	if boolerr != nil {
		fmt.Println(boolerr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "status值不符合规范",
			"data":    gin.H{},
		})
		return
	}
	var searchData models.Resume
	err := models.DB.Model(models.Resume{}).Where("id = ?", id).Where("is_delete", 0).First(&searchData).Error
	fmt.Println("------------", status, searchData.Follow)
	if statusBool == searchData.Follow {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "当前状态与保存状态一致，请检查参数",
			"data":    gin.H{},
		})
		return
	}
	updateError := models.DB.Model(models.Resume{}).Where("id = ?", id).Where("is_delete", 0).Update("Follow", statusBool).Error
	if updateError != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "更新数据库失败，请重试",
			"data":    gin.H{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    gin.H{},
	})
}

// AddUserResume
// @Tags 简历相关
// @Summary 添加简历和用户信息
// @Param name query string true "name"
// @Param email query string true "email"
// @Param resumeUrl query string true "resumeUrl"
// @Param phone query string false "phone"
// @Param gender query string false "gender"
// @Param employmentIntention query string false "employmentIntention"
// @Param confirmEnrollment query string false "confirmEnrollment"
// @Param jobbed query string false "jobbed"
// @Param level query string false "level"
// @Param targetCompany query string false "targetCompany"
// @Param postSalary query string false "postSalary"
// @Param timeInduction query string false "timeInduction"
// @Param firstContactTime query string false "firstContactTime"
// @Param personCharge query string false "personCharge"
// @Param remarks query string false "remarks"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/addUserResume [post]
func (list ListController) AddUserResume(c *gin.Context) {
	json := make(map[string]string)
	c.ShouldBindJSON(&json)

	name := json["name"]
	phone := json["phone"]
	email := json["email"]
	gender := json["gender"]
	employmentIntention := json["employmentIntention"] // 入职意向
	confirmEnrollment := json["confirmEnrollment"]     // 是否确认入职
	jobbed := json["jobbed"]                           // 技术岗位
	level := json["level"]                             // 级别
	targetCompany := json["targetCompany"]             // 目标公司
	postSalary := json["postSalary"]                   // 薪资
	timeInduction := json["timeInduction"]             // 入职时间
	firstContactTime := json["firstContactTime"]       // 首次联系时间
	personCharge := json["personCharge"]               // 入职负责人
	remarks := json["remarks"]                         // 备注
	resumeUrl := json["resumeUrl"]                     // 简历url
	if name == "" || email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "用户名或者邮箱不能为空",
		})
		return
	}
	if resumeUrl == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "请上传简历之后再提交",
		})
		return
	}

	resumeInfo := models.Resume{
		Name:                name,
		Phone:               phone,
		Email:               email,
		Gender:              gender,
		EmploymentIntention: employmentIntention,
		ConfirmEnrollment:   confirmEnrollment,
		Jobbed:              jobbed,
		Level:               level,
		TargetCompany:       targetCompany,
		PostSalary:          postSalary,
		TimeInduction:       timeInduction,
		FirstContactTime:    firstContactTime,
		PersonCharge:        personCharge,
		Remarks:             remarks,
		ResumeUrl:           resumeUrl,
	}
	err := models.DB.Model(models.Resume{}).Create(&resumeInfo).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "用户数据插入失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"date":    gin.H{},
	})
}

// Upload
// @Tags 公共方法
// @Summary 上传文件
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/upload [post]
func (list ListController) Upload(c *gin.Context) {
	fmt.Println("上传文件interface start")
	var obj interface{}
	//req := httplib.Post("http://192.168.1.17:8080/upload")
	req := httplib.Post("http://192.168.1.17:8080/group1")
	fmt.Println("req", req)
	//req.PostFile("file", "filename") //注意不是全路径
	//req.Param("output", "json")
	//req.Param("scene", "default")
	//req.Param("path", "")
	req.ToJSON(&obj)
	fmt.Print(obj)
	c.JSON(http.StatusOK, gin.H{
		"data": obj,
		"req":  req,
	})
	/*defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 单文件
	file, _ := c.FormFile("file")
	dir := "./file"
	os.MkdirAll(dir, 0666)

	filename := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	// 上传文件至指定的完整文件路径
	dst := path.Join(dir, filename)
	c.SaveUploadedFile(file, dst)
	returnPath := dst
	c.JSON(http.StatusOK, gin.H{
		"code":     "200",
		"data":     returnPath,
		"path":     dir,
		"filename": filename,
	})*/
}
