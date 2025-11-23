package sinch

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/sinch"
	"github.com/spf13/cobra"
)

func newStartVerificationCmd() *cobra.Command {
	var params sinch.StartVerificationParams

	cmd := &cobra.Command{
		Use:   "start-verification",
		Short: "Start Verification",
		Long:  `See API documentation: https://developers.sinch.com/docs/verification/api-reference/verification/tag/Verifications-start/#tag/Verifications-start/operation/StartVerification`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := sinch.NewClient(nil, os.Getenv("SINCH_APP_KEY"), os.Getenv("SINCH_APP_SECRET"))

			resp, err := client.StartVerification(params)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVarP(&params.Identity.Type, "identity-type", "t", "number", "")
	cmd.Flags().StringVarP(&params.Identity.Endpoint, "identity-endpoint", "e", "", "E.164 formatted phone number")
	cmd.Flags().StringVarP(&params.Method, "method", "m", "sms", "The type of the verification request")
	cmd.MarkFlagRequired("identity-endpoint")

	return cmd
}
