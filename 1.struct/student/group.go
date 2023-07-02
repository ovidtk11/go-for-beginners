package student

type Group struct {
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	Students []student `json:"students"`
}
