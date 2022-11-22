package v1

import (
	"fmt"
	"net/http"

	"github.com/burxondv/student/api/models"
	"github.com/burxondv/student/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /student [post]
// @Summary Create a student
// @Description Create a student
// @Tags student
// @Accept json
// @Produce json
// @Param user body models.CreateStudentRequest true "Student"
// @Success 201 {object} models.Student
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateStudent(c *gin.Context) {
	req := models.CreateStudentRequest{
		Students: make([]*models.CreateStudent, 0),
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	slc := convert(req)

	fmt.Println(slc)

	err = h.storage.Student().Create(slc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}

// @Router /student [get]
// @Summary Get student
// @Description Get student
// @Tags student
// @Accept json
// @Produce json
// @Param filter query models.GetAllParam false "Filter"
// @Success 200 {object} models.GetStudentResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetStudent(c *gin.Context) {
	req, err := validateGetParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	resp, err := h.storage.Student().Get(&repo.GetStudentParam{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getStudentResponse(resp))
}

func getStudentResponse(data *repo.GetStudentResult) *models.GetStudentResponse {
	response := models.GetStudentResponse{
		Students: make([]*models.Student, 0),
		Count:   data.Count,
	}

	for _, student := range data.Students {
		u := parseStudentModel(student)
		response.Students = append(response.Students, &u)
	}

	return &response
}

func parseStudentModel(student *repo.Student) models.Student {
	return models.Student{
		ID:          student.ID,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Username:    student.Username,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		CreatedAt:   student.CreatedAt,
	}
}

func convert(st models.CreateStudentRequest) []*repo.Student {
	req := make([]*repo.Student, 0)
	for _, s := range st.Students {
		var rm repo.Student
		rm.FirstName = s.FirstName
		rm.LastName = s.LastName
		rm.Username = s.Username
		rm.Email = s.Email
		rm.PhoneNumber = s.PhoneNumber
		req = append(req, &rm)
	}

	return req
}
