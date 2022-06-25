package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()
	router:=mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getRequestedPerson).Methods("GET")

	//Had issues with mux supporting query parameters
	//router.HandleFunc("/people", getPeopleWithFields).Queries("first_name", "{first_name:[a-zA-Z]+}", "last_name", "{last_name:[a-zA-Z]+}").Methods("GET")
	//router.HandleFunc("/people", getPeopleWithPhoneNumber).Queries("phone_number", "{phone_number:+}").Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router));
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	firstName := values.Get("first_name")
	lastName := values.Get("last_name")
	phoneNumber := values.Get("phone_number")
	if firstName != "" && lastName != "" {
		fmt.Println("Getting people with first name: %s and last name: %s", firstName, lastName)
    	people:=models.FindPeopleByName(firstName, lastName)
    	resp := []string {}
    	for _, personage := range people {
    		peep, _:=personage.ToJSON()
    		resp = append(resp, peep)
    	}
    	jsonResp, err := json.Marshal(resp)
    	if err != nil {
    		fmt.Errorf("Error happened marshaling map json for people by first and last name")
    	}
    	w.Write(jsonResp)
    	return
	} else if phoneNumber != "" {
		params, err := url.ParseQuery(r.URL.RawQuery)
		fmt.Println("Query Params: ")
		for key, value := range params {
			fmt.Printf("  %v = %v\n", key, value)
		}
		phoneNumber = strings.Replace(phoneNumber, " ", "+", 1)
		fmt.Println("Getting people with phone number:", phoneNumber)
    	people:=models.FindPeopleByPhoneNumber(phoneNumber)
    	resp := []string {}
    	for _, personage := range people {
    		peep, _:=personage.ToJSON()
    		resp = append(resp, peep)
    	}
    	jsonResp, err := json.Marshal(resp)
    	if err != nil {
    		fmt.Errorf("Error happened marshaling map json for people by first and last name")
    	}
    	w.Write(jsonResp)
	} else {
		fmt.Println("Getting people")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		resp := []string {}

		allPeople := models.AllPeople()
		for _, personage := range allPeople {
			peep, _:=personage.ToJSON()
			resp = append(resp, peep)
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Errorf("Error happened marshaling map json for all people")
		}
		w.Write(jsonResp)
		return
	}
}

func getRequestedPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	uuid, uuidErr := uuid.FromString(key)
	if uuidErr !=nil {
		fmt.Errorf("Invalid uuid: %s", key)
		w.WriteHeader(404)
		return
	}
	person, err := models.FindPersonByID(uuid)
	if err != nil || person == nil {
		fmt.Errorf("No one found with that id")
		w.WriteHeader(404)
		return
	} else {
		jsonRepresentation, err:=person.ToJSON()
		if err != nil {
			fmt.Errorf("Error happened marshaling map json for uuid: %s", key)
			w.WriteHeader(404)
		} else {
			w.WriteHeader(http.StatusOK)
			resp, _:=json.Marshal(jsonRepresentation)
			w.Write(resp)
		}
		return
	}
}
