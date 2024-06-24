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
	Speed    float64 `json:"speed"`
	Duration float64 `json:"duration"`
}

func main() {
	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Ingresar datos manualmente")
		fmt.Println("2. Salir")
		fmt.Print("Opción: ")
		var option int
		fmt.Scanf("%d\n", &option)

		switch option {
		case 1:
			enterManualData()
		case 2:
			fmt.Println("Cerrando cliente...")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func enterManualData() {
	var trips []TripData

	for {
		var distance, speed, duration float64

		fmt.Print("Ingrese la distancia: ")
		_, err := fmt.Scanf("%f\n", &distance)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			return
		}

		fmt.Print("Ingrese la velocidad: ")
		_, err = fmt.Scanf("%f\n", &speed)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			return
		}

		fmt.Print("Ingrese la duración: ")
		_, err = fmt.Scanf("%f\n", &duration)
		if err != nil {
			fmt.Println("Error de entrada. Intente de nuevo.")
			return
		}

		trip := TripData{
			Distance: distance,
			Speed:    speed,
			Duration: duration,
		}

		trips = append(trips, trip)

		sendTrips(trips)

		fmt.Print("¿Desea ingresar otro viaje? (s/n): ")
		var response string
		fmt.Scanf("%s\n", &response)
		if response != "s" {
			break
		}
	}
}

func sendTrips(trips []TripData) {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(trips)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}
