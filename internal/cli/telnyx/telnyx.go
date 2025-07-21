package telnyx

import (
	"github.com/spf13/cobra"
)

var TelnyxCmd = &cobra.Command{
	Use:   "telnyx",
	Short: "Telnyx",
	Long:  ``,
}

func init() {
	TelnyxCmd.AddCommand(createProfileCmd)
	TelnyxCmd.AddCommand(triggerSmsCmd)
}
