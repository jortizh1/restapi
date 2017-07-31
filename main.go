package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	size = 8
	cells = 64
)

type Move struct {
	Board string `json:"board"`
	Player string `json:"player"`
}

type Response struct {
	Start string
	End string
	Message string
}

type Output struct {
	Play string `json:"play"`
}

var	StatusBoard [size][size] string
var	StatusBoardPosition [size][size] string
var OutputNote = make(map[string]Response)
var player = []string{ "w", "b", "0"}
var Id int = 0

func in_array(val string, array []string) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}

func ReadBoard (move Move, response Response){
	var validstring bool
	var z, y int = 0, 1
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			validstring = in_array(move.Board[z:y], player)
			if validstring == true	{
				StatusBoard[i][j] = move.Board[z:y]
				StatusBoardPosition[i][j] = strconv.Itoa(y)
				z++
				y++
			} else {
				response.Message = "Caracteres"
				i, j = 9, 9
			}
		}
	}
	fmt.Printf("%q \n", StatusBoard)
	fmt.Printf("%s \n", StatusBoardPosition)
}

func MakePlayWhite (Player string, response Response, flag int) Response{
	var row, col int
	var cellvalue string
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if StatusBoard[i][j] == Player{
				row = i+1
				col = j+1
				//fmt.Printf("filaXXX: %d columna: %d \n", row, col)
				if col < 0{
					continue
				}
				cellvalue = StatusBoard[row][col]
				if cellvalue == "b"{
					row = i + 2
					col = j + 2
					//fmt.Printf("filaAAA: %d columna: %d \n", row, col)
					if col < 0{
						continue
					}
					if StatusBoard[row][col] == "0" {
						response.Start = StatusBoardPosition[i][j]
						response.End = StatusBoardPosition[row][col]
						response.Message = "Come"
						fmt.Printf("Come ficha 1 %s \n", response)
						Id++
						k := strconv.Itoa(Id)
						OutputNote[k] = response
						if flag == 1{
							i, j = 8, 8
						}
					}
				} else if cellvalue == "0"{
					response.Start = StatusBoardPosition[i][j]
					response.End = StatusBoardPosition[row][col]
					response.Message = "Valido"
					fmt.Printf("Movimiento Valido 1 %s \n", response)
					Id++
					k := strconv.Itoa(Id)
					OutputNote[k] = response
				}
				row = i+1
				col = j-1
				if col < 0{
					continue
				}
				cellvalue = StatusBoard[row][col]
				if cellvalue == "b"{
					row = i + 2
					col = j - 2
					//fmt.Printf("filaAAA: %d columna: %d \n", row, col)
					if col < 0{
						continue
					}
					if StatusBoard[row][col] == "0" {
						response.Start = StatusBoardPosition[i][j]
						response.End = StatusBoardPosition[row][col]
						response.Message = "Come"
						fmt.Printf("Come ficha 2 %s \n", response)
						Id++
						k := strconv.Itoa(Id)
						OutputNote[k] = response
						if flag == 1{
							i, j = 8, 8
						}
					}
				} else if cellvalue == "0"{
					response.Start = StatusBoardPosition[i][j]
					response.End = StatusBoardPosition[row][col]
					response.Message = "Valido"
					fmt.Printf("Movimiento Valido 2 %s \n", response)
					Id++
					k := strconv.Itoa(Id)
					OutputNote[k] = response
				}
			}
			//fmt.Printf("Fila: %d - Columna: %d - Postition: %s \n", i, j, StatusBoardPosition[i][j])
		}
	}
	return response
}

