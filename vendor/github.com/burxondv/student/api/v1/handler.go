package v1

import (
	"github.com/burxondv/student/api/models"
	"github.com/burxondv/student/config"
	"github.com/burxondv/student/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func validateGetParams(c *gin.Context) (*models.GetAllParam, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllParam{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: c.Query("search"),
	}, nil
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
