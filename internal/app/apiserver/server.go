package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"qask/internal/app/model"
	"qask/internal/app/questions"
	"qask/internal/app/store"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	store     store.Store
	questions questions.Questions
	logger    *logrus.Logger
	router    *mux.Router
}

func newServer(store store.Store, questions questions.Questions) *server {
	s := &server{
		store:     store,
		questions: questions,
		logger:    logrus.New(),
		router:    mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/questions", s.handleQuestionsGet()).Methods("GET")
	s.router.HandleFunc("/reports", s.handleReportsPost()).Methods("POST")
	s.router.HandleFunc("/users", s.handleUsersGet()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUsersPost()).Methods("POST")
	s.router.HandleFunc("/users/{id:[0-9]+}", s.handleGetUser()).Methods("GET")
}

func (s *server) logRequest(next http.Handler) http.Handler {
	type responseWriter struct {
		http.ResponseWriter
		code int
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started request %s %s", r.RequestURI, r.Method)

		start := time.Now()
		newResponseWriter := responseWriter{w, http.StatusOK}

		next.ServeHTTP(newResponseWriter, r)

		logger.Infof("result code = %d in = %f sec", newResponseWriter.code, time.Now().Sub(start).Seconds())
	})
}

func (s *server) handleGetUser() http.HandlerFunc {
	type request struct {
		From string `json:"from"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Empty body"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.From != "telegram" {
			s.error(w, r, http.StatusBadRequest, errors.New("Unknown From"))
			return
		}

		vars := mux.Vars(r)
		strID := vars["id"]
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Invalid user id"))
			return
		}

		user := s.store.User().FindUserByID(id)
		if user == nil {
			s.error(w, r, http.StatusNotFound, errors.New("User not found"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		s.respond(w, r, http.StatusOK, user)
	})
}

func (s *server) handleQuestionsGet() http.HandlerFunc {
	type request struct {
		TgID int64  `json:"tgID"`
		From string `json:"from"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.From != "telegram" {
			s.error(w, r, http.StatusBadRequest, errors.New("Unknown From"))
			return
		}

		q, err := s.questions.Questions().GetQuestion()
		for {
			if err == nil {
				break
			}

			s.logger.Infof("Get invalid question")
			q, err = s.questions.Questions().GetQuestion()
		}

		w.Header().Add("Content-Type", "application/json")
		s.respond(w, r, http.StatusOK, q)
	})
}

func (s *server) handleReportsPost() http.HandlerFunc {
	type request struct {
		model.User
		From    string `json:"from"`
		Message string `json:"message"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Empty body"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.From != "telegram" {
			s.error(w, r, http.StatusBadRequest, errors.New("Unknown From"))
			return
		}

		if req.From == "telegram" {
			user := s.store.User().FindUserByTgID(req.TgID)
			if user == nil {
				s.error(w, r, http.StatusBadRequest, errors.New("User not found"))
				return
			}

			if err := s.store.Report().CreateReport(int64(user.ID), req.Message); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.respond(w, r, http.StatusCreated, "{}")
		}

		s.respond(w, r, http.StatusBadRequest, "{}")
	})
}

func (s *server) handleUsersGet() http.HandlerFunc {
	type request struct {
		From string `json:"from"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Empty body"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.From != "telegram" {
			s.error(w, r, http.StatusBadRequest, errors.New("Unknown From"))
			return
		}

		users := s.store.User().GetUsers()

		w.Header().Add("Content-Type", "application/json")
		s.respond(w, r, http.StatusOK, users)
	})
}

func (s *server) handleUsersPost() http.HandlerFunc {
	type request struct {
		model.User
		From string `json:"from"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Empty body"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.From != "telegram" {
			s.error(w, r, http.StatusBadRequest, errors.New("Unknown From"))
			return
		}

		newUser := model.User{}
		newUser.FirstName = req.FirstName
		newUser.UserName = req.UserName
		newUser.TgID = req.TgID

		if err := s.store.User().CreateUser(&newUser); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, newUser)
	})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
