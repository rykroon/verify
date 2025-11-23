package twilio

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func newSendVerificationCmd() *cobra.Command {
	var serviceSid string
	var form twilio.SendVerificationForm

	cmd := &cobra.Command{
		Use:   "send-verification",
		Short: "Send Verification",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(
				os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"), nil,
			)

			result, err := client.SendVerification(serviceSid, form)
			if err != nil {
				return err
			}

			yamlBytes, err := yaml.Marshal(result)
			if err != nil {
				return fmt.Errorf("failed to encode as yaml: %w", err)
			}
			fmt.Println(string(yamlBytes))
			return nil
		},
	}

	cmd.Flags().StringVar(&serviceSid, "service-sid", "", "The SID of the verification Service.")
	cmd.Flags().StringVar(&form.To, "to", "", "The phone number or email to verify.")
	cmd.Flags().StringVar(&form.Channel, "channel", "", "The verification method to use.")
	cmd.MarkFlagRequired("service-sid")
	cmd.MarkFlagRequired("to")
	cmd.MarkFlagRequired("channel")

	return cmd
}
