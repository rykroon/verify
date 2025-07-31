package cli

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/cli/telnyx"
	"github.com/rykroon/verify/internal/cli/twilio"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "verify",
	Short: "The Alpha and Omega of Phone Verification",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(telnyx.TelnyxCmd)
	rootCmd.AddCommand(twilio.TwilioCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
