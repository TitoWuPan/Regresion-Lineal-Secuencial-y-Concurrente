package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type TrainingData struct {
	X []float64
	Y []float64
}

func generateTrainingData(numSamples int) TrainingData {
	rand.Seed(time.Now().UnixNano())
	var data TrainingData
	for i := 0; i < numSamples; i++ {
		x := rand.Float64() * 10
		y := float64(rand.Intn(30001) + 45000)
		data.X = append(data.X, x)
		data.Y = append(data.Y, y)
	}
	return data
}

func linearRegression(X, Y []float64) (float64, float64) {
	var sumX, sumY, sumXY, sumXX float64
	n := float64(len(X))

	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumXX += X[i] * X[i]
	}

	m := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - m*sumX) / n

	return m, b
}

func linearRegressionConcurrent(X, Y []float64, wg *sync.WaitGroup, ch chan<- float64) {
	defer wg.Done()

	numBatches := 4
	batchSize := len(X) / numBatches

	results := make(chan float64, numBatches*2)

	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := start + batchSize
		if i == numBatches-1 {
			end = len(X)
		}
		wg.Add(1)
		go calculateBatch(X[start:end], Y[start:end], wg, results)
	}

	var sumM, sumB float64
	for i := 0; i < numBatches; i++ {
		sumM += <-results
		sumB += <-results
	}

	m := sumM / float64(numBatches)
	b := sumB / float64(numBatches)

	ch <- m
	ch <- b
}

func calculateBatch(X, Y []float64, wg *sync.WaitGroup, ch chan<- float64) {
	defer wg.Done()
	var sumX, sumY, sumXY, sumXX float64

	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumXX += X[i] * X[i]
	}

	ch <- sumXY
	ch <- sumXX
	ch <- sumX
	ch <- sumY
}

func Secuencial(trainingData TrainingData) time.Duration {

	start := time.Now()

	m, b := linearRegression(trainingData.X, trainingData.Y)
	end := time.Since(start)

	fmt.Println("Caso Secuencial: ")
	fmt.Println("Coeficiente m:", m)
	fmt.Println("Coeficiente b:", b)

	fmt.Println("Tiempo de respuesta:", end)

	// partido5 := 5.0
	// ventasPartido5 := m*partido5 + b
	// fmt.Println("Número de entradas vendidas para el partido número 5:", int(ventasPartido5), " \n")

	return time.Duration(end.Nanoseconds())
}

func Concurrente(trainingData TrainingData) time.Duration {
	ch := make(chan float64, 2)

	start := time.Now()
	wg.Add(1)
	go linearRegressionConcurrent(trainingData.X, trainingData.Y, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	m := <-ch
	b := <-ch
	end := time.Since(start)

	fmt.Println("Caso Concurrente: ")
	fmt.Println("Coeficiente m:", m)
	fmt.Println("Coeficiente b:", b)

	fmt.Println("Tiempo de respuesta:", end)

	// partido5 := 5.0
	// ventasPartido5 := m*partido5 + b
	// fmt.Println("Número de entradas vendidas para el partido número 5:", int(ventasPartido5), " \n")

	return time.Duration(end.Nanoseconds())
}

func main() {
	numSamples := 1000000
	trainingData := generateTrainingData(numSamples)
	var TestSecuencial []time.Duration
	var TestConcurrente []time.Duration

	for i := 0; i < 1000; i++ {
		mutex.Lock()
		var a = Secuencial(trainingData)
		mutex.Unlock()

		TestSecuencial = append(TestSecuencial, a)

		mutex.Lock()
		var b = Concurrente(trainingData)
		mutex.Unlock()

		TestConcurrente = append(TestConcurrente, b)
	}

}
