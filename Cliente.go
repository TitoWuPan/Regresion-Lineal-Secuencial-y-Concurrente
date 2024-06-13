package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type TripData struct {
	Distance float64 `json:"distance"`
	Price    float64 `json:"Price"`
	Duration float64 `json:"duration"`
}

func main() {
	for {
		var distance, speed, duration float64

		fmt.Print("Ingrese la distancia: ")
		_, err := fmt.Scanf("%f\n", &distance)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			continue
		}

		fmt.Print("Ingrese la velocidad: ")
		_, err = fmt.Scanf("%f\n", &speed)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			continue
		}

		fmt.Print("Ingrese la duración: ")
		_, err = fmt.Scanf("%f\n", &duration)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			continue
		}

		trip := TripData{
			Distance: distance,
			Speed:    speed,
			Duration: duration,
		}

		sendTrip(trip)

		fmt.Print("¿Desea ingresar otro viaje? (s/n): ")
		var response string
		fmt.Scanf("%s\n", &response)
		if response != "s" {
			fmt.Println("Cerrando cliente...")
			os.Exit(0)
		}
	}
}

func sendTrip(trip TripData) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(&trip)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}
