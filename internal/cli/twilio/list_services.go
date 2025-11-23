package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

func newListServicesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-services",
		Short: "List Services",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

			content, err := client.ListServices()
			if err != nil {
				return err
			}

			utils.PrintContent(content)
			return nil
		},
	}

	return cmd
}
