package request

type IDS struct {
	ID int `form:"id" valid:"required~id不能为空,int~id格式不正确"`
}
