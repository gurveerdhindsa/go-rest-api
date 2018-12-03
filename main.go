package main
import (
  "github.com/gorilla/mux"
  "encoding/json"
  "log"
  "net/http"
)

var people [] Person  // Mocked array

func main() {
  InitializeMockData()
  router := mux.NewRouter()

  //  Routes
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", AddPerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

  //  Start HTTP server at localhost:8000
  log.Fatal(http.ListenAndServe(":8000", router))
}

/*
  initializeMockData
  Mock data into people array
*/
func InitializeMockData() {
  people = append(people, Person{ID: "1", Firstname: "Bob", Lastname: "Doe", Address: &Address{City: "Ottawa"}})
  people = append(people, Person{ID: "2", Firstname: "John", Lastname: "Smith", Address: &Address{City: "Mississauga"}})
  people = append(people, Person{ID: "3", Firstname: "Mickey", Lastname: "Mouse"})
}

/*
  GetPeople - get all people
  /people (GET)
*/
func GetPeople(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(people)
}

/*
  GetPerson - get all people with a specified ID
  params: id
  /people/{id} (GET)
*/
func GetPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

/*
  AddPerson - add a person with a specified ID
  params: id
  /people/{id} (POST)
*/
func AddPerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}

/*
  DeletePerson - delete a person with a specified ID
  params: id
  /people/{id} (POST)
*/
func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for i, item := range people {
    if item.ID == params["id"] {
      people = append(people[:i], people[i+1:]...)
      break
    }
    json.NewEncoder(w).Encode(people)
  }
}

/*
  Person
*/
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

/*
  Address
*/
type Address struct {
    City  string `json:"city,omitempty"`
}
