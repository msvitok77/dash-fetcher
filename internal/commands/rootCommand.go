package commands

import (
	"fmt"

	"github.com/msvitok77/dash-fetcher/internal/commands/dash"
	"github.com/spf13/cobra"
)

const exampleUsage = `
dash-fetcher -p file:///home/user/test/samples/1080.mpd

dash-fetcher -p http://localhost:8080/1.mpd

dash-fetcher -p https://localhost/1.mpd
`

type Root struct {
	mpdURL  string
	command *cobra.Command
}

// NewRoot prepares a new cobra command which encapsulates the whole fetcher logic
func NewRoot() *Root {
	var root Root
	root.command = &cobra.Command{
		Use:                   "dash-fetcher -p [URL]",
		Short:                 "dash-fetcher parses mpd files",
		Long:                  "💡 dash-fetcher parses mpd files from given URL, supports local files and http/s files",
		Example:               exampleUsage,
		DisableFlagsInUseLine: true,
		Run:                   root.process,
	}

	root.command.Flags().StringVarP(&root.mpdURL, "mpdURL", "p", "", "URL to process .mpd files")

	return &root
}

func (r *Root) process(cmd *cobra.Command, args []string) {
	if r.mpdURL == "" {
		fmt.Println("\n💀 empty 'mpdURL'")
		return
	}

	if err := dash.Parse(r.mpdURL); err != nil {
		fmt.Printf("💀 %v", err)
	}
}

// Execute executes the underlying cobra command
func (r *Root) Execute() error {
	return r.command.Execute()
}
