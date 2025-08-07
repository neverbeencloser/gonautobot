// //go:build unit
// //+build unit
package nautobot_test

import (
	nautobot "github.com/neverbeencloser/gonautobot"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"os"
	"testing"
)

const testURL = "https://demo.nautobot.com/api/"

var testClient *nautobot.Client

// TestMain function is used to initialize any data prior to running tests. Since
// Go testing does not execute tests in any specific order, the TestMain function can handle
// build and tear down of test data for us.
func TestMain(m *testing.M) {
	testClient = nautobot.New(
		nautobot.WithToken("aaaaaaaaaaaa"),
		nautobot.WithHTTPClient(http.DefaultClient),
		nautobot.WithEndpoint(testURL),
	)
	gock.InterceptClient(testClient.Request.Client)

	defer gock.Off()
	defer gock.RestoreClient(testClient.Request.Client)

	exitVal := m.Run()
	os.Exit(exitVal)
}
