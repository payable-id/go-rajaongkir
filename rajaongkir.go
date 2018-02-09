package ro

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//rajaOngkir entry struct for Raja Ongkir
type rajaOngkir struct {
	apiKey string
	client *httpRequest
	Env    Env
}

// New function, create rajaOngkir pointer
// required paramter, your Raja Ongkir API KEY and timeout(http request timeout)
func New(apiKey string, timeout time.Duration) RajaOngkirService {
	httpRequest := newRequest(timeout)
	return &rajaOngkir{
		apiKey: apiKey,

		client: httpRequest,

		//set default env to starter,
		//latter just simply change with rajaOngkir.Env = Basic
		Env: Starter,
	}
}

//call function
func (r *rajaOngkir) call(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	headers["key"] = r.apiKey
	path = r.Env.String() + path
	return r.client.exec(method, path, body, v, headers)
}

//GetProvince
func (r *rajaOngkir) GetProvince(q QueryRequest) ServiceResult {

	//make headers map
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var (
		//path
		path string

		//responseWrapper single result
		responseWrapper struct {
			Rajaongkir struct {
				Query   interface{} `json:"query"`
				Status  Status      `json:"status"`
				Results Province    `json:"results"`
			} `json:"rajaongkir"`
		}

		//responseWrapperList list result
		responseWrapperList struct {
			Rajaongkir struct {
				Query   interface{} `json:"query"`
				Status  Status      `json:"status"`
				Results []Province  `json:"results"`
			} `json:"rajaongkir"`
		}
	)

	if q.ProvinceID != "" {
		path = fmt.Sprintf("province?id=%s", q.ProvinceID)

		err := r.call("GET", path, nil, &responseWrapper, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapper.Rajaongkir.Results}
	} else {
		path = "province"

		err := r.call("GET", path, nil, &responseWrapperList, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapperList.Rajaongkir.Results}
	}
	return ServiceResult{Error: errors.New("Invalid Request")}
}

//GetCity
func (r *rajaOngkir) GetCity(q QueryRequest) ServiceResult {

	//make headers map
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	var (
		//path
		path string

		//responseWrapper single result
		responseWrapper struct {
			Rajaongkir struct {
				Query   interface{} `json:"query"`
				Status  Status      `json:"status"`
				Results City        `json:"results"`
			} `json:"rajaongkir"`
		}

		//responseWrapperList list result
		responseWrapperList struct {
			Rajaongkir struct {
				Query   interface{} `json:"query"`
				Status  Status      `json:"status"`
				Results []City      `json:"results"`
			} `json:"rajaongkir"`
		}
	)

	if q.CityID != "" && q.ProvinceID == "" {
		path = fmt.Sprintf("city?id=%s", q.CityID)

		err := r.call("GET", path, nil, &responseWrapper, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapper.Rajaongkir.Results}

	} else if q.CityID != "" && q.ProvinceID != "" {
		path = fmt.Sprintf("city?id=%s&province=%s", q.CityID, q.ProvinceID)

		err := r.call("GET", path, nil, &responseWrapper, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapper.Rajaongkir.Results}
	} else if q.CityID == "" && q.ProvinceID != "" {
		path = fmt.Sprintf("city?province=%s", q.ProvinceID)

		err := r.call("GET", path, nil, &responseWrapperList, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapperList.Rajaongkir.Results}
	} else {
		path = "city"

		err := r.call("GET", path, nil, &responseWrapperList, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		return ServiceResult{Result: responseWrapperList.Rajaongkir.Results}
	}
	return ServiceResult{Error: errors.New("Invalid Request")}
}

//GetCost
func (r *rajaOngkir) GetCost(q QueryRequest) ServiceResult {

	//make headers map
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	if q.Origin != "" && q.Destination != "" && q.Weight > 0 && q.Courier != "" {

		var costWrapper CostWrapper

		query := url.Values{}
		query.Set("origin", q.Origin)
		query.Set("destination", q.Destination)

		weightString := strconv.Itoa(q.Weight)
		query.Set("weight", weightString)
		query.Set("courier", q.Courier)

		err := r.call("POST", "cost", strings.NewReader(query.Encode()), &costWrapper, headers)

		if err != nil {
			return ServiceResult{Error: err}
		}

		var cost Cost
		cost.OriginDetails = costWrapper.Rajaongkir.OriginDetails
		cost.DestinationDetails = costWrapper.Rajaongkir.DestinationDetails
		cost.Providers = costWrapper.Rajaongkir.Results

		return ServiceResult{Result: cost}
	}

	return ServiceResult{Error: errors.New("Invalid Request")}
}
