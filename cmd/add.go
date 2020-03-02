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

var addCmd = &cobra.Command{
	Use:   "add <name> <seed>",
	Short: "Create a new entry",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fullName := fmt.Sprintf("%s%s", config.NAME_PREFIX, name)
		token := strings.ReplaceAll(strings.ToUpper(args[1]), " ", "")

		forceFlag, _ := cmd.Flags().GetBool("force")

		sess, err := session.NewSessionWithOptions(session.Options{})
		svc := ssm.New(sess, aws.NewConfig().WithRegion(config.AWS_REGION))
		input := &ssm.PutParameterInput{
			Name:      aws.String(fullName),
			Value:     aws.String(token),
			Type:      aws.String("SecureString"),
			Overwrite: aws.Bool(forceFlag),
		}
		_, err = svc.PutParameter(input)
		if err != nil {
			fmt.Printf("error while writing to parameter store: %s\n", err.Error())
			return
		}
		fmt.Println("Saved successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("force", "f", false, "Overwrite existing entries")
}
