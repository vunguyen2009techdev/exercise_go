package main

import (
	"encoding/csv"
	"exercise_go/src/db"
	dto "exercise_go/src/dto"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
	"xorm.io/xorm"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("SQL_URI").(string)
	h := db.Init(dbUrl)
	db := GetCollectionRecord(h)

	// Open the file
	csvfile, err := os.Open("./src/data/sentence_all_records.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var records dto.Records
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, dto.Record{
			SentenceId:    record[0],
			Text:          record[1],
			VieSentenceId: record[2],
			VieText:       record[3],
			AudioUrl:      record[4],
		})
	}

	InsertRecord(db, records)
	fmt.Println("DONE SEEDING")
}

func GetCollectionRecord(h *xorm.Engine) *xorm.Session {
	return h.Table("records")
}

func InsertRecord(db *xorm.Session, data dto.Records) {
	fmt.Println("here")
	status, err := db.Insert(&data)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("status: ", status)
}
