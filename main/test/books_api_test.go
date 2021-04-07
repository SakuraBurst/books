package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/SakuraBurst/books.git/main/controllers"
	"github.com/SakuraBurst/books.git/main/driver"
	"github.com/SakuraBurst/books.git/main/models"
	"github.com/SakuraBurst/books.git/main/repository"
	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////////////
// для работы тестов нужно положить .evn в папку test
//////////////////////////////////////////////////////////////

type booksRequest struct {
	id         string
	shouldPass bool
}

type postBookRequest struct {
	request    *strings.Reader
	shouldPass bool
}

type putBookRequest struct {
	request    *strings.Reader
	id         string
	shouldPass bool
}

var BookRepository repository.BookRepository
var BooksController controllers.Controler
var FirstBook string = `{"Title":"Гарри повар и филосовское яйцо","ID":1,"Author":"Kora0108","Year":"2012-01-01T00:00:00Z"}`
var RightTestBookMock string = `{"Title":"Гарри повар и автотесты","Author":"SakuraBurst","Year":"03-04-2021"}`
var WrongTestBookMock string = `{"Title":"Гарри повар и автотесты","Author":"SakuraBurst"}`

func TestMain(m *testing.M) {
	testSetup()
	os.Exit(m.Run())
}

func testSetup() {
	database := driver.ConnectDatabase("DB_URL")
	BookRepository = repository.BookRepository{Database: database}
	BooksController = controllers.Controler{Repository: BookRepository}
}

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksController.GetBooks)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestGetBook(t *testing.T) {
	// обращение к тестовой базе, где айди 1 точно будет, а 1337 точно не будет
	testTable := []booksRequest{
		{"1", true},
		{"1337", false},
	}

	for _, testCase := range testTable {
		path := fmt.Sprintf("/books/%s", testCase.id)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/books/{id}", BooksController.GetBook).Methods("GET")
		router.ServeHTTP(rr, req)
		if rr.Code == http.StatusOK && !testCase.shouldPass {
			t.Errorf("handler should have failed on id %s: got %v want %v",
				testCase.id, rr.Code, http.StatusNotFound)
		}
		if rr.Code != http.StatusOK && testCase.shouldPass {
			t.Errorf("handler should have failed on id %s: got %v want %v",
				testCase.id, rr.Code, http.StatusOK)
		}
		if rr.Code == http.StatusOK && testCase.shouldPass {
			expectedMock := regexp.MustCompile(`{"Title":"Гарри повар и филосовское яйцо","ID":1,"Author":"Kora0108","Year":"2012-01-01T00:00:00Z"}`)
			if !expectedMock.MatchString(rr.Body.String()) {
				t.Errorf("handler should have failed body on id %s: got %v want %v",
					testCase.id, rr.Body, expectedMock)
			}
		}
	}
}

func TestAddBook(t *testing.T) {
	testTable := []postBookRequest{
		{request: strings.NewReader(RightTestBookMock), shouldPass: true},
		{request: strings.NewReader(WrongTestBookMock), shouldPass: false},
	}
	books := make([]models.Book, 0)
	for _, testCase := range testTable {
		req, err := http.NewRequest("POST", "/books", testCase.request)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(BooksController.AddBook)

		handler.ServeHTTP(rr, req)

		if rr.Code == http.StatusOK && !testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusNotFound)
		}
		if rr.Code != http.StatusOK && testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}
		if rr.Code == http.StatusOK && testCase.shouldPass {
			decoder := json.NewDecoder(rr.Body)
			decoder.Decode(&books)
		}
	}

	// не уверен на счет этого, должен ли респонс отдаваться в порядке записи в бд? но удалять тестовую запись надо, а лучших идей у меня нет
	testedBookId := strconv.Itoa(books[len(books)-1].ID)
	BookRepository.DeleteBookFromDatabase(testedBookId)
}

func TestUpdateBook(t *testing.T) {
	testTable := []putBookRequest{
		{request: strings.NewReader(FirstBook), id: "1", shouldPass: true},
		{request: strings.NewReader(WrongTestBookMock), id: "1", shouldPass: false},
		{request: strings.NewReader(RightTestBookMock), id: "1337", shouldPass: false},
	}
	for _, testCase := range testTable {
		path := fmt.Sprintf("/books/%s", testCase.id)
		req, err := http.NewRequest("PUT", path, testCase.request)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/books/{id}", BooksController.UpdateBook).Methods("PUT")

		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusOK && !testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusNotFound)
		}
		if rr.Code != http.StatusOK && testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}
		if rr.Code == http.StatusOK && testCase.shouldPass {
			expectedMock := regexp.MustCompile(`{"Title":"Гарри повар и филосовское яйцо","ID":1,"Author":"Kora0108","Year":"2012-01-01T00:00:00Z"}`)
			if !expectedMock.MatchString(rr.Body.String()) {
				t.Errorf("handler should have failed body on id %s: got %v want %v",
					testCase.id, rr.Body, expectedMock)
			}
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testRowId, err := addTestBook(t)
	if err != nil {
		t.Fatal(err)
	}
	testTable := []booksRequest{
		{id: testRowId, shouldPass: true},
		{id: "1337", shouldPass: false},
	}
	for _, testCase := range testTable {
		path := fmt.Sprintf("/books/%s", testCase.id)
		req, err := http.NewRequest("DELETE", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/books/{id}", BooksController.DeleteBook).Methods("DELETE")

		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusOK && !testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusNotFound)
		}
		if rr.Code != http.StatusOK && testCase.shouldPass {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}
	}
}

func addTestBook(t *testing.T) (string, error) {
	reader := strings.NewReader(RightTestBookMock)
	books := make([]models.Book, 0)
	req, err := http.NewRequest("POST", "/books", reader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksController.AddBook)

	handler.ServeHTTP(rr, req)
	if rr.Code == http.StatusOK {
		decoder := json.NewDecoder(rr.Body)
		decoder.Decode(&books)
		return strconv.Itoa(books[len(books)-1].ID), nil
	}
	return "", errors.New("problem with add new row")
}
