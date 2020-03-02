package cmd

import (
	"fmt"
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
		svc := ssm.New(sess, aws.NewConfig().WithRegion(config.AWS_REGION))
		input := &ssm.DescribeParametersInput{}
		data, err := svc.DescribeParameters(input)
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		fmt.Println("Found the following 2FA entries:")
		for _, entry := range data.Parameters {
			if strings.HasPrefix(*entry.Name, config.NAME_PREFIX) {
				cleanName := strings.Replace(*entry.Name, config.NAME_PREFIX, "", 1)
				fmt.Printf("- %s\n", cleanName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
