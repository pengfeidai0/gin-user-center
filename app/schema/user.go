package schema

type AddUser struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Avatar   string `form:"avatar" json:"avatar"`
}
