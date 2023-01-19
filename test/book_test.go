package test

import (
	"add/models"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var s int64

func TestBook(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id:=createBook(t)
			// updateBook(t,id)
			deleteBook(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createBook(t *testing.T) string {
	response := &models.CategoryBook{}

	request := &models.CreateBook{
		Name:       faker.Name(),
		Price: float64(faker.RandomUnixTime()),
		Description: faker.Paragraph(),
	}

	resp, err := PerformRequest(http.MethodPost, "/book", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateBook(t *testing.T, id string) string {
	response := &models.UpdateBook{}
	request := &models.UpdateBook{
		Name:       faker.Name(),
		Price: float64(faker.RandomUnixTime()),
		Description: faker.Paragraph(),
	}

	resp, err := PerformRequest(http.MethodPut, "/book/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteBook(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/book/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 202)
	}

	return ""
}
