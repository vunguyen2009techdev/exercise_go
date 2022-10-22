package dto

type Link struct {
	SentenceId    string `json:"sentenceId"`
	TranslationId string `json:"translationId"`
}

type Links []Link

type LinkStrs []string
