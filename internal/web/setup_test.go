package web

import (
	"os"
	"testing"

	"github.com/ArnaudLasnier/pingpong/internal/database/models/factory"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	testWebServer *webServer
	testFactory   *factory.Factory
)

func TestMain(m *testing.M) {
	testWebServer = NewTestWebServer()
	testFactory = factory.New()
	exitCode := m.Run()
	os.Exit(exitCode)

}
