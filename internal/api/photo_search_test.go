package api

import (
	"github.com/tidwall/gjson"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPhotos(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app, router, ctx := NewApiTest()

		GetPhotos(router, ctx)
		r := PerformRequest(app, "GET", "/api/v1/photos?count=10")
		len := gjson.Get(r.Body.String(), "#")
		assert.LessOrEqual(t, int64(3), len.Int())
		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("invalid request", func(t *testing.T) {
		app, router, ctx := NewApiTest()
		GetPhotos(router, ctx)
		result := PerformRequest(app, "GET", "/api/v1/photos?xxx=10")
		assert.Equal(t, http.StatusBadRequest, result.Code)
	})
}
