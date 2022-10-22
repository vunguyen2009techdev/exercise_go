package db

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Init(sqlUri string) *xorm.Engine {
	parts := strings.SplitN(sqlUri, "://", 2)
	if len(parts) != 2 {
		err := fmt.Errorf("invalid: %s", sqlUri)
		log.Fatalln(err)
	}

	db, err := xorm.NewEngine(parts[0], parts[1])

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
