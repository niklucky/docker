package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/niklucky/docker/internal/spawner/docker"
	"github.com/stretchr/testify/assert"
)

func TestServer_handleGetContainers(t *testing.T) {
	d := docker.New(&docker.Config{})
	s := New(&Config{
		SecretKey: "123",
	}, d)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/container/5", nil)

	s.handleGetContainers(rec, req, httprouter.Params{
		{Key: "id", Value: "5"},
	})
	fmt.Println(rec.Body.String())
	assert.Equal(t, rec.Body.String(), "hello, 5!")
}
