package cmd

import (
	"net"

	"github.com/spf13/cobra"
)

// NewFoxCommand New Fox Command
func NewFoxCommand(c net.Conn) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fox",
		Short: "Fox Command",
		Run: func(cmd *cobra.Command, args []string) {
			c.Write([]byte("Hello Fox!\n"))
		},
	}

	return cmd
}
