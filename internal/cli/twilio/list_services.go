package twilio

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "List Services",
	Long:  ``,
	RunE:  runListServicesCmd,
}

func runListServicesCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	result, err := client.ListServices()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(result)
	if err != nil {
		return nil
	}

	fmt.Println(string(data))
	return nil
}

func init() {

}
