package column

type TitleS struct {
	Title string `form:"title" valid:"required~专栏标题不能为空,stringlength(10|50)~专栏标题10-50个字符"`
}
