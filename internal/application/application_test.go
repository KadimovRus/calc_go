package application

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	// Проверка успешного запроса

	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult string
	}{
		{
			name:           "simple",
			expression:     `{"expression": "1+1"}`,
			expectedResult: "result: 2.000000",
		},
		{
			name:           "priority",
			expression:     `{"expression": "(2+2)*2"}`,
			expectedResult: "result: 8.000000",
		},
		{
			name:           "priority",
			expression:     `{"expression": "2+2*2"}`,
			expectedResult: "result: 6.000000",
		},
		{
			name:           "/",
			expression:     `{"expression": "1/2"}`,
			expectedResult: "result: 0.500000",
		},
	}

	for _, tc := range testCasesSuccess {
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tc.expression))
		w := httptest.NewRecorder()

		CalcHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		expectedBody := tc.expectedResult
		if w.Body.String() != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
		}
	}

	testCasesUnsuccessfull := []struct {
		name           string
		expression     string
		expectedResult string
	}{
		{
			name:           "simple",
			expression:     `{"expression": "1+1/"}`,
			expectedResult: "err: expression is not valid\n",
		},
		{
			name:           "priority",
			expression:     `{"expression": "(2+2)/0"}`,
			expectedResult: "err: division by zero\n",
		},
		{
			name:           "priority",
			expression:     `{"expression": "2+2*2)"}`,
			expectedResult: "err: unbalanced parentheses\n",
		},
	}

	for _, tc := range testCasesUnsuccessfull {
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tc.expression))
		w := httptest.NewRecorder()

		CalcHandler(w, req)

		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
		}

		expectedBody := tc.expectedResult
		if w.Body.String() != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
		}
	}
}
