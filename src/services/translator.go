package services

import (
	"exercise_go/src/dto"
	"fmt"
	"net/http"
	"strconv"

	cons "exercise_go/src/constant"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Translator interface {
	GetTranslates(c *gin.Context)
}

type translator struct {
	DB *xorm.Engine
}

func NewTranslatorService(db *xorm.Engine) Translator {
	return &translator{
		DB: db,
	}
}

func (t *translator) GetTranslates(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("page_number", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	ormSession := t.GetData(pageSize, pageNumber)
	var records dto.Records
	err := ormSession.Find(&records)
	if err != nil {
		fmt.Println("err: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, records)
}

func (t *translator) GetData(limit int, page int) *xorm.Session {
	sqlQuery := `1`
	offsetQuery := (page - 1) * limit

	sqlCommand := `SELECT * FROM ` + cons.TableRecord + ` WHERE ` + sqlQuery + ` ORDER BY sentenceId ASC LIMIT ` + fmt.Sprint(limit) + ` OFFSET ` + fmt.Sprint(offsetQuery)

	ormSession := t.GetCollectionRecord().SQL(sqlCommand)
	return ormSession
}

func (t *translator) GetCollectionRecord() *xorm.Session {
	return t.DB.Table("records")
}
