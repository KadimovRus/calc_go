package application

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcHandlerSuccess(t *testing.T) {
	// Проверка успешного запроса

	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult string
	}{
		{
			name:           "simple",
			expression:     `{"expression": "1+1"}`,
			expectedResult: "{result: 2.000000}",
		},
		{
			name:           "priority",
			expression:     `{"expression": "(2+2)*2"}`,
			expectedResult: "{result: 8.000000}",
		},
		{
			name:           "priority",
			expression:     `{"expression": "2+2*2"}`,
			expectedResult: "{result: 6.000000}",
		},
		{
			name:           "/",
			expression:     `{"expression": "1/2"}`,
			expectedResult: "{result: 0.500000}",
		},
	}

	for _, tc := range testCasesSuccess {
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tc.expression))
		w := httptest.NewRecorder()

		CalcHandler(w, req)

		assert.Equal(t, w.Code, http.StatusOK, "they should be equal")
		assert.Equal(t, w.Body.String(), tc.expectedResult, "they should be equal")
	}
}

func TestCalcHandlerUnsuccessful(t *testing.T) {

	testCasesUnsuccessful := []struct {
		name           string
		expression     string
		expectedResult string
	}{
		{
			name:           "expressionNotValid",
			expression:     `{"expression": "1+1/"}`,
			expectedResult: "{err: expression is not valid}\n",
		},
		{
			name:           "divisionByZero",
			expression:     `{"expression": "(2+2)/0"}`,
			expectedResult: "{err: expression is not valid}\n",
		},
		{
			name:           "unbalancedParentheses",
			expression:     `{"expression": "2+2*2)"}`,
			expectedResult: "{err: expression is not valid}\n",
		},
	}

	for _, tc := range testCasesUnsuccessful {
		req, _ := http.NewRequest("POST", "/", strings.NewReader(tc.expression))
		w := httptest.NewRecorder()

		CalcHandler(w, req)
		assert.Equal(t, w.Code, http.StatusUnprocessableEntity, "they should be equal")
		assert.Equal(t, w.Body.String(), tc.expectedResult, "they should be equal")
	}
}
