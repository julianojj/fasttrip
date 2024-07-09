package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/core/usecases"
)

type (
	BookingController struct {
		makeBooking *usecases.MakeBooking
	}
	ErrorResponse struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

func NewBookingController(
	makeBooking *usecases.MakeBooking,
) *BookingController {
	return &BookingController{
		makeBooking,
	}
}

// MakeBooking godoc
//
//	@Summary		Faz nova reserva
//	@Description	Faz nova reserva
//	@Tags			booking
//	@Accept			json
//	@Produce		json
//	@Param			input	body		usecases.MakeBookingInput	true	"MakeBookingInput"
//	@Success		201		{array}		usecases.MakeBookingOutput
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		422		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/make_booking [post]
func (c *BookingController) MakeBooking(w http.ResponseWriter, r *http.Request) {
	var input usecases.MakeBookingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	output, err := c.makeBooking.Execute(&input)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		case *exceptions.NotFoundException:
			statusCode = http.StatusNotFound
		}

		handleError(w, statusCode, err.Error(), err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response", err)
		return
	}
}

func handleError(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Message: message,
		Code:    statusCode,
	})
	log.Printf("Error %d: %s - %v", statusCode, message, err)
}
