package user

import (
	"net/http"

	"github.com/nadeem-baig/MHPS-backend/config"
	"github.com/nadeem-baig/MHPS-backend/service/auth"
	"github.com/nadeem-baig/MHPS-backend/types"
	"github.com/nadeem-baig/MHPS-backend/utils"
	"github.com/nadeem-baig/MHPS-backend/utils/logger"
)

// HomeHandler responds with a welcome message.
func HomeHandler(h *config.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.JSONResponse(w, config.Response{Message: "Welcome to the Go HTTP API!"}, http.StatusOK)
	}
}

// RegisterHandler processes JSON input data and responds.
func RegisterHandler(h *config.Handler, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.RegisterUserPayload

		// Validate and parse JSON request payload
		if err := utils.ParseJson(r, &payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// check if user exists
		_, err := store.GetUserByEmail(payload.Email)
		if err == nil {
			utils.JSONResponse(w, config.Response{Message: "User already exists"}, http.StatusConflict)
			return
		}
		hashedPassword, err := auth.HashPassword(payload.Password)
		if err != nil {
			logger.Errorf("Failed to hash password")
		}
		// create user
		user := types.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  hashedPassword,
		}
		if err := store.CreateUser(user); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, config.Response{Message: "User registered successfully"}, http.StatusCreated)
	}
}

// RegisterHandler processes JSON input data and responds.
func LoginHandler(h *config.Handler, store UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.LoginUserPayload

		// Validate and parse JSON request payload
		if err := utils.ParseJson(r, &payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}

		u, err := store.GetUserByEmail(payload.Email)
		if err != nil {
			utils.JSONResponse(w, config.Response{Message: "Invalid User details"}, http.StatusBadRequest)
			return
		}

		if !auth.ComparePassword(u.Password, payload.Password) {
			utils.JSONResponse(w, config.Response{Message: "Invalid User details"}, http.StatusBadRequest)
			return
		}
		secret := []byte(config.AppConfigs.JWTSecret)
		token, err := auth.CreateJWT(secret, u.ID)
		if err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		utils.JSONResponse(w, config.Response{Message: token}, http.StatusCreated)
	}
}

// RegisterHandler processes JSON input data and responds.
func RegisterMemberHandler(h *config.Handler, store MembersStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.RegisterMemberPayload

		// Validate and parse JSON request payload
		if err := utils.ParseJson(r, &payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// validate payload
		if err := utils.Validate.Struct(payload); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusBadRequest)
			return
		}
		// check if user exists
		exists, err := store.CheckMemberExistsByAadhaar(payload.AadhaarNumber)
		if err != nil {
			utils.JSONResponse(w, config.Response{Message: "Error checking member existence"}, http.StatusInternalServerError)
			return
		} else if exists {
			utils.JSONResponse(w, config.Response{Message: "User already exists"}, http.StatusConflict)
			return
		} 


		// create user
		user := &types.Member{
			AadhaarNumber: payload.AadhaarNumber,
			Address:       payload.Address,
			BloodGroup:    payload.BloodGroup,
			ContactNumber: payload.ContactNumber,
			DateOfBirth:   payload.DateOfBirth,
			Education:     payload.Education,
			Email:         payload.Email,
			FatherName:    payload.FatherName,
			MaritalStatus: payload.MaritalStatus,
			Name:          payload.Name,
			StdPin:        payload.StdPin,
		}
		if err := store.InsertMember(user); err != nil {
			utils.JSONResponse(w, config.Response{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, config.Response{Message: "User registered successfully"}, http.StatusCreated)
	}
}
