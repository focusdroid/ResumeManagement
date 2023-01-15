package models

import (
	"gorm.io/gorm"
)

/*
简历表
*/

type Resume struct {
	gorm.Model
	UserId              string `gorm:"column:user_id;type:varchar(255);" db:"column:user_id" json:"user_id" form:"user_id" comment:"用户id默认使用邮箱"`
	Name                string `gorm:"column:name;type:varchar(50);" db:"column:name" form:"name" json:"name" comment:"简历人员姓名"`
	Gender              string `gorm:"column:gender;type:varchar(50);" db:"column:gender" form:"gender" json:"gender" comment:"简历人员性别"`
	Phone               string `gorm:"column:phone;type:varchar(11);" db:"column:phone" form:"phone" json:"phone" comment:"候选人手机号"`
	Email               string `gorm:"column:email;type:varchar(255);" db:"column:email" json:"email" form:"email" comment:"候选人邮箱"`
	UploadUserEmail     string `gorm:"column:upload_user_email;type:varchar(255);" db:"column:upload_user_email" json:"upload_user_email" form:"email" comment:"管理人员邮箱"`
	Jobbed              string `gorm:"column:jobbed;type:varchar(50);" db:"column:jobbed" form:"jobbed" json:"jobbed" comment:"岗位"`
	Level               string `gorm:"column:level;type:varchar(50);" db:"column:level" form:"level" json:"level" comment:"级别"`
	TargetCompany       string `gorm:"column:target_company;type:varchar(50);" db:"column:target_company" form:"target_company" json:"target_company" comment:"目标公司"`
	FirstContactTime    string `gorm:"column:first_contact_time;type:varchar(50);" db:"column:first_contact_time" form:"first_contact_time" json:"first_contact_time" comment:"首次联系时间"`
	EmploymentIntention string `gorm:"column:employment_intention;type:varchar(50);" db:"column:employment_intention" form:"employment_intention" json:"employment_intention" comment:"入职意向"`
	ConfirmEnrollment   string `gorm:"column:confirm_enrollment;type:varchar(50);" db:"column:confirm_enrollment" form:"confirm_enrollment" json:"confirm_enrollment" comment:"是否确认入职"`
	PostSalary          string `gorm:"column:post_salary;type:varchar(50);" db:"column:post_salary" form:"post_salary" json:"post_salary" comment:"岗位工资"`
	TimeInduction       string `gorm:"column:time_induction;type:varchar(50);" db:"column:time_induction" form:"time_induction" json:"time_induction" comment:"几号入职"`
	PersonCharge        string `gorm:"column:person_charge;type:varchar(50);" db:"column:person_charge" form:"person_charge" json:"person_charge" comment:"入职负责人"`
	Follow              bool   `gorm:"column:follow;type:tinyint;default:0;" db:"column:follow" form:"follow" json:"follow" comment:"是否重点关注"`
	ResumeUrl           string `gorm:"column:resume_url;type:varchar(256);" db:"column:resume_url" form:"resume_url" json:"resume_url" comment:"简历url"`
	ResumeStatus        string `gorm:"column:resume_status;type:varchar(20);default: 1" db:"column:resume_status" form:"resume_status" json:"resume_status" comment:"简历状态0禁止查看管理员可以看,1正常查看"`
	IsDelete            bool   `gorm:"column:is_delete;tinyint;default:0;" db:"column:is_delete" form:"is_delete" json:"is_delete" comment:"0未删除,1已删除"`
	Remarks             string `gorm:"column:remarks;type:varchar(255)" db:"column:remarks" form:"remarks" json:"remarks" comment:"备注信息"`
}

type ResumeInterface struct {
	ID                  uint   `json:"id"`
	UserId              string `json:"user_id"`
	Name                string `json:"name"`
	Gender              string `json:"gender"`
	Phone               string `json:"phone"`
	Email               string `json:"email"`
	Jobbed              string `json:"jobbed"`
	Level               string `json:"level"`
	TargetCompany       string `json:"target_company"`
	FirstContactTime    string `json:"first_contact_time"`
	EmploymentIntention string `json:"employment_intention"`
	ConfirmEnrollment   string `json:"confirm_enrollment"`
	PostSalary          string `json:"post_salary"`
	TimeInduction       string `json:"time_induction"`
	PersonCharge        string `json:"person_charge"`
	Follow              bool   `json:"follow"`
	ResumeUrl           string `json:"resume_url;"`
	ResumeStatus        string `json:"resume_status"`
	Remarks             string `json:"remarks"`
}

func (resume Resume) TableName() string {
	return "resume"
}
