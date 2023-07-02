package student

type Group struct {
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	Students []Student `json:"students"`
}
