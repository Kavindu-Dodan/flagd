package ofrep

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/open-feature/flagd/core/pkg/evaluator"
	"github.com/open-feature/flagd/core/pkg/logger"
	"github.com/open-feature/flagd/core/pkg/service/ofrep"
	"github.com/rs/xid"
)

const (
	key              = "key"
	singleEvaluation = "/ofrep/v1/evaluate/flags/{key}"
	bulkEvaluation   = "/ofrep/v1/evaluate/{path:flags\\/|flags}"
)

type handler struct {
	Logger    *logger.Logger
	evaluator evaluator.IEvaluator
}

func NewHandler(logger *logger.Logger, evaluator evaluator.IEvaluator) http.Handler {
	h := handler{
		logger,
		evaluator,
	}

	router := mux.NewRouter()
	router.HandleFunc(singleEvaluation, h.HandleFlagEvaluation).Methods("POST")
	router.HandleFunc(bulkEvaluation, h.HandleBulkEvaluation).Methods("POST")
	return router
}

func (h *handler) HandleFlagEvaluation(w http.ResponseWriter, r *http.Request) {
	requestID := xid.New().String()

	// obtain flag key
	vars := mux.Vars(r)
	flagKey := vars[key]
	if vars == nil || flagKey == "" {
		h.writeJSONToResponse(
			http.StatusInternalServerError,
			ofrep.InternalError{ErrorDetails: "failed to obtain the flag key from the request"}, w)
		return
	}

	request := ofrep.Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		// Special handling for invalid context
		h.writeJSONToResponse(http.StatusBadRequest, ofrep.ContextErrorResponseAndStatus(flagKey), w)
		return
	}

	context := map[string]any{}
	if res, ok := request.Context.(map[string]any); ok {
		context = res
	} else {
		h.Logger.WarnWithID(requestID, "provided context does not comply with flagd, continuing ignoring the context")
	}

	evaluation := h.evaluator.ResolveAsAnyValue(r.Context(), requestID, flagKey, context)
	if evaluation.Error != nil {
		status, evaluationError := ofrep.ValueToStatusAndError(evaluation)
		h.writeJSONToResponse(status, evaluationError, w)
	} else {
		h.writeJSONToResponse(http.StatusOK, ofrep.SuccessResponseFrom(evaluation), w)
	}
}

func (h *handler) HandleBulkEvaluation(w http.ResponseWriter, r *http.Request) {
	requestID := xid.New().String()

	request := ofrep.Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		// Special handling for invalid context
		h.writeJSONToResponse(http.StatusBadRequest, ofrep.ContextErrorResponseForBulkAndStatus(), w)
		return
	}

	context := map[string]any{}
	if res, ok := request.Context.(map[string]any); ok {
		context = res
	} else {
		h.Logger.WarnWithID(requestID, "provided context does not comply with flagd, continuing ignoring the context")
	}

	evaluations := h.evaluator.ResolveAllValues(r.Context(), requestID, context)
	h.writeJSONToResponse(http.StatusOK, ofrep.BulkEvaluationResponseFrom(evaluations), w)
}

func (h *handler) writeJSONToResponse(status int, payload interface{}, w http.ResponseWriter) {
	// first marshal payload
	marshal, err := json.Marshal(payload)
	if err != nil {
		// always a 500
		h.Logger.Warn(fmt.Sprintf("error marshelling the response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(marshal)
	if err != nil {
		h.Logger.Warn(fmt.Sprintf("error while writing response: %v", err))
		return
	}
}
