package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func mean(data []float64, result chan<- float64) {
	total := 0.0
	for _, val := range data {
		total += val
	}
	result <- total / float64(len(data))
}

func covariance(x, y []float64, meanX, meanY float64, result chan<- float64) {
	cov := 0.0
	for i := 0; i < len(x); i++ {
		cov += (x[i] - meanX) * (y[i] - meanY)
	}
	result <- cov / float64(len(x))
}

func variance(data []float64, meanVal float64, result chan<- float64) {
	varn := 0.0
	for _, val := range data {
		varn += math.Pow(val-meanVal, 2)
	}
	result <- varn / float64(len(data))
}

func linearRegression(x, y []float64) (float64, float64) {
	meanXChan := make(chan float64)
	meanYChan := make(chan float64)
	covChan := make(chan float64)
	varXChan := make(chan float64)
	varYChan := make(chan float64)

	go mean(x, meanXChan)
	go mean(y, meanYChan)

	meanX := <-meanXChan
	meanY := <-meanYChan

	go variance(x, meanX, varXChan)
	go variance(y, meanY, varYChan)

	varX := <-varXChan
	//varY := <-varYChan

	go covariance(x, y, meanX, meanY, covChan)

	covXY := <-covChan

	slope := covXY / varX
	intercept := meanY - slope*meanX

	return intercept, slope
}

func main() {
	urls := []string{
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_1.csv",
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_2.csv",
		"https://raw.githubusercontent.com/TitoWuPan/Regresion-Lineal-Secuencial-y-Concurrente/TA3/Dataset/NYC_SPEED_3.csv",
	}

	combinedData := [][]float64{}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error al descargar el archivo:", err)
			return
		}
		defer resp.Body.Close()

		reader := csv.NewReader(resp.Body)
		reader.FieldsPerRecord = -1

		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error al parsear el archivo CSV:", err)
			return
		}

		for _, row := range records {
			var floatRow []float64
			for _, value := range row {
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}
				floatRow = append(floatRow, floatValue)
			}
			combinedData = append(combinedData, floatRow)
		}
	}

	fmt.Println(len(combinedData))

	xData := make([]float64, len(combinedData))
	yData := make([]float64, len(combinedData))

	for i, row := range combinedData {
		// fmt.Println(row)
		for j, value := range row {
			if j == 1 {
				yData[i] = value // distance - Variable independiente
			}
			if j == 2 {
				xData[i] = value // speed - Variable dependiente
			}
		}
	}

	intercept, slope := linearRegression(xData, yData)

	fmt.Printf("Intercepto: %.4f\n", intercept)
	fmt.Printf("Pendiente: %.4f\n", slope)
}
