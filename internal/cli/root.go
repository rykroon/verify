package cli

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/cli/telnyx"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "verify",
	Short: "The Alpha and Omega of Phone Verification",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(telnyx.TelnyxCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
