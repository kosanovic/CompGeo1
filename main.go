package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the .dat file
	file, err := os.Open("strecken/s_1000_1.dat")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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
				return
			}
			values = append(values, value)
		}

		graphs = append(graphs, Graph{p1X: values[0], p1Y: values[1], p2X: values[2], p2Y: values[3]})
	}

	fmt.Println(graphs)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

type Graph struct {
	p1X float64
	p1Y float64
	p2X float64
	p2Y float64
}
