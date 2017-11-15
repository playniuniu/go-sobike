package bikelib

import "fmt"

// OfoJSON struct is api json response
type OfoJSON struct {
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
	Values    struct {
		Info struct {
			ZoomLevel      int           `json:"zoomLevel"`
			Cars           []OfoCar      `json:"cars"`
			RedPacketAreas []interface{} `json:"redPacketAreas"`
			Time           int           `json:"time"`
		} `json:"info"`
	} `json:"values"`
}

// OfoCar struct is used for ofo car
type OfoCar struct {
	Carno      string  `json:"carno"`
	Ordernum   string  `json:"ordernum"`
	UserIDLast string  `json:"userIdLast"`
	Lng        float64 `json:"lng"`
	Lat        float64 `json:"lat"`
}

func (el OfoCar) String() string {
	return fmt.Sprintf("Ofo Num: %v Pos: ( %v, %v )", el.Carno, el.Lng, el.Lat)
}
