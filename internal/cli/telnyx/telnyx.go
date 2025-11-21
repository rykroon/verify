package telnyx

import (
	"github.com/spf13/cobra"
)

var TelnyxCmd = &cobra.Command{
	Use:   "telnyx",
	Short: "Telnyx",
	Long:  ``,
}

var verifyProfileId string

func init() {
	TelnyxCmd.AddCommand(newListProfilesCmd())
	TelnyxCmd.AddCommand(newRetrieveProfileCmd())
	TelnyxCmd.AddCommand(newCreateProfileCmd())
	TelnyxCmd.AddCommand(newUpdateProfileCmd())
	TelnyxCmd.AddCommand(newTriggerSmsCmd())
	TelnyxCmd.AddCommand(newVerifyCodeCmd())
	TelnyxCmd.AddCommand(newListTempaltesCmd())
}
