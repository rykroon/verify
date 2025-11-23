package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

func newCreateServiceCmd() *cobra.Command {
	var params twilio.CreateServiceParams

	var cmd = &cobra.Command{
		Use:   "create-service",
		Short: "Create Service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

			resp, err := client.CreateService(params)
			if err != nil {
				return err
			}

			utils.PrintResponse(resp)
			return nil
		},
	}

	cmd.Flags().StringVarP(&params.FriendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	cmd.Flags().IntVar(&params.CodeLength, "code-length", 0, "The length of the code")
	cmd.MarkFlagRequired("friendly-name")

	return cmd
}
