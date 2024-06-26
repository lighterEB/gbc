package main

import (
	"context"
	"encoding/json"
	"gbc/backend/api"
	"gbc/backend/response"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) QueryData() interface{} {
	products, err := api.QueryData()
	if err != nil {
		return response.NewErrorResponse[string](500, "Failed to query data")
	}
	productsJson, err := json.Marshal(products)
	if err != nil {
		return response.NewErrorResponse[string](500, "Failed to marshal data")
	}
	productsRaw := json.RawMessage(productsJson)
	return response.NewSuccessResponse(productsRaw)
}

func (a *App) KeyGenRequest(info string) interface{} {
	return api.KeyGenRequest(info)
} 
