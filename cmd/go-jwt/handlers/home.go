package handlers

import (
	"go-jwt/interval/templates"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleHome(ctx *gin.Context) {
	templates.Home().Render(ctx.Request.Context(), ctx.Writer)
}
