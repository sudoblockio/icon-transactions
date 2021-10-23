package rest

import (
	"encoding/json"
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/geometry-labs/icon-transactions/config"
	"github.com/geometry-labs/icon-transactions/crud"
)

type TransactionsQuery struct {
	Limit       int    `query:"limit"`
	Skip        int    `query:"skip"`
	From        string `query:"from"`
	To          string `query:"to"`
	Type        string `query:"type"`
	BlockNumber int    `query:"block_number"`
	Method      string `query:"method"`
}

func TransactionsAddHandlers(app *fiber.App) {

	prefix := config.Config.RestPrefix + "/transactions"

	app.Get(prefix+"/", handlerGetTransactions)
	app.Get(prefix+"/details/:hash", handlerGetTransactionDetails)
	app.Get(prefix+"/address/:address", handlerGetTransactionAddress)
	app.Get(prefix+"/internal/:hash", handlerGetInternalTransactionsByHash)
	app.Get(prefix+"/internal/address/:address", handlerGetInternalTransactionsAddress)
	app.Get(prefix+"/token-transfers", handlerGetTokenTransfers)
	app.Get(prefix+"/token-transfers/address/:address", handlerGetTokenTransfersAddress)
	app.Get(prefix+"/token-transfers/token-contract/:token_contract_address", handlerGetTokenTransfersTokenContract)
}

