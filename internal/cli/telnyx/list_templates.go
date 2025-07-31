package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/pkg/telnyx"
	"github.com/spf13/cobra"
)

var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "List Verification Profile Templates",
	Long:  ``,
	RunE:  runListTemplates,
}

func runListTemplates(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	result, err := client.ListMessageTemplates()
	if err != nil {
		return err
	}
	bytes_, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(bytes_))
	return nil
}

func init() {

}
