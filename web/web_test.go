package web_test

// https://elithrar.github.io/article/testing-http-handlers-go/
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"toba.tech/app/lib/auth"
	"toba.tech/app/lib/config"
	"toba.tech/app/lib/web"
	"toba.tech/app/lib/web/encoding"
	"toba.tech/app/lib/web/header/accept"
	"toba.tech/app/lib/web/header/content"
	"toba.tech/app/lib/web/mime"
)

var (
	c = config.HTTP{
		FromZip:    "",
		FromFolder: "static",
	}
	modulePaths = []string{"module1", "module2"}
	authPaths   = map[string]*auth.AuthProvider{
		"auth/dropbox": auth.Providers[auth.Dropbox],
	}
	handler = web.Handle(c, modulePaths, authPaths)
)

func get(t *testing.T, path string) *http.Response {
	assert.NotNil(t, handler)

	r := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.Header.Add(accept.Encoding, encoding.GZip)

	handler(w, r)

	return w.Result()
}

func TestMissingFile(t *testing.T) {
	res := get(t, "/no-such-file")
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestModulePath(t *testing.T) {
	res := get(t, "/module1")
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, mime.HTML, res.Header.Get(content.Type))
}

func TestLogo(t *testing.T) {
	res := get(t, "/img/logo.svg")
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, mime.SVG, res.Header.Get(content.Type))
}

func TestJavascript(t *testing.T) {
	res := get(t, "/js/common.js")
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, mime.JavaScript, res.Header.Get(content.Type))
	assert.Equal(t, encoding.GZip, res.Header.Get(content.Encoding))
}