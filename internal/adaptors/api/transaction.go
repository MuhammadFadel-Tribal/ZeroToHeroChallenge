package api

import (
	"encoding/json"
	"net/http"
	"zerotoherochallenge/internal/models"
	"zerotoherochallenge/internal/services/Transaction"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type TransactionController struct {
	log            *zap.SugaredLogger
	validate       *validator.Validate
	transactionSvc *Transaction.TransactionService
}

func TransactionControllerInitializer(server *HTTPServer, logger *zap.SugaredLogger, v *validator.Validate, ts *Transaction.TransactionService) *TransactionController {
	b := &TransactionController{
		log:            logger,
		validate:       v,
		transactionSvc: ts,
	}

	// Load routes
	server.Router.Group(func(r chi.Router) {
		r.Get("/transactions", b.handleGetAll)
		r.Get("/transactions/{id}", b.handleGetOne)
		r.Post("/transactions", b.handleCreate)
	})

	return b
}

func (b *TransactionController) handleGetAll(w http.ResponseWriter, r *http.Request) {
	result, ok := b.transactionSvc.FindAllTransactions(r.Context())

	if !ok {
		b.log.Errorf("There is no transactions")
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}

func (b *TransactionController) handleGetOne(w http.ResponseWriter, r *http.Request) {
	beneficiaryID := chi.URLParam(r, "id")

	if len(beneficiaryID) == 0 {
		b.log.Errorf("Beneficiary Id is required")
		return
	}

	result, ok := b.transactionSvc.FindTransactionDetails(r.Context(), beneficiaryID)

	if !ok {
		b.log.Errorf("Transaction is not exist")
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}

func (b *TransactionController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var body models.CreateTransactionDto

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		b.log.Errorf("Malformed body: %s", err)
		RenderJSON(r.Context(), w, http.StatusNotAcceptable, err)
		return
	}

	//body.CreatedAt = time.Now().UTC()
	err = b.validate.Struct(&body)
	if err != nil {
		b.log.Errorf("Validation error: %s", err)
		RenderJSON(r.Context(), w, http.StatusBadRequest, err)
		return
	}

	// Call the service
	result, status := b.transactionSvc.AddTransaction(r.Context(), &body)

	if status == false {
		b.log.Errorf("Add transaction failed. %v", err)
		RenderJSON(r.Context(), w, http.StatusInternalServerError, status)
		return
	}

	RenderJSON(r.Context(), w, http.StatusOK, result)
}
