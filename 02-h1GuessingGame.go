//Eoghan O'Connor
//Go Problem sheet 2
//02-h1GuessingGame.go

package main

import(

	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")// makes browser render all html tags 

	fmt.Fprintf(w,"<h1>Guessing Game</h1>")
	
}

func main(){

	http.HandleFunc("/",requestHandler)
	http.ListenAndServe(":8080",nil)

}