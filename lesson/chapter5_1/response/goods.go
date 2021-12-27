package response

type Goods struct {
	ID         int `json:"id"`
	Category *GoodsCategory `json:"category"`
	Name       string `json:"name"`
	CreateTime DateTime `json:"createTime"`
}

type GoodsCategory struct {
	ID         int `json:"id"`
	Name       string `json:"name"`
	CreateTime DateTime `json:"createTime"`
}
