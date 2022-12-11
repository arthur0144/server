package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"server/service"
)

func Create(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := readBody(r)
		if err != nil {
			response500(w, err)
			return
		}

		var req CreateRequest
		err = json.Unmarshal(reqBody, &req)
		if err != nil {
			response500(w, err)
			return
		}

		var resp CreateResponse
		resp.Id = s.CreateUser(req.Name, req.Age)
		response(w, http.StatusCreated, resp)
	}
}

func GetAll(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp BaseResponse
		resp.Message = s.GetAllUsers()
		response(w, http.StatusOK, resp)
	}
}

func MakeFriends(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := readBody(r)
		if err != nil {
			response500(w, err)
			return
		}

		var req MakeFriendsRequest
		if err := json.Unmarshal(reqBody, &req); err != nil {
			response500(w, err)
			return
		}

		name1, name2, err := s.MakeFriends(req.TargetId, req.SourceId)
		if err != nil {
			response400(w, err)
			return
		}

		var resp BaseResponse
		resp.Message = name1 + " и " + name2 + " теперь друзья"
		response(w, http.StatusOK, resp)
	}
}

func DeleteUser(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := readBody(r)
		if err != nil {
			response500(w, err)
			return
		}

		var req DeleteUserRequest
		if err := json.Unmarshal(reqBody, &req); err != nil {
			response500(w, err)
			return
		}

		name, err := s.DeleteUser(req.UserId)
		if err != nil {
			response400(w, err)
			return
		}

		var resp BaseResponse
		resp.Message = name + " удален"
		response(w, http.StatusOK, resp)
	}
}

func GetFriends(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response500(w, err)
			return
		}

		var resp BaseResponse
		resp.Message, err = s.GetAllUserFriends(req)
		if err != nil {
			response400(w, err)
			return
		}
		response(w, http.StatusOK, resp)
	}
}

func UpdateAge(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := readBody(r)
		if err != nil {
			response500(w, err)
			return
		}

		req, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response500(w, err)
			return
		}

		var age UpdateAgeRequest
		if err := json.Unmarshal(reqBody, &age); err != nil {
			response500(w, err)
			return
		}

		var resp BaseResponse
		resp.Message, err = s.UpdateAge(req, age.UserAge)
		if err != nil {
			response400(w, err)
			return
		}
		response(w, http.StatusOK, resp)
	}
}
