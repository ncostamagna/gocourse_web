package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
	}
)

func MakeEndpoints(s Service) Endpoints {

	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}

}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "invalid request format"})
			return
		}

		if req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "first name is required"})
			return
		}

		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "last name is required"})
			return
		}

		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
	}
}
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.GetAll()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: users})
	}
}
func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]
		user, err := s.Get(id)
		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "invalid request format"})
			return
		}

		if req.FirstName != nil && *req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "first name is required"})
			return
		}

		if req.LastName != nil && *req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "last name is required"})
			return
		}

		path := mux.Vars(r)
		id := path["id"]

		if err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "success"})
	}
}

func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.Vars(r)
		id := path["id"]

		if err := s.Delete(id); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "success"})
	}
}
