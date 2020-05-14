package schema

type Register struct {
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Name     string `form:"name" json:"name" binding:"required,gt=0,lte=32"`
	Password string `form:"password" json:"password" binding:"required,gt=0,lte=32"`
	Avatar   string `form:"avatar" json:"avatar"`
}

type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required,len=11"`
	Password string `form:"password" json:"password" binding:"required,gt=0,lte=32"`
}
