package application

import (
	"encoding/json"
	"errors"
	"finalTask/pkg/calculation"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		errorMessage := fmt.Sprintf("{err: %s}", err.Error())
		if errors.Is(err, calculation.ErrInvalidExpression) ||
			errors.Is(err, calculation.ErrUnbalancedParentheses) ||
			errors.Is(err, calculation.ErrDivisionByZero) ||
			errors.Is(err, calculation.ErrInvalidTypeOfOperation) {
			http.Error(w, errorMessage, http.StatusUnprocessableEntity)
		} else {
			http.Error(w, errorMessage, http.StatusInternalServerError)
		}

	} else {
		fmt.Fprintf(w, "{result: %f}", result)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s %s",
			r.Method,
			r.URL.Hostname(),
			r.URL.Path,
			time.Since(start))
	})
}

func (a *Application) RunServer() error {
	mux := http.NewServeMux()
	handler := http.HandlerFunc(CalcHandler)
	mux.Handle("/api/v1/calculate", loggingMiddleware(handler))
	return http.ListenAndServe(":"+a.config.Addr, mux)
}
