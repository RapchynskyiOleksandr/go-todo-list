package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the not protected route!"})
}

func (h *Handler) signIn(c *gin.Context) {

}
