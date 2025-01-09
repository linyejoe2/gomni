package utils

type Auth struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
}

type Remote struct {
	IP   string `json:"ip"`
	Name string `json:"name"`
	Auth Auth   `json:"auth"`
}

type Remotes struct {
	Remotes []Remote `json:"remotes"`
}
