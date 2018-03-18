package validators

type CompanyCreateIsValid struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Telephone string `form:"telephone" json:"telephone" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
}

type CompanyLoginIsValid struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}