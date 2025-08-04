package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "List Verification Profile Templates",
	Long:  ``,
	RunE:  runListTemplates,
}

func runListTemplates(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	jsonBytes, err := client.ListMessageTemplates()
	if err != nil {
		return err
	}
	var m map[string]any
	if err := json.Unmarshal(jsonBytes, &m); err != nil {
		return err
	}
	yamlBytes, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlBytes))
	return nil
}

func init() {

}
