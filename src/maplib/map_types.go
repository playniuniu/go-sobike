package maplib

// GaodeJSON struct
type GaodeJSON struct {
	Count    string     `json:"count"`
	Geocodes []GaodeGeo `json:"geocodes"`
	Info     string     `json:"info"`
	Infocode string     `json:"infocode"`
	Status   string     `json:"status"`
}

// GaodeGeo struct
type GaodeGeo struct {
	Adcode   string `json:"adcode"`
	Building struct {
		Name []interface{} `json:"name"`
		Type []interface{} `json:"type"`
	} `json:"building"`
	City             string `json:"city"`
	Citycode         string `json:"citycode"`
	District         string `json:"district"`
	FormattedAddress string `json:"formatted_address"`
	Level            string `json:"level"`
	Location         string `json:"location"`
	Neighborhood     struct {
		Name []interface{} `json:"name"`
		Type []interface{} `json:"type"`
	} `json:"neighborhood"`
	Number   []interface{} `json:"number"`
	Province string        `json:"province"`
	Street   []interface{} `json:"street"`
	Township []interface{} `json:"township"`
}
