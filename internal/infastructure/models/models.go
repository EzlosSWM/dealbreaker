package models

type Cards struct {
	Cards []Card
}

type Card struct {
	ID       uint   `json:"id"`
	JokeType string `json:"joke_type"`
	Joke     string `json:"joke"`
	Topic    string `json:"topic"`
}