func MakePlayBlack (Player string, response Response, flag int) Response{
	var row, col int
	var cellvalue string
	for i := 7; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			if StatusBoard[i][j] == Player{
				row = i - 1
				col = j - 1
				//fmt.Printf("filaXXX: %d columna: %d \n", row, col)
				if row < 0{
					continue
				}
				cellvalue = StatusBoard[row][col]
				if cellvalue == "w"{
					row = i - 2
					col = j - 2
					//fmt.Printf("filaAAA: %d columna: %d \n", row, col)
					if row < 0{
						continue
					}
					if StatusBoard[row][col] == "0" {
						response.Start = StatusBoardPosition[i][j]
						response.End = StatusBoardPosition[row][col]
						response.Message = "Come"
						fmt.Printf("Come ficha 1 %s \n", response)
						Id++
						k := strconv.Itoa(Id)
						OutputNote[k] = response
						if flag == 1{
							i, j = -2, -2
						}
					}
				} else if cellvalue == "0"{
					response.Start = StatusBoardPosition[i][j]
					response.End = StatusBoardPosition[row][col]
					response.Message = "Valido"
					fmt.Printf("Movimiento Valido 1 %s \n", response)
					Id++
					k := strconv.Itoa(Id)
					OutputNote[k] = response
				}
				row = i - 1
				col = j + 1
				if col < 0{
					continue
				}
				cellvalue = StatusBoard[row][col]
				if cellvalue == "w"{
					row = i - 2
					col = j + 2
					//fmt.Printf("filaAAA: %d columna: %d \n", row, col)
					if row < 0 || col > 7{
						continue
					}
					if StatusBoard[row][col] == "0" {
						response.Start = StatusBoardPosition[i][j]
						response.End = StatusBoardPosition[row][col]
						response.Message = "Come"
						fmt.Printf("Come ficha 2 %s \n", response)
						Id++
						k := strconv.Itoa(Id)
						OutputNote[k] = response
						if flag == 1{
							i, j = -2, -2
						}
					}
				} else if cellvalue == "0"{
					response.Start = StatusBoardPosition[i][j]
					response.End = StatusBoardPosition[row][col]
					response.Message = "Valido"
					fmt.Printf("Movimiento Valido 2 %s \n", response)
					Id++
					k := strconv.Itoa(Id)
					OutputNote[k] = response
				}
			}
			//fmt.Printf("Fila: %d - Columna: %d - Postition: %s \n", i, j, StatusBoardPosition[i][j])
		}
	}
	return response
}

func GetMoves(w http.ResponseWriter, r *http.Request) {
	var move Move
	var response Response
	var length int
	var validPlayer bool
	var outputnote []Response

	err := json.NewDecoder(r.Body).Decode(&move)
	if err != nil{
		log.Println("Error en consulta")
	}

	length =  len(move.Board)
	validPlayer = in_array(move.Player, player)
	Message := Output{}

	if length != cells{ //|| validPlayer != true{
		Message.Play = "Longitud te tablero No Valida";
	} else if validPlayer != true{
		Message.Play = "Jugador No Valido";
	} else {
		log.Printf("Board: %s", move.Board)
		log.Printf("Player: %s", move.Player)

		ReadBoard(move, response)
		if move.Player == "w" {
			response = MakePlayWhite(move.Player, response, 2)
		} else {
			response = MakePlayBlack(move.Player, response, 2)
		}
		if response.End != ""{
			for _, v:= range OutputNote {
				outputnote = append(outputnote, v)
			}
		} else {
			Message.Play = "No se Encontró Jugada";
		}
	}
	w.Header().Set("Contend-Type", "application/json")
	j, err := json.Marshal(outputnote)
	if err != nil{
		log.Println("Error en consulta")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	OutputNote = nil
	OutputNote = make(map[string]Response)
}

func GetMove(w http.ResponseWriter, r *http.Request){

	var move Move
	var response Response
	var length int
	var validPlayer bool

	err := json.NewDecoder(r.Body).Decode(&move)
	if err != nil{
		log.Println("Error en consulta")
	}

	length =  len(move.Board)
	validPlayer = in_array(move.Player, player)
	Message := Output{}

	if length != cells{ //|| validPlayer != true{
		Message.Play = "Longitud te tablero No Valida";
	} else if validPlayer != true{
		Message.Play = "Jugador No Valido";
	} else {
		log.Printf("Board: %s", move.Board)
		log.Printf("Player: %s", move.Player)

		ReadBoard(move, response)

		if move.Player == "w" {
			response = MakePlayWhite(move.Player, response, 1)
		} else {
			response = MakePlayBlack(move.Player, response, 1)
		}
		if response.End != ""{
			Message.Play = response.Start + "-" + response.End
		} else {
			Message.Play = "No se Encontró Jugada";
		}
	}
	w.Header().Set("Contend-Type", "application/json")
	j, err := json.Marshal(Message)
	if err != nil{
		log.Println("Error en consulta")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/GetMove", GetMove).Methods("POST")
	r.HandleFunc("/api/GetMoves", GetMoves).Methods("POST")

	server := &http.Server{
		Addr: ":8585",
		Handler:r,
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,
	}
	log.Println("Listening http://localhost:8585 ...")
	server.ListenAndServe()
}