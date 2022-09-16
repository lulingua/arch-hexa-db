package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lulingua/hexa-arch/application/ports"
	"github.com/lulingua/hexa-arch/domain"
)

type GorillaHandler struct {
	service   ports.PersonService
	MuxRouter *mux.Router
}

func NewGorillaHandler(s ports.PersonService) *GorillaHandler {
	r := mux.NewRouter()

	return &GorillaHandler{
		service:   s,
		MuxRouter: r,
	}
}

func (h *GorillaHandler) RunServer(port string) {
	if port == "default" {
		port = ":8888"
	}

	log.Println("Server listining on port", port)
	log.Fatalln(http.ListenAndServe(port, h.MuxRouter))
}

func (h *GorillaHandler) SetupRoutes() {
	h.MuxRouter.HandleFunc("/person", h.ListPerson).Methods("GET")
	h.MuxRouter.HandleFunc("/person", h.AddPerson).Methods("POST")
}

type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type HttpHandler struct {
	Persons []domain.Person
}

// POST
func (g GorillaHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPerson domain.Person
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseInfo{
			Status: http.StatusInternalServerError,
			Data:   err,
		})
		return
	}

	g.service.AddPerson(newPerson)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusCreated,
		Data:   newPerson,
	})
}

// GET ALL

func (g GorillaHandler) ListPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   g.service.ListPerson(),
	})
}

// GET PERSON BY NAME

/*func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	nameParam := param["name"]

	var newPerson domain.Person
	for _, person := range slicedb.Persons {
		if person.Name == (nameParam) {
			newPerson = person
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseInfo{
		Status: http.StatusOK,
		Data:   newPerson,
	})
}
*/
