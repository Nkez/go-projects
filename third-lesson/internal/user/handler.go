package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"

	"golang/basic/internal/handlers"
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
		fmt.Println("AAAAAAAAAAAa")
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
	for i, c := range a {
		fmt.Println(i, c)
		if c.Id == str {
			s = append(s, c.Name)
		}
	}
	return s
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/students", h.GetInfo)
	router.GET("/students/1", h.GetById)
	router.GET("/students/grade=11", h.GetAllEleven)
	router.POST("/students/add", h.AddStudent)
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
	var todo Student
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("ADDStudent"); err != nil {
			panic(err)
		}
	}

}
