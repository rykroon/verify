package twilio

import (
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Create Service",
	Long:  ``,
	RunE:  runCreateServiceCmd,
}

type createServiceParams struct {
	friendlyName string
}

var csp createServiceParams

func runCreateServiceCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	params := twilio.NewCreateServiceParams()
	params.SetFriendlyName(csp.friendlyName)
	result, err := client.CreateService(params)
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
	createServiceCmd.Flags().StringVarP(&csp.friendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	createServiceCmd.MarkFlagRequired("friendly-name")
}
