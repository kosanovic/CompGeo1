package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Richtige Werte : 11, 732, 77126
	var graphsFirstFile = dataLoading("s_1000_1.dat")
	start := time.Now()
	// Call your function
	var amount1 = amountOfInterceptingGraphs(graphsFirstFile)
	// Get the time again and calculate the duration
	duration := time.Since(start)
	fmt.Println("In the first data set the amount of crossing graphs is ", amount1)
	// Print the duration
	fmt.Println("Time taken for first calculation:", duration)

	var graphsSecondFile = dataLoading("s_10000_1.dat")
	start2 := time.Now()
	// Call your function
	var amount2 = amountOfInterceptingGraphs(graphsSecondFile)
	// Get the time again and calculate the duration
	duration2 := time.Since(start2)
	fmt.Println("In the second data set the amount of crossing graphs is ", amount2)
	// Print the duration
	fmt.Println("Time taken for second calculation:", duration2)

	var graphsThirdFile = dataLoading("s_100000_1.dat")
	start3 := time.Now()
	// Call your function
	var amount3 = amountOfInterceptingGraphs(graphsThirdFile)
	// Get the time again and calculate the duration
	duration3 := time.Since(start3)
	fmt.Println("In the third data set the amount of crossing graphs is ", amount3)
	// Print the duration
	fmt.Println("Time taken for third calculation:", duration3)

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
		for j := i; j < len(graphs); j++ {
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

func areIntercepting(graph1 Graph, graph2 Graph) bool {
	// Bestimme die Punkte P und Q von Graph 1
	p1 := Point{graph1.p1X, graph1.p1Y}
	p2 := Point{graph1.p2X, graph1.p2Y}
	// Bestimme die Punkte R1 und R2 von Graph 2
	q1 := Point{graph2.p1X, graph2.p1Y}
	q2 := Point{graph2.p2X, graph2.p2Y}

	if ccw(p1, p2, q1) == 0 && ccw(p1, p2, q2) == 0 {
		return isPointOnLine(p1, p2, q1) || isPointOnLine(p1, p2, q2)
	} else if ccw(p1, p2, q1)*ccw(p1, p2, q2) <= 0 && ccw(q1, q2, p1)*ccw(q1, q2, p2) <= 0 {
		return true
	}
	return false
}

func isPointOnLine(p1, p2, q Point) bool {
	// Überprüfung, ob q auf der Strecke p1-p2 liegt
	if (q.X >= p1.X && q.X <= p2.X) || (q.X >= p2.X && q.X <= p1.X) {
		return true
	}

	return false
}
func ccw(p1, p2, p3 Point) float64 {
	return (p2.Y-p1.Y)*(p3.X-p2.X) - (p2.X-p1.X)*(p3.Y-p2.Y)
}

type Graph struct {
	p1X float64
	p1Y float64
	p2X float64
	p2Y float64
}

type Point struct {
	X, Y float64
}
