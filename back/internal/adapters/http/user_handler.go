package http

import (
	"net/http"

	appuser "github.com/RoundRonin/name-subject-to-change/back/internal/application/user"
	"github.com/RoundRonin/name-subject-to-change/back/internal/application/user/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *appuser.Service
}

func NewUserHandler(svc *appuser.Service) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.CreateUser)
	r.GET("/users/:id", h.GetUser)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "User info"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appuser.ToUserResponse(u))
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get user info by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	u, err := h.svc.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, appuser.ToUserResponse(u))
}
