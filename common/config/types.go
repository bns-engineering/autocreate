package config

//Configuration Struct
type Configuration struct {
	Debug   bool
	Mambu   Mambu
	Clients []Client
}

//Mambu Struct
type Mambu struct {
	Endpoint      string
	Authorization string
}

type Client struct {
	ClientID string
}
