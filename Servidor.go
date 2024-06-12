package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type TripData struct {
	Distance float64 `json:"distance"`
	Speed    float64 `json:"speed"`
	Duration float64 `json:"duration"`
}

var (
	tripDataChan = make(chan TripData, 100)
	wg           sync.WaitGroup
	modelUpdate  = make(chan struct{})
	beta0, beta1 float64
	allTrips     []TripData
	mu           sync.Mutex
)

func main() {
	go startServer()
	go processTrips()

	menu()
}

func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		wg.Add(1)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer wg.Done()

	var trip TripData
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&trip)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return
	}

	tripDataChan <- trip
}

func processTrips() {
	for trip := range tripDataChan {
		mu.Lock()
		allTrips = append(allTrips, trip)
		mu.Unlock()

		recalculateModel()
		modelUpdate <- struct{}{}
	}
}

func recalculateModel() {
	mu.Lock()
	defer mu.Unlock()

	var (
		n     = len(allTrips)
		sumX  float64
		sumY  float64
		sumX2 float64
		sumXY float64
	)

	for _, trip := range allTrips {
		x := trip.Distance
		y := trip.Speed
		sumX += x
		sumY += y
		sumX2 += x * x
		sumXY += x * y
	}

	if n > 1 {
		beta1 = (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
		beta0 = (sumY - beta1*sumX) / float64(n)
	}
}

func menu() {
	for {
		fmt.Println("===== Menú del Servidor =====")
		fmt.Println("1. Ver modelo actualizado")
		fmt.Println("2. Salir")
		fmt.Print("Seleccione una opción: ")

		var option int
		fmt.Scanf("%d\n", &option)

		switch option {
		case 1:
			fmt.Println(allTrips)
			showModel()
		case 2:
			fmt.Println("Cerrando el servidor...")
			close(tripDataChan)
			wg.Wait()
			os.Exit(0)
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func showModel() {
	fmt.Println("Esperando actualización del modelo...")
	select {
	case <-modelUpdate:
		mu.Lock()
		fmt.Printf("Modelo actualizado: y = %.2fx + %.2f\n", beta1, beta0)
		mu.Unlock()
	case <-time.After(5 * time.Second):
		fmt.Println("No hay nuevas actualizaciones.")
	}
}
