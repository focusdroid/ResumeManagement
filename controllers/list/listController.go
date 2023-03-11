package list

import (
	"ResumeManagement/helper"
	"ResumeManagement/middleware"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListController struct{}

// GetBacklogList
// @Tags 简历方法
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
	userinfo, err := helper.AnalysisTokenGetUserInfo(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userinfo)
	fmt.Println("获取简历列表")
	pageNum := (page - 1) * pageSize
	var (
		resumeList []models.ResumeInterface
		total      []models.ResumeInterface
		errs       error
	)
	fmt.Println("middleware.CurrentEmail", middleware.CurrentEmail)
	if name == "" && email == "" {
		errs = models.DB.Model(models.Resume{}).Where("is_delete", 0).Where("upload_user_email = ?", middleware.CurrentEmail).Offset(pageNum).Limit(pageSize).Find(&resumeList).Error
		errs = models.DB.Model(models.Resume{}).Where("is_delete", 0).Where("upload_user_email = ?", middleware.CurrentEmail).Offset(pageNum).Limit(pageSize).Find(&total).Error
	} else {
		errs = models.DB.Model(models.Resume{}).Where("is_delete", 0).Where("upload_user_email = ?", middleware.CurrentEmail).Where("name= ? or email = ? ", name, email).Offset(pageNum).Limit(pageSize).Find(&resumeList).Error
		errs = models.DB.Model(models.Resume{}).Where("is_delete", 0).Where("upload_user_email = ?", middleware.CurrentEmail).Where("name= ? or email = ? ", name, email).Offset(pageNum).Limit(pageSize).Find(&total).Error
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
		"code": "200",
		"data": gin.H{
			"data":     resumeList,
			"total":    len(total),
			"page":     page,
			"pageSize": pageSize,
		},
		"text":    "简历列表接口",
		"message": "success",
	})
}

// MainResumeList
// @Tags 简历方法
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
	findErr := models.DB.Model(models.Resume{}).Where("is_delete", 0).Where("upload_user_email = ?", middleware.CurrentEmail).Where("follow", 1).Offset(pageNumber).Limit(pageSize).Find(&mainResumeList).Error
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
// @Tags 简历方法
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
	fmt.Println("------------", status, statusBool, searchData.Follow)
	if statusBool == searchData.Follow {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "当前状态与保存状态一致，请检查参数",
			"data":    gin.H{},
		})
		return
	}
	updateError := models.DB.Model(models.Resume{}).Where("id = ?", id).Where("upload_user_email = ?", middleware.CurrentEmail).Where("is_delete", 0).Update("Follow", statusBool).Error
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
// @Tags 简历方法
// @Summary 添加简历
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
// @Description /list/addUserResume
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/addUserResume [post]
func (list ListController) AddUserResume(c *gin.Context) {
	json := map[string]string{
		"name":                 "",
		"phone":                "",
		"email":                "",
		"gender":               "",
		"employment_intention": "",
		"confirm_enrollment":   "",
		"jobbed":               "",
		"level":                "",
		"target_company":       "",
		"post_salary":          "",
		"time_induction":       "",
		"first_contact_time":   "",
		"person_charge":        "",
		"remarks":              "",
		"resumeUrl":            "",
	}
	c.ShouldBindJSON(&json)

	name := json["name"]
	phone := json["phone"]
	email := json["email"]
	gender := json["gender"]
	employmentIntention := json["employment_intention"] // 入职意向
	confirmEnrollment := json["confirm_enrollment"]     // 是否确认入职
	jobbed := json["jobbed"]                            // 技术岗位
	level := json["level"]                              // 级别
	targetCompany := json["target_company"]             // 目标公司
	postSalary := json["post_salary"]                   // 薪资
	timeInduction := json["time_induction"]             // 入职时间
	firstContactTime := json["first_contact_time"]      // 首次联系时间
	personCharge := json["person_charge"]               // 入职负责人
	remarks := json["remarks"]                          // 备注
	resumeUrl := json["resumeUrl"]                      // 简历url
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
	token := c.GetHeader("token")
	userinfo, infoerr := helper.ParseToken(c, token)
	if infoerr != nil {
		fmt.Println(infoerr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "获取信息失败，请重新上传",
		})
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
		UploadUserEmail:     userinfo.Email,
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
// @Description /list/upload
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/upload [post]
func (list ListController) Upload(c *gin.Context) {
	fmt.Println("上传文件interface start")
	//var obj interface{}
	////req := httplib.Post("http://192.168.1.17:8080/upload")
	//req := httplib.Post("http://www.asmie.live:8080")
	//fmt.Println("req", req)
	////req.PostFile("file", "upload") //注意不是全路径
	//req.Param("output", "json")
	//req.Param("scene", "default")
	//req.Param("path", "/group/upload")
	//req.ToJSON(&obj)
	//fmt.Print(obj)
	//c.JSON(http.StatusOK, gin.H{
	//	"data": obj,
	//	"req":  req,
	//})
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

// ResumeDetail
// @Tags 简历方法
// @Summary 简历详情
// @Param id query int true "id"
// @Description { id: 1}
// @Description url: /list/detail
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/detail [get]
func (list ListController) ResumeDetail(c *gin.Context) {
	var (
		id      interface{}
		idError error
	)
	id, idError = strconv.Atoi(c.Query("id"))
	if idError != nil {
		helper.AbnormalEncapsulation(c, "id参数异常")
		return
	}
	_, ok := id.(int)
	if !ok {
		helper.AbnormalEncapsulation(c, "id参数异常")
		return
	}
	//fmt.Println(reflect.TypeOf(id))
	var resumeDetail models.ResumeInterface
	findError := models.DB.Model(&models.Resume{}).Where("id = ?", id).Where("is_delete = ?", 0).First(&resumeDetail).Error
	if findError != nil {
		helper.AbnormalEncapsulation(c, "查询数据异常或id不存在")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    resumeDetail,
		"message": "success",
	})
	return
}

// ResumeDelete
// @Tags 简历方法
// @Summary 删除简历
// @Param id query int true "id"
// @Description { id: 1}
// @Description url: /list/detail
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/delete [get]
func (list ListController) ResumeDelete(c *gin.Context) {
	var (
		id      interface{}
		idError error
	)
	id, idError = strconv.Atoi(c.Query("id"))
	if idError != nil {
		helper.AbnormalEncapsulation(c, "id参数异常")
		return
	}
	_, ok := id.(int)
	if !ok {
		helper.AbnormalEncapsulation(c, "id参数异常")
		return
	}
	findError := models.DB.Model(&models.Resume{}).Where("id = ? and is_delete = ?", id, 0).Update("is_delete", 1).Error
	if findError != nil {
		helper.AbnormalEncapsulation(c, "查询数据异常或id不存在")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    gin.H{},
		"message": "success",
	})
	return
}

