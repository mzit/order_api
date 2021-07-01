package models

type OrderOrigin struct {
	Model
	Data     string `json:"data"`
	Md5      string `json:"md5"`
	OrderSn  string `json:"order_sn"`
	StoreId  int    `json:"store_id"`
	Platform int    `json:"platform"`
}

func (o OrderOrigin) TableName() string {
	return "order_origin"
}

func GetOrders(offset int, pageSize int, where interface{}) (orders []OrderOrigin) {
	db.Where(where).Offset(offset).Limit(pageSize).Find(&orders)

	return
}

func GetOrderTotal(where interface{}) (count int64) {
	db.Model(&OrderOrigin{}).Where(where).Count(&count)

	return
}

func AddOrder(orderSn string, data string, storeId int, platform int, md5 string) bool {
	db.Create(&OrderOrigin{
		Data:     data,
		Md5:      md5,
		OrderSn:  orderSn,
		StoreId:  storeId,
		Platform: platform,
	})
	return true
}
