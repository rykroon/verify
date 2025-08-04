package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/twilio"
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

	rawJson, err := client.ListServices()
	if err != nil {
		return err
	}

	var m map[string]any
	if err := json.Unmarshal(rawJson, &m); err != nil {
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

}
