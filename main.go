package main 

import(
	"fmt"
	"log"
	 "encoding/json"
	 "math/rand"
	 "net/http"
	 "strconv"
	 "github.com/gorilla/mux"

)

type Movie struct{
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`

}


type Director{
  Firstname string `json:"firstname"`
  Lastname string `json:"lastname"`
}

var movies []Movie


func getMovies(w http.responseWriter, r * http.Request){
	w.Header().set("Content-Type","application/json")
	json.NewEncoder(w).encode(movies)
}




func main(){
	r := mux.NewRouter()

	movies = append(movies,Movie{ID: "1",Isbn:"221313", Title:"Movie One", Director : &Director{Firstname:"John", Lastname:"Doe"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenandServe(":8000",r))
}

