package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var graphsFirstFile = dataLoading("s_1000_1.dat")
	var amount = amountOfInterceptingGraphs(graphsFirstFile)

	fmt.Print(amount)
}

func dataLoading(filename string) []Graph {

	// Open the .dat file
	file, err := os.Open("strecken/" + filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	var graphs []Graph

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into fields
		fields := strings.Fields(line)

		// Convert the fields to float64 values
		var values []float64
		for _, field := range fields {
			value, err := strconv.ParseFloat(field, 64)
			if err != nil {
				fmt.Println("Error parsing float value:", err)
			}
			values = append(values, value)
		}

		graphs = append(graphs, Graph{p1X: values[0], p1Y: values[1], p2X: values[2], p2Y: values[3]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return graphs
}

func amountOfInterceptingGraphs(graphs []Graph) int {
	var amount int
	for i := 0; i < len(graphs); i++ {
		for j := 0; j < len(graphs); j++ {
			if i != j {
				var isTrue = areIntercepting(graphs[i], graphs[j])
				if isTrue {
					amount++
				}
			}
		}
	}
	return amount
}

func areIntercepting(graph1, graph2 Graph) bool {
	m1 := getGraphGradient(graph1)
	m2 := getGraphGradient(graph2)

	b1 := getB(m1, graph1.p1X, graph1.p1Y)
	b2 := getB(m2, graph2.p1X, graph2.p1Y)

	x := (b2 - b1) / (m1 - m2)

	if (x >= graph1.p1X && x <= graph1.p2X) && (x >= graph2.p1X && x <= graph2.p2X) {
		return true
	} else {
		return false
	}
}

func getGraphGradient(graph Graph) float64 {
	return (graph.p2Y - graph.p1Y) / (graph.p2X - graph.p1X)
}

func getB(gradient, x, y float64) float64 {
	return y - gradient*x
}

type Graph struct {
	p1X float64
	p1Y float64
	p2X float64
	p2Y float64
}