package api

import (
	db "Bankstore/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	//TODO: add routes to router
	router.POST("/accounts, server.CreateAccount")
	server.router = router
	return server
}

type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required, oneof=USD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	//Десериализация входящего JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse((err)))
		return
	}

	//На основе входящего CreateAccountRequest создаем CreateAccountParams
	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}
	
	//На основе arg создаем аккаунт
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// errorResponse return gin.H -> map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//Start server method

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}