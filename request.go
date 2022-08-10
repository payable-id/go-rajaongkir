package ro

//QueryRequest set of query paramters
type QueryRequest struct {
	SubDistrictID string
	CityID        string
	ProvinceID    string
	Origin        string
	Destination   string
	Weight        int
	Courier       string
}
