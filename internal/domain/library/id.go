package library

type id struct {
	text string
}

func (id *id) Text() string {
	return id.text
}