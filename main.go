package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	// Echo: Return the matrix as a string in matrix format.
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		records := readFile(w, r)
		var response string

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}

		fmt.Fprint(w, response)
	})

	// Invert: Return the matrix as a string in matrix format where the columns and rows are inverted
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		records := readFile(w, r)
		var response []string

		for i, row := range records {
			var tmp []string
			for j := range row {
				tmp = append(tmp, records[j][i])
			}
			response = append(response, strings.Join(tmp, ","))
		}

		fmt.Fprint(w, strings.Join(response, "\n"))
	})

	// Flatten: Return the matrix as a 1 line string, with values separated by commas.
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		records := readFile(w, r)
		var response []string

		for _, row := range records {
			response = append(response, strings.Join(row, ","))
		}

		fmt.Fprint(w, strings.Join(response, ","))
	})

	// Sum: Return the sum of the integers in the matrix
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		records := readFile(w, r)
		var response int = 0

		for i, row := range records {
			for j, _ := range row {
				cell, err := strconv.Atoi(records[i][j])
				if err != nil {
					fmt.Fprintf(w, "error 3: %s", err.Error())
					return
				}
				response += cell
			}
		}

		fmt.Fprint(w, response)
	})

	// Multiply: Return the product of the integers in the matrix
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		records := readFile(w, r)
		var response int = 1

		for i, row := range records {
			for j, _ := range row {
				cell, err := strconv.Atoi(records[i][j])
				if err != nil {
					fmt.Fprintf(w, "error 3: %s", err.Error())
					return
				}
				response *= cell
			}
		}

		fmt.Fprint(w, response)
	})

	http.ListenAndServe(":8080", nil)
}

func readFile(w http.ResponseWriter, r *http.Request) [][]string {
	file, _, err := r.FormFile("file")

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error 1: %s", err.Error())))
		return nil
	}

	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error 2: %s", err.Error())))
		return nil
	}

	return records
}
