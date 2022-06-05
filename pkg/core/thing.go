package core

type ThingBase struct {
	Name              string  `json:"name" binding:"required"`
	Type              string  `json:"type" binding:"required"`
	Remainder         float32 `json:"remainder"`
	RemainderType     string  `json:"remainder_type"`
	IsBlocked         bool    `json:"is_blocked"`
	BlockedTimeStart  int64   `json:"blocked_time_start"`
	BlockedTimeEnd    int64   `json:"blocked_time_end"`
	NeedAdminApproval bool    `json:"need_admin_approval"`
}

type Thing struct {
	ThingBase
	Id           int    `json:"id"`
	ImageURL     string `json:"image_url"`
	CompanyId    int    `json:"company_id"`
	DepartmentId int    `json:"department_id"`
}
