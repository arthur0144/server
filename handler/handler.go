package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"server/service"
)

func Create(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var req CreateRequest
		if err := json.Unmarshal(reqBody, &req); err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		var resp CreateResponse
		resp.Id = s.CreateUser(req.Name, req.Age)

		data, err := json.Marshal(resp)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		response(w, http.StatusCreated, data)
	}
}

func GetAll(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := []byte(s.GetAllUsers())
		response(w, http.StatusOK, resp)
	}
}

func MakeFriends(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var req MakeFriendsRequest
		if err := json.Unmarshal(reqBody, &req); err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		name1, name2, errF := s.MakeFriends(req.TargetId, req.SourceId)
		resp := BaseResponse{
			Message: name1 + " и " + name2 + " теперь друзья",
		}
		if errF != nil {
			resp.Message = errF.Error()
			data, err := json.Marshal(resp)
			if err != nil {
				response(w, http.StatusInternalServerError, []byte(err.Error()))
				return
			}
			response(w, http.StatusBadRequest, data)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		response(w, http.StatusOK, data)
	}
}

func DeleteUser(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var req DeleteUserRequest
		if err := json.Unmarshal(reqBody, &req); err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}
		name, errF := s.DeleteUser(req.UserId)
		resp := BaseResponse{
			Message: name + " удален ",
		}
		if errF != nil {
			resp.Message = errF.Error()
			data, err := json.Marshal(resp)
			if err != nil {
				response(w, http.StatusInternalServerError, []byte(err.Error()))
				return
			}
			response(w, http.StatusBadRequest, data)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			response(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		response(w, http.StatusOK, data)
	}
}

func GetFriends(s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: раскомментировать и использовать этот id, чтобы получить друзей и написать ответ
		// id := chi.URLParam(r, "id")
		response(w, http.StatusOK, []byte("success"))
	}
}

func response(w http.ResponseWriter, statusCode int, resp []byte) {
	w.WriteHeader(statusCode)
	_, _ = w.Write(resp)
}
