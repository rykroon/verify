package twilio

import (
	"encoding/json"
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

var csp *twilio.CreateServiceParams

func runCreateServiceCmd(cmd *cobra.Command, args []string) error {
	client := twilio.NewClient(nil, os.Getenv("TWILIO_API_KEY_SID"), os.Getenv("TWILIO_API_KEY_SECRET"))

	resp, err := client.CreateService(csp)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	var m map[string]any
	if err := json.Unmarshal(resp.Body, &m); err != nil {
		return err
	}

	rawYaml, err := yaml.Marshal(m)
	if err != nil {
		return nil
	}

	fmt.Println(string(rawYaml))
	return nil
}

func init() {
	createServiceCmd.Flags().StringVarP(&csp.FriendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	createServiceCmd.MarkFlagRequired("friendly-name")
}