// Transactions
// @Summary Get Transactions
// @Description get historical transactions
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param from query string false "find by from address"
// @Param to query string false "find by to address"
// @Param type query string false "find by type"
// @Param block_number query int false "find by block number"
// @Param method query string false "find by method"
// @Router /api/v1/transactions [get]
// @Success 200 {object} []models.TransactionAPIList
// @Failure 422 {object} map[string]interface{}
func handlerGetTransactions(c *fiber.Ctx) error {
	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	// NOTE: TEMP casting string types for type field
	if params.Type == "regular" {
		params.Type = "transaction"
	} else if params.Type == "internal" {
		params.Type = "log"
	}

	// Get Transactions
	transactions, count, err := crud.GetTransactionModel().SelectManyAPI(
		params.Limit,
		params.Skip,
		params.From,
		params.To,
		params.Type,
		params.BlockNumber,
		params.Method,
	)
	if err != nil {
		zap.S().Warnf("Transactions CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve transactions"}`)
	}

	if len(*transactions) == 0 {
		// No Content
		c.Status(204)
	}

	// Set X-TOTAL-COUNT
	if count != -1 {
		// Filters given, count some
		c.Append("X-TOTAL-COUNT", strconv.FormatInt(count, 10))
	} else {
		// No filters given, count all
		// Total count in the transaction_counts table
		counter, err := crud.GetTransactionCountModel().SelectCount("regular")
		if err != nil {
			counter = 0
			zap.S().Warn("Could not retrieve transaction count: ", err.Error())
		}

		c.Append("X-TOTAL-COUNT", strconv.FormatUint(counter, 10))
	}

	body, _ := json.Marshal(&transactions)
	return c.SendString(string(body))
}

// Transactions by Address
// @Summary Get Transactions by address
// @Description get transactions by address
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param address path string true "address"
// @Router /api/v1/transactions/address/{address} [get]
// @Success 200 {object} models.TransactionAPIList
// @Failure 422 {object} map[string]interface{}
func handlerGetTransactionAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		c.Status(422)
		return c.SendString(`{"error": "address required"}`)
	}

	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	transactions, err := crud.GetTransactionModel().SelectManyByAddressAPI(
		params.Limit,
		params.Skip,
		address,
	)
	if err != nil {
		c.Status(500)
		return c.SendString(`{"error": "no transactions found"}`)
	}

	// X-TOTAL-COUNT
	count, err := crud.GetTransactionCountByAddressModel().SelectCount(address)
	if err != nil {
		count = 0
		zap.S().Warn("Could not retrieve transaction count: ", err.Error())
	}

	c.Append("X-TOTAL-COUNT", strconv.FormatUint(count, 10))

	body, _ := json.Marshal(&transactions)
	return c.SendString(string(body))
}

// Transaction Details
// @Summary Get Transaction Details
// @Description get details of a transaction
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param hash path string true "transaction hash"
// @Router /api/v1/transactions/details/{hash} [get]
// @Success 200 {object} models.TransactionAPIDetail
// @Failure 422 {object} map[string]interface{}
func handlerGetTransactionDetails(c *fiber.Ctx) error {
	hash := c.Params("hash")

	if hash == "" {
		c.Status(422)
		return c.SendString(`{"error": "hash required"}`)
	}

	transaction, err := crud.GetTransactionModel().SelectOneAPI(hash, -1)
	if err != nil {
		c.Status(404)

		zap.S().Warn(err.Error())
		return c.SendString(`{"error": "no transaction found"}`)
	}

	body, _ := json.Marshal(&transaction)
	return c.SendString(string(body))
}

// Internal transactions by hash
// @Summary Get internal transactions by hash
// @Description Get internal transactions by hash
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param hash path string true "find by hash"
// @Router /api/v1/transactions/internal/{hash} [get]
// @Success 200 {object} []models.TransactionInternalAPIList
// @Failure 422 {object} map[string]interface{}
func handlerGetInternalTransactionsByHash(c *fiber.Ctx) error {
	hash := c.Params("hash")
	if hash == "" {
		c.Status(422)
		return c.SendString(`{"error": "hash required"}`)
	}

	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	if hash == "" {
		c.Status(422)
		return c.SendString(`{"error": "hash required"}`)
	}

	internalTransactions, err := crud.GetTransactionModel().SelectManyInternalAPI(
		params.Limit,
		params.Skip,
		hash,
	)
	if err != nil {
		c.Status(500)
		return c.SendString(`{"error": "no internal transaction found"}`)
	}

	if len(*internalTransactions) == 0 {
		// No Content
		c.Status(204)
	}

	body, _ := json.Marshal(&internalTransactions)
	return c.SendString(string(body))
}

// Internal transactions by address
// @Summary Get internal transactions by address
// @Description Get internal transactions by address
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param address path string true "find by address"
// @Router /api/v1/transactions/internal/address/{address} [get]
// @Success 200 {object} []models.TransactionInternalAPIList
// @Failure 422 {object} map[string]interface{}
func handlerGetInternalTransactionsAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		c.Status(422)
		return c.SendString(`{"error": "address required"}`)
	}

	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	internalTransactions, err := crud.GetTransactionModel().SelectManyInternalByAddressAPI(
		params.Limit,
		params.Skip,
		address,
	)
	if err != nil {
		c.Status(500)
		return c.SendString(`{"error": "no internal transaction found"}`)
	}

	if len(*internalTransactions) == 0 {
		// No Content
		c.Status(204)
	}

	// X-TOTAL-COUNT
	count, err := crud.GetTransactionInternalCountByAddressModel().SelectCount(address)
	if err != nil {
		count = 0
		zap.S().Warn("Could not retrieve transaction count: ", err.Error())
	}

	c.Append("X-TOTAL-COUNT", strconv.FormatUint(count, 10))

	body, _ := json.Marshal(&internalTransactions)
	return c.SendString(string(body))
}

// TokenTransfers
// @Summary Get token transfers
// @Description get historical token transfers
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param from query string false "find by from address"
// @Param to query string false "find by to address"
// @Param block_number query int false "find by block number"
// @Router /api/v1/transactions/token-transfers [get]
// @Success 200 {object} []models.TokenTransfer
// @Failure 422 {object} map[string]interface{}
func handlerGetTokenTransfers(c *fiber.Ctx) error {
	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	// Get Transactions
	tokenTransfers, count, err := crud.GetTokenTransferModel().SelectMany(
		params.Limit,
		params.Skip,
		params.From,
		params.To,
		params.BlockNumber,
	)
	if err != nil {
		zap.S().Warnf("Transactions CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve transactions"}`)
	}

	if len(*tokenTransfers) == 0 {
		// No Content
		c.Status(204)
	}

	// Set X-TOTAL-COUNT
	if count != -1 {
		// Filters given, count some
		c.Append("X-TOTAL-COUNT", strconv.FormatInt(count, 10))
	} else {
		// No filters given, count all
		// Total count in the token_transfer_counts table
		counter, err := crud.GetTransactionCountModel().SelectCount("token_transfer")
		if err != nil {
			counter = 0
			zap.S().Warn("Could not retrieve token transfer count: ", err.Error())
		}

		c.Append("X-TOTAL-COUNT", strconv.FormatUint(counter, 10))
	}

	body, _ := json.Marshal(&tokenTransfers)
	return c.SendString(string(body))
}

