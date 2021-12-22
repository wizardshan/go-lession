package article

type TitleS struct {
	Title string `form:"title" valid:"required~文章标题不能为空,stringlength(10|50)~文章标题10-50个字符"`
}
