package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Measurement struct {
	Min float64
	Max float64
	Sum float64
	Count int64
}

func main() {
	start := time.Now()

	measurementsChan := carregarMedicoes("measurements.txt")

	datas := processarMedicoes(measurementsChan)

	imprimirResultados(datas)

	elapsed := time.Since(start)
	fmt.Println("Tempo de execução:", elapsed)
}

func carregarMedicoes(nomeArquivo string) <-chan string {
	measurementsChan := make(chan string, 1024) // Buffer

	go func() {
		defer close(measurementsChan)

		arquivo, err := os.Open(nomeArquivo)
		if err != nil {
			panic(err)
		}
		defer arquivo.Close()

		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			measurementsChan <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	return measurementsChan
}

func processarMedicoes(measurementsChan <-chan string) map[string]Measurement {
	datas := make(map[string]Measurement)
	var mu sync.Mutex // Mutex para garantir acesso seguro ao mapa

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ { // número de workers
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rawData := range measurementsChan {
				semicolon := strings.Index(rawData, ";")
				location := rawData[:semicolon]
				rawTemp := rawData[semicolon+1:]

				temp, _ := strconv.ParseFloat(rawTemp, 64)

				mu.Lock()
				measurement, ok := datas[location]
				if !ok {
					measurement = Measurement{
						Min:   temp,
						Max:   temp,
						Sum:   temp,
						Count: 1,
					}
				} else {
					measurement.Min = math.Min(measurement.Min, temp)
					measurement.Max = math.Max(measurement.Max, temp)
					measurement.Sum += temp
					measurement.Count++
				}
				datas[location] = measurement
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	return datas
}

func imprimirResultados(datas map[string]Measurement) {
	locations := make([]string, 0, len(datas))
	for name := range datas {
		locations = append(locations, name)
	}
	sort.Strings(locations)

	for _, name := range locations {
		measurement := datas[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Println()
}

// func generateDataTemp(){
// 	cts := []string{"Hamburg", "Bulawayo", "Palembang", "St. John's", "Cracow", "Rio de Janeiro", "Tokyo", "Sydney", "Cairo", "London"}

// 	file, err := os.Create("measurements.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// buffer
// 	w := bufio.NewWriter(file)
// 	defer w.Flush()

// 	// billion rows
// 	for i := 0; i < 1000000000; i++ {
// 		ct := cts[rand.Intn(len(cts))]
// 		temp := rand.Float64()*55 - 10

// 		line := fmt.Sprintf("%s;%.1f\n", ct, temp)
// 		w.WriteString(line)
// 	}
// }