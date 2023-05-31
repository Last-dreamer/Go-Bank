package api

import (
	db "bank/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type GetAllAccountsRequest struct {
	PageNo   int32 `form:"page_no" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) GetAllAccounts(c *gin.Context) {
	var req GetAllAccountsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.GetListAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageNo - 1) * req.PageSize,
	}

	list, err := server.store.GetListAccount(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, list)
}
