package models

type OrderGoodsOrigin struct {
	Model
	Data     string `json:"data"`
	Md5      string `json:"md5"`
	OrderSn  string `json:"order_sn"`
	StoreId  int    `json:"store_id"`
	Platform int    `json:"platform"`
}

func (og OrderGoodsOrigin) TableName() string {
	return "order_goods_origin"
}

func GetOrderGoodsByOrderSn(orderSn string) (og OrderGoodsOrigin) {
	db.Where("order_sn = ?", orderSn).First(&og)

	return
}

func ExistOrderGoodsByOrderSn(orderSn string) bool {
	var orderGoods OrderGoodsOrigin
	db.Select("id").Where("order_sn = ?", orderSn).First(&orderGoods)

	if orderGoods.ID > 0 {
		return true
	}

	return false
}