// TokenTransfersAddress
// @Summary Get token transfer by address
// @Description get historical token transfers by address
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param address path string true "find by address"
// @Router /api/v1/transactions/token-transfers/address/{address} [get]
// @Success 200 {object} []models.TokenTransfer
// @Failure 422 {object} map[string]interface{}
func handlerGetTokenTransfersAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		c.Status(422)
		return c.SendString(`{"error": "address required"}`)
	}

	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	// Get Transactions
	tokenTransfers, err := crud.GetTokenTransferModel().SelectManyByAddress(
		params.Limit,
		params.Skip,
		address,
	)
	if err != nil {
		zap.S().Warnf("Transactions CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve transactions"}`)
	}

	if len(*tokenTransfers) == 0 {
		// No Content
		c.Status(204)
	}

	// Set X-TOTAL-COUNT
	// Token transfer by address
	count, err := crud.GetTokenTransferCountByAddressModel().SelectCount(address)
	if err != nil {
		count = 0
		zap.S().Warn("Could not retrieve token transfer count: ", err.Error())
	}

	c.Append("X-TOTAL-COUNT", strconv.FormatUint(count, 10))

	body, _ := json.Marshal(&tokenTransfers)
	return c.SendString(string(body))
}

// TokenTransfersTokenContract
// @Summary Get token transfers by token contract
// @Description get historical token transfers by token contract
// @Tags Transactions
// @BasePath /api/v1
// @Accept */*
// @Produce json
// @Param limit query int false "amount of records"
// @Param skip query int false "skip to a record"
// @Param token_contract_address path string true "find by token contract address"
// @Router /api/v1/transactions/token-transfers/token-contract/{token_contract_address} [get]
// @Success 200 {object} []models.TokenTransfer
// @Failure 422 {object} map[string]interface{}
func handlerGetTokenTransfersTokenContract(c *fiber.Ctx) error {
	tokenContractAddress := c.Params("token_contract_address")
	if tokenContractAddress == "" {
		c.Status(422)
		return c.SendString(`{"error": "token_contract_address required"}`)
	}

	params := new(TransactionsQuery)
	if err := c.QueryParser(params); err != nil {
		zap.S().Warnf("Transactions Get Handler ERROR: %s", err.Error())

		c.Status(422)
		return c.SendString(`{"error": "could not parse query parameters"}`)
	}

	// Default Params
	if params.Limit <= 0 {
		params.Limit = 25
	}

	// Check Params
	if params.Limit < 1 || params.Limit > config.Config.MaxPageSize {
		c.Status(422)
		return c.SendString(`{"error": "limit must be greater than 0 and less than 101"}`)
	}

	// Get Transactions
	tokenTransfers, err := crud.GetTokenTransferModel().SelectManyByTokenContractAddress(
		params.Limit,
		params.Skip,
		tokenContractAddress,
	)
	if err != nil {
		zap.S().Warnf("Transactions CRUD ERROR: %s", err.Error())
		c.Status(500)
		return c.SendString(`{"error": "could not retrieve transactions"}`)
	}

	if len(*tokenTransfers) == 0 {
		// No Content
		c.Status(204)
	}

	// X-TOTAL-COUNT
	count, err := crud.GetTokenTransferCountByTokenContractModel().SelectLargestCountByTokenContract(tokenContractAddress)
	if err != nil {
		count = 0
		zap.S().Warn("Could not retrieve token transfer count: ", err.Error())
	}

	c.Append("X-TOTAL-COUNT", strconv.FormatUint(count, 10))

	body, _ := json.Marshal(&tokenTransfers)
	return c.SendString(string(body))
}
