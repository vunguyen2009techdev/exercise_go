package main

import (
	"bufio"
	"encoding/csv"
	dto "exercise_go/src/dto"
	"fmt"
	"os"
	"strings"
)

func main() {
	fsentences, err := os.Open("./src/data/sentences.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened Sentences CSV file")
	defer fsentences.Close()

	reader := csv.NewReader(bufio.NewReader(fsentences))
	reader.Comma = ';'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	csvSentences, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	flinks, err := os.Open("./src/data/links.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened Links CSV file")
	defer flinks.Close()

	reader2 := csv.NewReader(bufio.NewReader(flinks))
	reader2.Comma = ';'
	reader2.LazyQuotes = true
	reader2.FieldsPerRecord = -1
	csvLinks, err := reader2.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	faudios, err := os.Open("./src/data/sentences_with_audio.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened Sentences With Audio CSV file")
	defer faudios.Close()

	reader3 := csv.NewReader(bufio.NewReader(faudios))
	reader3.Comma = ';'
	reader3.LazyQuotes = true
	reader3.FieldsPerRecord = -1
	csvAudios, err := reader3.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	var eng_sentences dto.Sentences
	var vie_sentences dto.Sentences
	var links dto.Links
	var audios dto.Audios

	for _, line := range csvSentences {
		if len(line) < 7 {
			if strings.Contains(line[0], "	eng	") {
				split := strings.Split(line[0], "	eng	")
				eng_sentences = append(eng_sentences, &dto.Sentence{
					SentenceId: split[0],
					Lang:       "eng",
					Text:       split[1],
				})
			}

			if strings.Contains(line[0], "	vie	") {
				split := strings.Split(line[0], "	vie	")
				vie_sentences = append(vie_sentences, &dto.Sentence{
					SentenceId: split[0],
					Lang:       "vie",
					Text:       split[1],
				})
			}
		}
	}

	for _, line := range csvLinks {
		if len(line) < 7 {
			split := strings.Split(line[0], "	")
			links = append(links, dto.Link{
				SentenceId:    split[0],
				TranslationId: split[1],
			})
		}
	}

	for _, line := range csvAudios {
		if len(line) < 7 {
			if strings.Contains(line[0], "http") {
				splitStr := strings.Split(line[0], "	")
				split := strings.Split(line[0], "http")
				audios = append(audios, dto.Audio{
					SentenceId:   splitStr[0],
					AttributeUrl: fmt.Sprintf("http%+v", split[1]),
				})
			}
		}
	}

	// Create csv file
	f, err := os.Create("./src/data/sentence_all_records.csv")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	// End create csv file

	for _, eng := range eng_sentences {
		for _, audio := range audios {
			if eng.SentenceId == audio.SentenceId {
				if audio.AttributeUrl != "" {
					eng.AudioUrl = fmt.Sprintf("%+v/%+v/%+v.mp3", audio.AttributeUrl, eng.Lang, eng.SentenceId)
				}
			}
		}

		for _, link := range links {
			if link.SentenceId == eng.SentenceId {
				eng.Links = append(eng.Links, link.TranslationId)
			}
		}

		for _, link := range eng.Links {
			for _, vieData := range vie_sentences {
				if link == vieData.SentenceId {
					eng.VieSentenceId = vieData.SentenceId
					eng.VieText = vieData.Text
				}
			}
		}

		row := []string{eng.SentenceId, eng.Text, eng.VieSentenceId, eng.VieText, eng.AudioUrl}
		fmt.Println("result: ", row)
		if err := w.Write(row); err != nil {
			fmt.Println("err writing record to file", err)
		}
	}
}
