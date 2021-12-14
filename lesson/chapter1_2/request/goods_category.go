package request

type GoodsCategory struct {
	ID  int `form:"id" valid:"required~id不能为空,int~id格式不正确"`
	Name string `form:"mobile" valid:"required~分类名称不能为空,stringlength(2|10)~分类名称2-10个字符"`
	Desc string `form:"desc" valid:"required~分类介绍不能为空"`
}