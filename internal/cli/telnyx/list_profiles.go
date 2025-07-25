package telnyx

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rykroon/verify/internal/telnyx"
	"github.com/spf13/cobra"
)

var listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "List verification profiles.",
	Long:  ``,
	RunE:  runListProfiles,
}

func runListProfiles(cmd *cobra.Command, args []string) error {
	client := telnyx.NewClient(os.Getenv("TELNYX_API_KEY"))
	result, err := client.ListVerifyProfiles()
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
