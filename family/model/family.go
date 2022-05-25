package model

type FamilyTreeBase struct {
	CartId        int     `json:"cart_id"`
	ID            int     `json:"id"`
	Uuid          string  `json:"uuid"`
	AwemeId       string  `json:"awemeId" gorm:"column:aweme_id"`
	GoodsId       string  `json:"goodsId" gorm:"column:goods_id"`
	Title         string  `json:"title"`
	Platform      string  `json:"platform"`
	Subtitle      string  `json:"subtitle"`
	Url           string  `json:"url"`
	ImgUrl        string  `json:"imgUrl" gorm:"column:img_url"`
	OriginalPrice float64 `json:"originalPrice" gorm:"column:original_price"`
	VisitNum      int     `json:"visitNum" gorm:"column:visit_num"`
	MinPrice      float64 `json:"minPrice" gorm:"column:min_price"`
	MaxPrice      float64 `json:"maxPrice" gorm:"column:max_price"`
	Sales         int     `json:"sales"`
	ShopId        string  `json:"shopId" gorm:"column:shop_id"`
	ShopName      string  `json:"shopName" gorm:"column:shop_name"`
	BrandName     string  `json:"brandName" gorm:"column:brand_name"`
	FinalPrice    float64 `json:"finalPrice" gorm:"column:final_price"`
	TbCouponInfo  string  `json:"tbCouponInfo" gorm:"column:tb_coupon_info"`
}

func (FamilyTreeBase) TableName() string {
	return "dy_cart_goods"
}
