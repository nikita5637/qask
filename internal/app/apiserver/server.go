package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"qask/internal/app/model"
	"qask/internal/app/questions"
	"qask/internal/app/store"

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
	s.router.HandleFunc("/questions", s.handleQuestionsGet()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUsersPost()).Methods("POST")
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

		s.respond(w, r, http.StatusCreated, "")
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
