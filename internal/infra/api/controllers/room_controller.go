package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julianojj/fastrip/internal/core/exceptions"
	"github.com/julianojj/fastrip/internal/core/usecases"
)

type RoomController struct {
	registerRoom *usecases.RegisterRoom
	getRooms     *usecases.GetRooms
}

func NewRoomController(
	registerRoom *usecases.RegisterRoom,
	getRooms *usecases.GetRooms,
) *RoomController {
	return &RoomController{
		registerRoom,
		getRooms,
	}
}

// RegisterRoom godoc
//
//	@Summary		Registra novo quarto
//	@Description	Registra novo quarto
//	@Tags			room
//	@Accept			json
//	@Produce		json
//	@Param			input	body		usecases.RegisterRoomInput	true	"RegisterRoomInput"
//	@Success		201		{array}		usecases.RegisterRoomOutput
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		422		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/register_room [post]
func (c *RoomController) RegisterRoom(w http.ResponseWriter, r *http.Request) {
	var input *usecases.RegisterRoomInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	output, err := c.registerRoom.Execute(input)
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

func (c *RoomController) GetRooms(w http.ResponseWriter, r *http.Request) {
	output, err := c.getRooms.Execute()
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
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&output); err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to encode response", err)
		return
	}
}
