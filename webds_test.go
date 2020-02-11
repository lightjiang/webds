package webds

import (
	"github.com/lightjiang/utils/log"
	"net/http"
	"testing"
)

type testHandler struct {
	upgrade *Server
}

func (t *testHandler) init() {
	t.upgrade = New(Config{
		BinaryMessages: false,
	})
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("receive " + r.Header.Get("id"))
	conn, err := t.upgrade.Upgrade(w, r)
	if err != nil {
		log.HandlerErrs(err)
		return
	}
	conn.OnInner("/inner/abc", func() {})
	conn.OnInner("/inner/123", func() {})
	conn.Wait()
}

func TestDefaultIDGenerator(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(DefaultIDGenerator(nil))
	}
}

func BenchmarkDefaultIDGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultIDGenerator(nil)
	}
}

func TestNew(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	log.Info().Msg("start webds server")
	h := testHandler{}
	h.init()
	http.ListenAndServe(":8080", &h)
}