// ResumeDeleted
// @Tags 管理员(admin)方法
// @Summary 获取已经删除的个人信息
// @Description
// @Description url: /list/deleted
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/deleted [get]
func (list ListController) ResumeDeleted(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    gin.H{},
		"message": "success",
	})
	return
}

// UpdateResumeInfo
// @Tags 简历方法
// @Summary 修改简历库个人信息
// @Description
// @Description url: /list/updateInfo
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /list/updateInfo [post]
func (list ListController) UpdateResumeInfo(c *gin.Context) {
	json := map[string]string{
		"name":                 "",
		"phone":                "",
		"email":                "",
		"gender":               "",
		"employment_intention": "",
		"confirm_enrollment":   "",
		"jobbed":               "",
		"level":                "",
		"target_company":       "",
		"post_salary":          "",
		"time_induction":       "",
		"first_contact_time":   "",
		"person_charge":        "",
		"remarks":              "",
		"resumeUrl":            "",
	}
	c.ShouldBindJSON(&json)
	id := json["id"]
	name := json["name"]
	phone := json["phone"]
	email := json["email"]
	gender := json["gender"]
	employmentIntention := json["employment_intention"] // 入职意向
	confirmEnrollment := json["confirm_enrollment"]     // 是否确认入职
	jobbed := json["jobbed"]                            // 技术岗位
	level := json["level"]                              // 级别
	targetCompany := json["target_company"]             // 目标公司
	postSalary := json["post_salary"]                   // 薪资
	timeInduction := json["time_induction"]             // 入职时间
	firstContactTime := json["first_contact_time"]      // 首次联系时间
	personCharge := json["person_charge"]               // 入职负责人
	remarks := json["remarks"]                          // 备注
	resumeUrl := json["resumeUrl"]                      // 简历url
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
	userinfo, infoErr := helper.AnalysisTokenGetUserInfo(c)
	if infoErr != nil {
		fmt.Println(infoErr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "获取信息失败，请重新上传",
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
		UploadUserEmail:     userinfo.Email,
	}
	ids, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"data":    gin.H{},
			"message": "id类型异常error",
			"text":    err,
		})
		return
	}
	updateError := models.DB.Model(&models.Resume{}).Where("is_delete = ?", 0).Where("id = ?", ids).Updates(resumeInfo).Error
	if updateError != nil {
		fmt.Println("updateError", updateError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"data":    gin.H{},
			"message": "update error",
			"text":    updateError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    gin.H{},
		"message": "success",
	})
}
