package ro

//ServiceResult
type ServiceResult struct {
	Result interface{}
	Error  error
}

//RajaOngkirService
//Generic abstraction for Raja Ongkir Service
type RajaOngkirService interface {
	GetProvince(QueryRequest) ServiceResult
	GetCity(QueryRequest) ServiceResult
	GetCost(QueryRequest) ServiceResult
}
