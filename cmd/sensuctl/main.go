package main

import (
	"os"

	"github.com/sensu/sensu-go/cli"
	"github.com/sensu/sensu-go/cli/commands"
	hooks "github.com/sensu/sensu-go/cli/commands/hooks"
	"github.com/sensu/sensu-go/cli/commands/root"
	"github.com/sensu/sensu-go/command"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := root.Command()
	sensuCli := cli.New(rootCmd.PersistentFlags())

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		return hooks.ConfigurationPresent(cmd, sensuCli)
	}

	commands.AddCommands(rootCmd, sensuCli)

	if err := rootCmd.Execute(); err != nil {
		if commandErr, ok := err.(command.CommandErrorer); ok {
			os.Exit(commandErr.ExitStatus())
		}
		os.Exit(1)
	}
}
