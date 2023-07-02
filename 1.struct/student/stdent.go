package student

type Student struct {
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Nickname  string     `json:"nickname"`
	age       int        `json:"age"`
	height    int        `json:"height"`
	weight    int        `json:"weight"`
	hobby     []string   `json:"hobby"`
	favMovie  []favMovie `json:"fav_movie"`
}

type favMovie struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Genre string `json:"genre"`
}
