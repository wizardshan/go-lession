package goodscategory

type Name struct {
	Val string `form:"name" valid:"required~分类名称不能为空,stringlength(2|10)~分类名称2-10个字符"`
}

func (req *Name) Name() string {
	return req.Val
}

type Desc struct {
	Val string `form:"desc" valid:"required~分类介绍不能为空"`
}

func (req *Name) Desc() string {
	return req.Val
}