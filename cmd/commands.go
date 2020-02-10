package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var statusConfigPath string

var rootCmd = &cobra.Command{
	Use:   "status",
	Short: "Status is an open source implementation of ping based status page.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatalf("Error displaying help: %s", err.Error())
		}
	},
}
