package twilio

import (
	"github.com/spf13/cobra"
)

var TwilioCmd = &cobra.Command{
	Use:   "twilio",
	Short: "Twilio",
	Long:  ``,
}

func init() {
	TwilioCmd.AddCommand(sendVerificationCmd)
	TwilioCmd.AddCommand(checkVerificationCmd)
	TwilioCmd.AddCommand(listServicesCmd)
	TwilioCmd.AddCommand(createServiceCmd)
}
