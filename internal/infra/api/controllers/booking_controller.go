package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/core/usecases"
)

type (
	BookingController struct {
		makeBooking *usecases.MakeBooking
		getBookings *usecases.GetBookings
		getBooking  *usecases.GetBooking
	}
	ErrorResponse struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
)

func NewBookingController(
	makeBooking *usecases.MakeBooking,
	getBookings *usecases.GetBookings,
	getBooking *usecases.GetBooking,
) *BookingController {
	return &BookingController{
		makeBooking,
		getBookings,
		getBooking,
	}
}

// MakeBooking godoc
//
//	@Summary		Faz nova reserva
//	@Description	Faz nova reserva
//	@Tags			booking
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Insert you bearer token"
//	@Param			input			body		usecases.MakeBookingInput	true	"MakeBookingInput"
//	@Success		201				{array}		usecases.MakeBookingOutput
//	@Failure		400				{object}	ErrorResponse
//	@Failure		404				{object}	ErrorResponse
//	@Failure		422				{object}	ErrorResponse
//	@Failure		500				{object}	ErrorResponse
//	@Router			/make_booking [post]
func (c *BookingController) MakeBooking(w http.ResponseWriter, r *http.Request) {
	var input usecases.MakeBookingInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body")
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
		handleError(w, statusCode, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

// GetBookings godoc
//
//	@Summary		Pegar reservas
//	@Description	Pegar reservas
//	@Tags			booking
//	@Produce		json
//	@Param			Authorization	header		string	true	"Insert you bearer token"
//	@Success		200				{array}		usecases.GetBookingsOutput
//	@Failure		422				{object}	ErrorResponse
//	@Failure		500				{object}	ErrorResponse
//	@Router			/get_bookings [get]
func (c *BookingController) GetBookings(w http.ResponseWriter, r *http.Request) {
	output, err := c.getBookings.Execute()
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		}
		handleError(w, statusCode, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

// GetBooking godoc
//
//	@Summary		Pegar reserva
//	@Description	Pegar reserva
//	@Tags			booking
//	@Produce		json
//	@Param			Authorization	header		string	true	"Insert you bearer token"
//	@Param			booking_id		path		string	true	"Booking ID"
//	@Success		200				{array}		usecases.GetBookingOutput
//	@Failure		422				{object}	ErrorResponse
//	@Failure		500				{object}	ErrorResponse
//	@Router			/get_booking/{booking_id} [get]
func (c *BookingController) GetBooking(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")
	output, err := c.getBooking.Execute(params)
	if err != nil {
		statusCode := http.StatusInternalServerError
		switch err.(type) {
		case *exceptions.DomainException:
			statusCode = http.StatusUnprocessableEntity
		case *exceptions.NotFoundException:
			statusCode = http.StatusNotFound
		}
		handleError(w, statusCode, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Message: message,
		Code:    statusCode,
	})
}
