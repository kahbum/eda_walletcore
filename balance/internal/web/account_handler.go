package web

import (
	"encoding/json"
	"net/http"

	"github.com/kahbum/eda_balance/internal/usecase/get_account_balance"
)

type WebAccountHandler struct {
	GetAccountBalanceUseCase get_account_balance.GetAccountBalanceUseCase
}

func NewWebAccountHandler(getAccountBalanceUseCase get_account_balance.GetAccountBalanceUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		GetAccountBalanceUseCase: getAccountBalanceUseCase,
	}
}

func (h *WebAccountHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("account_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dto := get_account_balance.GetAccountBalanceInputDTO{
		ID: id,
	}

	output, err := h.GetAccountBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
