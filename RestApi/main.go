package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
//Creacion de modelo o struc para los libros
type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Titulo string `json:"titulo"`
	Autor *Autor `json:"author"`
}
//Creacion de modelo o struc para los autor
type Autor struct {
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
}
 var books []Book
//funcion obtener todos los libros
func getBooks(w http.ResponseWriter,router*http.Request)  {
    w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}
//Obtener un solo libro
func getBook(w http.ResponseWriter,router*http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(router)
	for _,item := range books{
		if item.ID ==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Book{})
}
//crear libros
func createBook(w http.ResponseWriter,router*http.Request)  {

}
//actualizar libros
func updateBook(w http.ResponseWriter,router*http.Request)  {

}
//eliminar libros
func deleteBook(w http.ResponseWriter,router*http.Request)  {

}
func main() {
	// Inicio de enrutador
	router :=mux.NewRouter()
	books=append(books,Book{ID: "1",ISBN:"4444555",Titulo:"C Program",Autor:&Autor{Nombre:"Juan",Apellido:"Vernal"}})
	books=append(books,Book{ID: "2",ISBN:"666996",Titulo:"Java Program",Autor:&Autor{Nombre:"Tucan",Apellido:"Vernal"}})
	//Controladores/Endpoints

	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8001",router))
}
