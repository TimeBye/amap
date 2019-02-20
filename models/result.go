package models

type Poi struct {
	ID           string      `json:"id"`
	Parent       interface{} `json:"parent"`
	Name         string      `json:"name"`
	Tag          interface{} `json:"tag"`
	Type         string      `json:"type"`
	Typecode     string      `json:"typecode"`
	BizType      interface{} `json:"biz_type"`
	Address      interface{} `json:"address"`
	Location     string      `json:"location"`
	Tel          interface{} `json:"tel"`
	Postcode     interface{} `json:"postcode"`
	Website      interface{} `json:"website"`
	Email        interface{} `json:"email"`
	Pcode        string      `json:"pcode"`
	Pname        string      `json:"pname"`
	Citycode     string      `json:"citycode"`
	Cityname     string      `json:"cityname"`
	Adcode       string      `json:"adcode"`
	Adname       string      `json:"adname"`
	Importance   interface{} `json:"importance"`
	Shopid       interface{} `json:"shopid"`
	Shopinfo     string      `json:"shopinfo"`
	Poiweight    interface{} `json:"poiweight"`
	Gridcode     string      `json:"gridcode"`
	Distance     string      `json:"distance"`
	NaviPoiid    interface{} `json:"navi_poiid"`
	EntrLocation interface{} `json:"entr_location"`
	BusinessArea interface{} `json:"business_area"`
	ExitLocation interface{} `json:"exit_location"`
	Match        string      `json:"match"`
	Recommend    string      `json:"recommend"`
	Timestamp    interface{} `json:"timestamp"`
	Alias        interface{} `json:"alias"`
	IndoorMap    string      `json:"indoor_map"`
	IndoorData   struct {
		Cpid      interface{} `json:"cpid"`
		Floor     interface{} `json:"floor"`
		Truefloor interface{} `json:"truefloor"`
		Cmsid     interface{} `json:"cmsid"`
	} `json:"indoor_data"`
	GroupbuyNum string `json:"groupbuy_num"`
	DiscountNum string `json:"discount_num"`
	BizExt      struct {
		Rating interface{} `json:"rating"`
		Cost   interface{} `json:"cost"`
	} `json:"biz_ext"`
	Event    interface{} `json:"event"`
	Children interface{} `json:"children"`
	Photos   interface{} `json:"photos"`
}

type Result struct {
	Status     string `json:"status"`
	Count      string `json:"count"`
	Info       string `json:"info"`
	Infocode   string `json:"infocode"`
	Suggestion struct {
		Keywords interface{} `json:"keywords"`
		Cities   interface{} `json:"cities"`
	} `json:"suggestion"`
	Pois []Poi `json:"pois"`
}
