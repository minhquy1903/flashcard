package presenter

type RegisterResponse struct {
	Name     string
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken string
}
