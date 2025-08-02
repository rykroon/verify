package sinch

import (
	"github.com/spf13/cobra"
)

var SinchCmd = &cobra.Command{
	Use:   "sinch",
	Short: "Sinch",
	Long:  ``,
}

func init() {
	SinchCmd.AddCommand(startVerificationCmd)
	SinchCmd.AddCommand(reportVerificationCmd)

}
