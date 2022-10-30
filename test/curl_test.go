package test

import (
	"os"
	"testing"

	"github.com/barancanatbas/curl-go/curl"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestHelp(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			curl.Command(),
		},
	}

	testArgs := []string{execName, "curl", "-get", "https://google.com"}
	require.NoError(t, app.Run(testArgs))
}
