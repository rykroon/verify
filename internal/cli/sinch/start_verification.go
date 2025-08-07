package sinch

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
)

var startVerificationCmd = &cobra.Command{
	Use:   "start-verification",
	Short: "Start Verification",
	Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-start/#tag/Verifications-start/operation/StartVerification`,
	RunE:  runStartVerification,
}

var svp sinch.StartVerificationParams

func runStartVerification(cmd *cobra.Command, args []string) error {
	client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

	resp, err := client.StartVerification(svp)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	startVerificationCmd.Flags().StringVarP(&svp.Identity.Type, "identity-type", "t", "number", "")
	startVerificationCmd.Flags().StringVarP(&svp.Identity.Endpoint, "identity-endpoint", "e", "", "E.164 formatted phone number")
	startVerificationCmd.Flags().StringVarP(&svp.Method, "method", "m", "sms", "The type of the verification request")
	startVerificationCmd.MarkFlagRequired("identity-endpoint")
}
