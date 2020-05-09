package cyoa

//Story map
type Story map[string]Chapter

// Option struct
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}

// Chapter struct
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
