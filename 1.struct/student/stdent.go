package student

import "fmt"

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

// normal function

func A(gg Student) {
	fmt.Println("Hello", gg.Lastname)
}

func B(s string) {
	fmt.Println("Hello")
}

func C(st Student) string {
	return fmt.Sprintln("Hello", st.Firstname)
}

func E(firstname, lastname string) {

}

//receiver function

func (st *Student) D() {
	fmt.Println("Hello", st.Firstname)
	fmt.Println("Hello", st.Lastname)
}

func (st *Student) Oil() {
	fmt.Println("Hello", st.Firstname)
	fmt.Println("Hello", st.Lastname)
}

type Teacher struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Nickname  string `json:"nickname"`
}

func main() {
	s := Student{}
}

// 1. validate ชื่อ
