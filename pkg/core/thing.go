package core

type ThingBase struct {
	Name              string  `json:"name" binding:"required"`
	Type              string  `json:"type" binding:"required"`
	Remainder         float32 `json:"remainder"`
	RemainderType     string  `json:"remainder_type"`
	NeedAdminApproval bool    `json:"need_admin_approval"`
}

type Thing struct {
	ThingBase
	Id           int    `json:"id"`
	ImageURL     string `json:"image_url"`
	CompanyId    int    `json:"company_id"`
	DepartmentId int    `json:"department_id"`
}

type ThingActionBase struct {
	UserId    int    `json:"user_id" binding:"required"`
	ThingId   int    `json:"thing_id" binding:"required"`
	StartTime uint32 `json:"start_time" binding:"required"`
	EndTime   uint32 `json:"end_time"`
}

type ThingBlockAdd struct {
	ThingActionBase
	Reason string `json:"reason" binding:"required"`
}

type ThingBlock struct {
	ThingBlockAdd
	Id int `json:"id"`
}

type ThingUsageAdd struct {
	ThingActionBase
}

type ThingUsage struct {
	ThingUsageAdd
	Id         int  `json:"id"`
	IsApproved bool `json:"is_approved"`
	IsTaken    bool `json:"is_taken"`
}
