package ro

//Status response
type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

//Province response
type Province struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

//City response
type City struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

//CostWrapper
type CostWrapper struct {
	Rajaongkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status             Status     `json:"status"`
		OriginDetails      City       `json:"origin_details"`
		DestinationDetails City       `json:"destination_details"`
		Results            []Provider `json:"results"`
	} `json:"rajaongkir"`
}

//Cost
type Cost struct {
	OriginDetails      City       `json:"origin_details"`
	DestinationDetails City       `json:"destination_details"`
	Providers          []Provider `json:"results"`
}

//Provider response
type Provider struct {
	Code  string        `json:"code"`
	Name  string        `json:"name"`
	Costs []ServiceCost `json:"costs"`
}

//ServiceCost response
type ServiceCost struct {
	Service     string              `json:"service"`
	Description string              `json:"description"`
	Cost        []ServiceCostDetail `json:"cost"`
}

//ServiceCostDetail response
type ServiceCostDetail struct {
	Value        int    `json:"value"`
	EstimatedDay string `json:"etd"`
	Note         string `json:"note"`
}
