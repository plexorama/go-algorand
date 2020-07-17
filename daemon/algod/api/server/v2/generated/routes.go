// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
