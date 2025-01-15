package controllers

import (
	"crudappgolang/models"
	"crudappgolang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to get file")
		return
	}

	src, err := file.Open()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer src.Close()

	xlsx, err := excelize.OpenReader(src)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to parse Excel file")
		return
	}

	go models.ProcessExcel(xlsx)
	utils.RespondWithJSON(c, http.StatusOK, "File uploaded and processing started")
}

func ViewData(c *gin.Context) {
	data, err := models.GetCachedData()
	if err == nil {
		utils.RespondWithJSON(c, http.StatusOK, data)
		return
	}

	data, err = models.FetchAllUsers()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, data)
}

func EditRecord(c *gin.Context) {
	id := c.Param("id")
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := models.UpdateUser(id, input); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update record")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "Record updated successfully")
}
