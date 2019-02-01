package watcher

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wolfedale/k8s-watcher/cmd/watcher/get"
	"github.com/wolfedale/k8s-watcher/cmd/watcher/version"
)

const defaultLevel = logrus.WarnLevel

// Flags for the watcher command
type Flags struct {
	LogLevel string
}

// Main wraps Run, adding a log.Fatal(err) on error, and setting the log formatter
func Main() {
	// let's explicitly set stdout
	log.SetOutput(os.Stdout)
	// this formatter is the default, but the timestamps output aren't
	// particularly useful, they're relative to the command start
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	if err := Run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}
}

// Run runs the `watcher` root command
func Run() error {
	return NewCommand().Execute()
}

// NewCommand returns a new cobra.Command implementing the root command for watcher
func NewCommand() *cobra.Command {
	flags := &Flags{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "watcher",
		Short: "watcher is a tool for checking events from the Kubernetes clusters",
		Long:  "watcher checks Kubernetes clusters events",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Version: version.Version,
	}
	cmd.PersistentFlags().StringVar(
		&flags.LogLevel,
		"loglevel",
		defaultLevel.String(),
		"logrus log level ",
	)

	cmd.AddCommand(get.NewCommand())
	cmd.AddCommand(version.NewCommand())
	return cmd
}

func runE(flags *Flags, cmd *cobra.Command, args []string) error {
	level := defaultLevel
	parsed, err := log.ParseLevel(flags.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level '%s', defaulting to '%s'", flags.LogLevel, level)
	} else {
		level = parsed
	}
	log.SetLevel(level)
	return nil
}
