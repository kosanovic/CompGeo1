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
				var isTrue = areIntercepting2(graphs[i], graphs[j])
				if isTrue {
					amount++
				}
			}
		}
	}
	return amount
}

func areIntercepting2(graph1 Graph, graph2 Graph) bool {
	// if ccw < 0 -> punkt liegt rechts von g1
	// if ccw = 0 -> punkt liegt auf g1
	// if ccw > 0 -> punkt liegt links von g1
	// aus sicht von g1

	p := Point{graph1.p1X, graph1.p2X} // P (erster Punkt g1)
	q := Point{graph1.p2X, graph1.p2Y}
	r1 := Point{graph2.p1X, graph2.p1Y}
	r2 := Point{graph2.p2X, graph2.p2Y}

	ccw1 := p.X*q.Y + p.Y*r1.X + q.X*r1.Y - q.Y*r1.X - p.X*r1.Y - p.Y*q.X
	ccw2 := p.X*q.Y + p.Y*r2.X + q.X*r2.Y - q.Y*r2.X - p.X*r2.Y - p.Y*q.X

	if ccw1 < 0 {
		// r1 liegt rechts
		// liegt r2 recht, auf der strecke, oder links?
		if ccw2 < 0 {
			// r1 und r2 liegen rechts -> strecken schneiden sich nicht
			return false
		} else if ccw2 == 0 {
			// r1 liegt rechts und r2 liegt auf der strecke -> sonderfall
			return true
		} else {
			// r1 liegt rechts und r2 liegt links -> Schnittpunkt
			return true
		}
	} else if ccw1 == 0 {
		// r liegt auf der strecke
		// liegt r2 recht, auf der strecke, oder links?
		if ccw2 < 0 {
			// r1 liegt auf der strecke und r2 liegt rechts -> sonderfall
			return true
		} else if ccw2 == 0 {
			// r1 und r2 liegen auf der strecke -> sonderfall
			return true
		} else {
			// r1 liegt auf Strecke r2 liegt links -> sonderfall
			return true
		}
	} else {
		// r liegt auf der linken Seite
		if ccw2 < 0 {
			// r1 liegt links, r2 liegt rechts -> Schnittpunkt
			return true
		} else if ccw2 == 0 {
			// r1 liegt links, r2 liegt auf der Strecke -> sonderfall
			return true
		} else {
			// r1 und r2 liegen links -> Kein Schnittpunkt
			return false
		}
	}
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
