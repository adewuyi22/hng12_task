package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	//"strings"
)

// Response represents the JSON response format.
type Response struct {
	Number    int      `json:"number"`
	IsPrime   bool     `json:"is_prime"`
	IsPerfect bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum  int      `json:"digit_sum"`
	FunFact   string   `json:"fun_fact"`
}

// ErrorResponse represents the JSON error response format.
type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func main() {
	http.HandleFunc("/", classifyNumberHandler)
	http.ListenAndServe(":8080", nil)
}

func classifyNumberHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Number: numberStr,
			Error:  true,
		})
		return
	}

	properties := []string{}
	if isPrime(number) {
		properties = append(properties, "prime")
	}
	if isPerfect(number) {
		properties = append(properties, "perfect")
	}
	if isArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	digitSum := sumOfDigits(number)
	funFact := getFunFact(number)

	response := Response{
		Number:    number,
		IsPrime:   isPrime(number),
		IsPerfect: isPerfect(number),
		Properties: properties,
		DigitSum:  digitSum,
		FunFact:   funFact,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	sum := 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n && n != 1
}

func isArmstrong(n int) bool {
	sum := 0
	original := n
	for n != 0 {
		digit := n % 10
		sum += digit * digit * digit
		n /= 10
	}
	return sum == original
}

func sumOfDigits(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// getFunFact returns a fun fact for the given number.
func getFunFact(n int) string {
	// Use a simple static fact for demonstration purposes.
	if n == 371 {
		return "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
	}
	return fmt.Sprintf("%d is an interesting number.", n)
}
