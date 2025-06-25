package handlers

import (
	"net/http"
	"strings"

	"github.com/Kaa-dan/skill-management/common"
	"github.com/Kaa-dan/skill-management/managers"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	groupName   string
	userManager managers.UserManager
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		groupName:   "/api/users",
		userManager: userManager,
	}
}

// RegisterUserApis registers all user-related routes
func (h *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(h.groupName)
	{
		userGroup.POST("", h.Create)
		userGroup.GET("", h.List)
		userGroup.GET("/:user_id", h.Detail)
		userGroup.DELETE("/:user_id", h.Delete)
		userGroup.PATCH("/:user_id", h.Update)
	}
}

// Create handles user creation requests
func (h *UserHandler) Create(ctx *gin.Context) {
	userData := &common.UserCreationInput{}

	if err := ctx.ShouldBindJSON(userData); err != nil {
		h.handleError(ctx, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Validate input
	if err := h.validateUserInput(userData); err != nil {
		h.handleError(ctx, http.StatusBadRequest, "Validation failed", err)
		return
	}

	newUser, err := h.userManager.Create(userData)
	if err != nil {
		h.handleError(ctx, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	h.handleSuccess(ctx, http.StatusCreated, "User created successfully", newUser)
}

// List handles listing all users
func (h *UserHandler) List(ctx *gin.Context) {
	users, err := h.userManager.List()
	if err != nil {
		h.handleError(ctx, http.StatusInternalServerError, "Failed to retrieve users", err)
		return
	}

	h.handleSuccess(ctx, http.StatusOK, "Users retrieved successfully", users)
}

// Detail handles retrieving a specific user
func (h *UserHandler) Detail(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		h.handleError(ctx, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	user, err := h.userManager.Detail(userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Failed to retrieve user"

		// Handle specific errors
		if err == managers.ErrUserNotFound {
			statusCode = http.StatusNotFound
			message = "User not found"
		} else if err == managers.ErrInvalidUserID {
			statusCode = http.StatusBadRequest
			message = "Invalid user ID"
		}

		h.handleError(ctx, statusCode, message, err)
		return
	}

	h.handleSuccess(ctx, http.StatusOK, "User retrieved successfully", user)
}

// Delete handles user deletion
func (h *UserHandler) Delete(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		h.handleError(ctx, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	err := h.userManager.Delete(userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Failed to delete user"

		// Handle specific errors
		if err == managers.ErrUserNotFound {
			statusCode = http.StatusNotFound
			message = "User not found"
		} else if err == managers.ErrInvalidUserID {
			statusCode = http.StatusBadRequest
			message = "Invalid user ID"
		}

		h.handleError(ctx, statusCode, message, err)
		return
	}

	h.handleSuccess(ctx, http.StatusOK, "User deleted successfully", nil)
}

// Update handles user updates
func (h *UserHandler) Update(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		h.handleError(ctx, http.StatusBadRequest, "User ID is required", nil)
		return
	}

	userData := &common.UserCreationInput{}
	if err := ctx.ShouldBindJSON(userData); err != nil {
		h.handleError(ctx, http.StatusBadRequest, "Invalid request data", err)
		return
	}

	// Validate input
	if err := h.validateUserInput(userData); err != nil {
		h.handleError(ctx, http.StatusBadRequest, "Validation failed", err)
		return
	}

	updatedUser, err := h.userManager.Update(userData, userID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Failed to update user"

		// Handle specific errors
		if err == managers.ErrUserNotFound {
			statusCode = http.StatusNotFound
			message = "User not found"
		} else if err == managers.ErrInvalidUserID {
			statusCode = http.StatusBadRequest
			message = "Invalid user ID"
		}

		h.handleError(ctx, statusCode, message, err)
		return
	}

	h.handleSuccess(ctx, http.StatusOK, "User updated successfully", updatedUser)
}

// Helper methods

// handleError standardizes error responses
func (h *UserHandler) handleError(ctx *gin.Context, statusCode int, message string, err error) {
	response := gin.H{
		"success": false,
		"message": message,
		"data":    nil,
	}

	// In development, you might want to include the actual error
	// In production, be careful not to expose sensitive information
	if gin.Mode() == gin.DebugMode && err != nil {
		response["error"] = err.Error()
	}

	ctx.JSON(statusCode, response)
}

// handleSuccess standardizes success responses
func (h *UserHandler) handleSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// validateUserInput validates user input data
func (h *UserHandler) validateUserInput(userData *common.UserCreationInput) error {
	if userData == nil {
		return common.ErrInvalidInput
	}

	if strings.TrimSpace(userData.Fullname) == "" {
		return common.ErrMissingFullname
	}

	if strings.TrimSpace(userData.Email) == "" {
		return common.ErrMissingEmail
	}

	// Add email format validation if needed
	if !h.isValidEmail(userData.Email) {
		return common.ErrInvalidEmail
	}

	return nil
}

// isValidEmail performs basic email validation
func (h *UserHandler) isValidEmail(email string) bool {
	// Basic email validation - you might want to use a more robust solution
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Optional: Add middleware for logging, authentication, etc.
func (h *UserHandler) LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Add logging logic here
		ctx.Next()
	}
}

// Optional: Add middleware for authentication
func (h *UserHandler) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Add authentication logic here
		ctx.Next()
	}
}
