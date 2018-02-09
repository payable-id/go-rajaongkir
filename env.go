package ro

// Env type
type Env int

const (
	//Starter Env, enviroment constant for Starter Raja Ongkir Account
	Starter Env = iota

	//Basic Env, enviroment constant for Basic Raja Ongkir Account
	Basic Env = iota

	//Pro Env, enviroment constant for Pro Raja Ongkir Account
	Pro Env = iota
)

//String function, return Env as a string
func (e Env) String() string {
	switch e {
	case Starter:
		return "https://api.rajaongkir.com/starter"
	case Basic:
		return "https://api.rajaongkir.com/basic"
	case Pro:
		return "https://pro.rajaongkir.com/api"
	default:
		return "UNDEFINED"
	}
}
