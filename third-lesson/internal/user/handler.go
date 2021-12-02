package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang/basic/internal/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

var Students []Student

type Student struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

func ReadFromFile() []Student {
	js, err := os.Open("D:\\Go\\go-projects\\third-lesson\\mocks\\data.json")
	if err != nil {
		log.Fatal(err)

	}
	defer js.Close()

	data, err := ioutil.ReadAll(js)
	if err != nil {
		log.Fatal(err)
	}

	jsE := json.Unmarshal(data, &Students)
	if jsE != nil {
		log.Fatal(jsE)
	}
	return Students
}

func IsEleven() []string {
	s := make([]string, 0)
	a := ReadFromFile()
	for i, c := range a {
		fmt.Println(i, c)
		if c.Class == "11" {
			s = append(s, c.Name)
		}
	}
	return s
}

func NameById(str string) []string {
	s := make([]string, 0)
	a := ReadFromFile()
	for _, c := range a {
		if c.Id == str {
			s = append(s, c.Name)
		}
	}
	return s
}

func AddStudent() []Student {
	data := ReadFromFile()

	s := Student{"14", ":Lupa", "123"}
	data = append(data, s)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("D:\\Go\\go-projects\\third-lesson\\mocks\\data.json", dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/students", h.GetInfo)
	router.GET("/students/students?id=1", h.GetById)
	router.GET("/students/grade=11", h.GetAllEleven)
	router.GET("/students/new", h.AddStudent)
}

func (h *handler) GetInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := ReadFromFile()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}
func (h *handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := NameById("2")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}
func (h *handler) GetAllEleven(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := IsEleven()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) AddStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := AddStudent()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}
