package models

import (
	"fmt"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (m Movie) ToString() string {
	return fmt.Sprintf(
		" => Movie \"%s\" (Id: %s, ISBN: %s) directed by \"%s %s\"\n",
		m.Title, m.ID, m.Isbn, m.Director.FirstName, m.Director.LastName)
}
