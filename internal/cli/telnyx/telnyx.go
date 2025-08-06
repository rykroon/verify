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
	TelnyxCmd.AddCommand(listProfilesCmd)
	TelnyxCmd.AddCommand(retrieveProfileCmd)
	TelnyxCmd.AddCommand(createProfileCmd)
	TelnyxCmd.AddCommand(updateProfileCmd)
	TelnyxCmd.AddCommand(triggerSmsCmd)
	TelnyxCmd.AddCommand(verifyCodeCmd)
	TelnyxCmd.AddCommand(listTemplatesCmd)
}
