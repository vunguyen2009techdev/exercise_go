package dto

type Sentence struct {
	SentenceId    string   `json:"sentenceId"`
	Lang          string   `json:"lang"`
	Text          string   `json:"text"`
	Links         LinkStrs `json:"links"`
	VieSentenceId string   `json:"vieSentenceId"`
	VieText       string   `json:"vieText"`
	AudioUrl      string   `json:"audioUrl"`
}

type Sentences []*Sentence
