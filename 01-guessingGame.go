//Eoghan O'Connor
//Go Problem sheet 2
//01-guessingGame.go

package main

import(
	"fmt" // Formatted I/O
	"net/http" // Import For ussing http
)

func guessingGame(w http.ResponseWriter, r *http.Request){

		fmt.Fprintf(w,"Guessing Game")
}

func main(){

	http.HandleFunc("/",guessingGame)
	http.ListenAndServe(":8080",nil)
}

