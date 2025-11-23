package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

func newSendVerificationCmd() *cobra.Command {
	var params twilio.SendVerificationParams

	cmd := &cobra.Command{
		Use:   "send-verification",
		Short: "Send Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

			content, err := client.SendVerification(params.ServiceSid, params.SendVerificationForm)
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	cmd.Flags().StringVar(&params.ServiceSid, "service-sid", "", "The SID of the verification Service.")
	cmd.Flags().StringVar(&params.To, "to", "", "The phone number or email to verify.")
	cmd.Flags().StringVar(&params.Channel, "channel", "", "The verification method to use.")
	cmd.MarkFlagRequired("service-sid")
	cmd.MarkFlagRequired("to")
	cmd.MarkFlagRequired("channel")

	return cmd
}
