package sinch

import (
	"github.com/spf13/cobra"
)

var SinchCmd = &cobra.Command{
	Use:   "sinch",
	Short: "Sinch",
	Long:  ``,
}

var verificationId string

func init() {
	SinchCmd.AddCommand(startVerificationCmd)
	SinchCmd.AddCommand(reportVerificationCmd)
	SinchCmd.AddCommand(getVerificationCmd)

}
