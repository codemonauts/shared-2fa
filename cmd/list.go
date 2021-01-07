package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available entries",
	Run: func(cmd *cobra.Command, args []string) {
		sess, err := session.NewSessionWithOptions(session.Options{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		svc := ssm.New(sess, aws.NewConfig().WithRegion(config.AWS_REGION))
		input := &ssm.DescribeParametersInput{}

		var entries []string

		svc.DescribeParametersPages(input,
			func(page *ssm.DescribeParametersOutput, lastPage bool) bool {
				for _, entry := range page.Parameters {
					// Only parameters with our defined prefix are part of this tool
					if strings.HasPrefix(*entry.Name, config.NAME_PREFIX) {
						cleanName := strings.Replace(*entry.Name, config.NAME_PREFIX, "", 1)
						entries = append(entries, cleanName)
					}
				}
				return !lastPage
			})

		if len(entries) > 0 {
			// Sort entries alphabetically
			sort.Strings(entries)

			fmt.Println("Found the following entries:")
			for _, entry := range entries {
				fmt.Printf("- %s\n", entry)
			}
		} else {
			fmt.Println("Found no entries.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
