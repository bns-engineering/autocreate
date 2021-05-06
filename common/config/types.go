package config

//Configuration Struct
type Configuration struct {
	Debug        bool
	Mambu        Mambu
	Start        int
	Until        int
	ValueDate    string
	MaturityDate string
}

//Mambu Struct
type Mambu struct {
	Endpoint      string
	Authorization string
}

type Client struct {
	ClientID string
}
