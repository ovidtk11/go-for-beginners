package student

type student struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Nickname  string `json:"nickname"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
}

func NewStudent() student {
	return student{}
}

// Setter
func (s *student) SetFirstName(firstname string) {
	s.Firstname = firstname
}
func (s *student) SetLastname(lastname string) {
	s.Lastname = lastname
}
func (s *student) SetNickname(nickname string) {
	s.Nickname = nickname
}
func (s *student) SetAge(age int) {
	s.Age = age
}

func (s *student) SetHeight(height int) {
	s.Height = height
}

func (s *student) SetWeight(weight int) {
	s.Weight = weight
}

// Getter
func (s *student) GetFirstName() string {
	return s.Firstname
}

func (s *student) GetLastname() string {
	return s.Lastname
}
func (s *student) GetNickname() string {
	return s.Nickname
}
func (s *student) GetAge() int {
	return s.Age
}
func (s *student) GetHeight() int {
	return s.Height
}
func (s *student) GetWeight() int {
	return s.Weight
}
