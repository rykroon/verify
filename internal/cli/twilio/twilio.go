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
	TwilioCmd.AddCommand(newSendVerificationCmd())
	TwilioCmd.AddCommand(newCheckVerificationCmd())
	TwilioCmd.AddCommand(newListServicesCmd())
	TwilioCmd.AddCommand(newCreateServiceCmd())
}
