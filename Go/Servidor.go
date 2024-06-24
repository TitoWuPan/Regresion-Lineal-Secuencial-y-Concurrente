package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type TripData struct {
	Distance float64 `json:"distance"`
	Speed    float64 `json:"speed"`
	Duration float64 `json:"duration"`
}

type Model struct {
	Beta0 float64 `json:"beta0"`
	Beta1 float64 `json:"beta1"`
}

var (
	combinedData []TripData
	beta0, beta1 float64
	mu           sync.Mutex
)

func mean(data []float64) float64 {
	total := 0.0
	for _, val := range data {
		total += val
	}
	return total / float64(len(data))
}

func covariance(x, y []float64, meanX, meanY float64) float64 {
	cov := 0.0
	for i := 0; i < len(x); i++ {
		cov += (x[i] - meanX) * (y[i] - meanY)
	}
	return cov / float64(len(x))
}

func variance(data []float64, meanVal float64) float64 {
	varn := 0.0
	for _, val := range data {
		varn += math.Pow(val-meanVal, 2)
	}
	return varn / float64(len(data))
}

func linearRegression(x, y []float64) (float64, float64) {
	meanX := mean(x)
	meanY := mean(y)

	varX := variance(x, meanX)
	covXY := covariance(x, y, meanX, meanY)

	slope := covXY / varX
	intercept := meanY - slope*meanX

	return intercept, slope
}

func readCSVData(urls []string) ([]TripData, error) {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error al descargar el archivo: %v", err)
		}
		defer resp.Body.Close()

		reader := csv.NewReader(resp.Body)
		reader.FieldsPerRecord = -1

		records, err := reader.ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error al parsear el archivo CSV: %v", err)
		}

		for _, row := range records {
			var trip TripData
			distance, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				continue
			}
			speed, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				continue
			}
			duration, err := strconv.ParseFloat(row[0], 64)
			if err != nil {
				continue
			}
			trip = TripData{
				Distance: distance,
				Speed:    speed,
				Duration: duration,
			}
			combinedData = append(combinedData, trip)
		}
	}

	return combinedData, nil
}

func main() {
	urls := []string{
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_1.csv",
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_2.csv",
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_3.csv",
	}

	combinedData, err := readCSVData(urls)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	xData := make([]float64, len(combinedData))
	yData := make([]float64, len(combinedData))

	for i, trip := range combinedData {
		xData[i] = trip.Speed
		yData[i] = trip.Distance
	}

	beta1, beta0 = linearRegression(xData, yData)

	fmt.Printf("Speed = %.4f + (%.4f * Distance)\n", beta1, beta0)

	mux := http.NewServeMux()
	mux.HandleFunc("/model", modelHandler)
	mux.HandleFunc("/trips", tripsHandler)

	corsHandler := setupCORS(mux)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", corsHandler))
	}()

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error al abrir puerto:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor escuchando en puerto 9090...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			return
		}

		go handleClient(conn)
	}
}

func modelHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	model := Model{
		Beta0: beta0,
		Beta1: beta1,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model)
}

func tripsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(combinedData)
}

func setupCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var trips []TripData
	err := decoder.Decode(&trips)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, trip := range trips {
		combinedData = append(combinedData, trip)
	}

	xData := make([]float64, len(combinedData))
	yData := make([]float64, len(combinedData))

	for i, trip := range combinedData {
		xData[i] = trip.Speed
		yData[i] = trip.Distance
	}

	beta1, beta0 = linearRegression(xData, yData)

	fmt.Printf("Speed = %.4f + (%.4f * Distance)\n", beta1, beta0)
}
