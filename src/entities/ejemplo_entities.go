package entities

// Este fue nuestra primera estructura
type HolaMundo struct {
	Dato string `json:"dato"`
}

type ResponseDatos struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
