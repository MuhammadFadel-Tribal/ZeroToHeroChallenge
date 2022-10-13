package api

/*import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetAll)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":1,"Amount":50.47,"Currency":"USD","CreatedAt":"2022-07-19T13:02:01.440618Z"},{"ID":2,"Amount":8000.7865,"Currency":"MXN","CreatedAt":"2022-07-19T13:02:01.440618Z"},{"ID":3,"Amount":60789.674,"Currency":"EGP","CreatedAt":"2022-07-19T13:02:01.440618Z"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandleCreate(t *testing.T) {

	var jsonStr = []byte(`{"ID":4,"Amount":365.78,"Currency":"CLP","CreatedAt":"2022-07-19T13:02:01.440618Z"`)

	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Transaction created successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandleGetOne(t *testing.T) {

	req, err := http.NewRequest("GET", "/transactions/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("ID", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetOne)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ID":1,"Amount":50.47,"Currency":"USD","CreatedAt":"2022-07-19T13:02:01.440618Z"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}*/

/*func TestGetTransactionNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("ID", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTransaction)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}*/
