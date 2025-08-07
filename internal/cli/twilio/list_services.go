package twilio

import (
	"os"

	"github.com/rykroon/verify/internal/utils"
	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
)

var listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "List Services",
	Long:  ``,
	RunE:  runListServicesCmd,
}

func runListServicesCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.ListServices()
	if err != nil {
		return err
	}

	utils.PrintResponse(resp)
	return nil
}

func init() {

}
