package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rykroon/verify/internal/utils"
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

	req, err := client.NewCreateServiceRequest(params)
	if err != nil {
		return err
	}

	resp, err := utils.DoAndReadAll(http.DefaultClient, req)
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
	createServiceCmd.Flags().StringVarP(&csp.friendlyName, "friendly-name", "n", "", "A descriptive name for the service")
	createServiceCmd.MarkFlagRequired("friendly-name")
}
