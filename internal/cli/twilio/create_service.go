package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

var createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Create Service",
	Long:  ``,
	RunE:  runCreateServiceCmd,
}

var csp twilio.CreateServiceParams

func runCreateServiceCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.CreateService(csp)
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {
	createServiceCmd.Flags().StringVarP(&csp.FriendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	createServiceCmd.MarkFlagRequired("friendly-name")
}
