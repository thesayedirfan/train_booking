package entity


type Train struct {
	Name string
	From string
	To string
	Sections map[string][]Seat
}

type Seat struct {
	Number int64
	User *User
}