package main // import "ir-blaster.com/ir-blaster"

import (
	"os"

	"ir-blaster.com/ir-blaster/internal/glue"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: os.Args[0],
	}

	for _, cmdBuilder := range []func() (*cobra.Command, error){glue.BuildBasicHTMLFileCmd, glue.BuildGithubCmd} {
		cmd, err := cmdBuilder()
		if err != nil {
			log.Fatal(err)
		}

		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
