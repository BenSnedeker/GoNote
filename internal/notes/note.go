package notes

type Note struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsLocked bool   `json:"isLocked"`
}

type Change struct {
	ID        string `json:"id"`
	Changes   string `json:"changes"`
	Timestamp int64  `json:"timestamp"`
}

func NewNote(id, title, content string) *Note {
	return nil
}

func ApplyChange(change Change) {

}
