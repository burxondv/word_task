package v1

import (
	"net/http"

	"github.com/burxondv/word_task/api/models"
	"github.com/burxondv/word_task/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /word [post]
// @Summary Create a word
// @Description Create a word
// @Tags word
// @Accept json
// @Produce json
// @Param word body map[string]int true "Word"
// @Success 200
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateWord(c *gin.Context) {
	var (
		req = make(map[string]int)
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.storage.Word().Create(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "Successfully created")
}

// @Router /word [get]
// @Summary Get word
// @Description Get word
// @Tags word
// @Accept json
// @Produce json
// @Param filter query models.GetAllParam false "Filter"
// @Success 200 {object} models.Word
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetWord(c *gin.Context) {
	req, err := validateGetParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	resp, err := h.storage.Word().GetAll(&repo.GetWordParam{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getWordResponse(resp))
}

func getWordResponse(resp *repo.GetWordResult) *models.GetWordResponse {
	response := models.GetWordResponse{
		Words: make([]*models.Word, 0),
		Count: resp.Count,
	}

	for _, word := range resp.Words {
		u := parseWordModel(word)
		response.Words = append(response.Words, &u)
	}

	return &response
}

func parseWordModel(word *repo.GetWord) models.Word {
	return models.Word{
		Word:  word.Word,
		Point: word.Point,
	}
}
