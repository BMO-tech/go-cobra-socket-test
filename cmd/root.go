package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func newRootCommand(c net.Conn) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "test",
		Short: "Test root command",
	}

	rootCmd.AddCommand(
		NewFoxCommand(c),
	)

	return rootCmd
}

// Execute run the root command
func Execute(c net.Conn, command string) {
	rootCmd := newRootCommand(c)
	rootCmd.SetArgs([]string{command})
	if err := rootCmd.Execute(); err != nil {
		e := fmt.Sprintf("%s\n", err)
		c.Write([]byte(e))
	}
}
