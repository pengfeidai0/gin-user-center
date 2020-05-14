package schema

type UpdatePassword struct {
	Phone       string `form:"phone" json:"phone" binding:"required,len=11"`
	OldPassword string `form:"oldPassword" json:"oldPassword" binding:"required,gt=0,lte=32"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required,gt=0,lte=32"`
}
