package dto

type Audio struct {
	SentenceId   string `json:"sentenceId"`
	AuthorId     string `json:"authorId"`
	Author       string `json:"author"`
	License      string `json:"license"`
	AttributeUrl string `json:"attributeUrl"`
}

type Audios []Audio
