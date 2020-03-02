package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete an entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fullName := fmt.Sprintf("%s%s", config.NAME_PREFIX, name)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please type 'YES' if you really want to delete the entry for %q\n", name)
		text, _ := reader.ReadString('\n')
		if strings.TrimSuffix(text, "\n") == "YES" {
			sess, err := session.NewSessionWithOptions(session.Options{})
			svc := ssm.New(sess, aws.NewConfig().WithRegion(config.AWS_REGION))
			input := &ssm.DeleteParameterInput{
				Name: aws.String(fullName),
			}
			_, err = svc.DeleteParameter(input)
			if err != nil {
				fmt.Printf("Couldn't delete entry: %s\n", err.Error())
				return
			}

		} else {
			fmt.Println("Aborting.")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
