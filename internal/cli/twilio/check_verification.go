package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

func newCheckVerificationCmd() *cobra.Command {
	var params twilio.CheckVerificationParams

	cmd := &cobra.Command{
		Use:   "check-verification",
		Short: "Check Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

			content, err := client.CheckVerification(params.ServiceSid, params.CheckVerificationForm)
			if err != nil {
				return err
			}

			utils.PrintContent(content)

			return nil
		},
	}

	cmd.Flags().StringVarP(&params.ServiceSid, "service-sid", "s", "", "The SID of the verification Service.")
	cmd.Flags().StringVarP(&params.To, "to", "t", "", "The phone number or email to verify.")
	cmd.Flags().StringVarP(&params.VerificationSid, "verification-sid", "V", "", "A SID that uniquely identifies the Verification Check.")
	cmd.Flags().StringVarP(&params.Code, "code", "c", "", "The 4-10 character string being verified.")
	cmd.MarkFlagRequired("service-sid")
	cmd.MarkFlagsOneRequired("to", "verification-sid")
	cmd.MarkFlagRequired("code")

	return cmd
}
