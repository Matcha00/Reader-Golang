package ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"reader/Reader-Golang/APP/init"
	"reader/Reader-Golang/APP/Model"
)

func ReaderSelectAll(c *gin.Context)  {

	var readers  []Model.Reader

	//if dbError := init.MYSQLORM.Find(readers).Error; dbError != nil {
	//
	//
	//	c.AbortWithError(400, dbError)
	//	return
	//
	//}

	if dbError := initdb.MYSQLORM.Find(&readers).Error; dbError != nil {
		c.AbortWithError(400, dbError)
		return
	}

	var result  = make([]Model.ReaderJson,len(readers))

	for index, reader := range readers {

		result[index] = reader.Serializer()

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"data": result,
	})
	//c.JSON(http.StatusOK,readers)

}

func DeleteReaderSingel(c *gin.Context) {

	var reader Model.Reader
	readerUrl := c.Params.ByName("url")
    fmt.Print(readerUrl)
	if dbError := initdb.MYSQLORM.Where("reader_url = ?", readerUrl).Delete(&reader).Error; dbError != nil  {

		fmt.Print(dbError)

		c.AbortWithError(400, dbError)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete_successful",
	})



}

func CreateReader(c *gin.Context)  {

	//title := c.Params.ByName("title")
	//url := c.Params.ByName("url")
	//plan := c.Params.ByName("plan")

	var reader Model.Reader

	c.BindJSON(&reader)
	initdb.MYSQLORM.Create(&reader)
	c.JSON(http.StatusOK, gin.H{
		"message": "update_success",
		"test": reader,
	})
}

func UpdataReader(c *gin.Context)  {

	url := c.Params.ByName("url")

	var reader  Model.Reader

	if dbError := initdb.MYSQLORM.Where("reader_url = ?", url).First(&reader).Error; dbError != nil  {

		fmt.Print(dbError)

		c.AbortWithError(400, dbError)
	}


	c.BindJSON(&reader)
	initdb.MYSQLORM.Save(&reader)
	c.JSON(http.StatusOK, gin.H{
		"message": "updata_successs",
	})


}