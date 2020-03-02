package cmd

import (
	"fmt"
	"time"

	"github.com/codemonauts/shared-2fa/config"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xlzd/gotp"
)

type Entry struct {
	Seed string `json:"seed"`
}

var generateCmd = &cobra.Command{
	Use:   "generate <name>",
	Short: "Generate a token for the given entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fullName := fmt.Sprintf("%s%s", config.NAME_PREFIX, name)

		sess, err := session.NewSessionWithOptions(session.Options{})
		svc := ssm.New(sess, aws.NewConfig().WithRegion(config.AWS_REGION))

		input := &ssm.GetParameterInput{
			Name:           aws.String(fullName),
			WithDecryption: aws.Bool(true),
		}

		data, err := svc.GetParameter(input)

		if err != nil {
			fmt.Printf("Couldn't get entry from parameter store: %s\n", err.Error())
			return
		}

		totp := gotp.NewDefaultTOTP(*data.Parameter.Value)
		key, expiration := totp.NowWithExpiration()
		exp := expiration - time.Now().Unix()

		clipboardFlag, _ := cmd.Flags().GetBool("clipboard")
		if clipboardFlag {
			if clipboard.Unsupported {
				color.Red("Clipboard functionality is not supported on your system")
			} else {
				clipboard.WriteAll(key)
			}
		}
		fmt.Printf("%s (%s)\n", key, colorExpiration(exp))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolP("clipboard", "c", false, "Write token to clipboard")
}

func colorExpiration(duration int64) string {
	s := fmt.Sprintf("Expires in %ds", duration)
	if duration < 10 {
		return color.RedString(s)
	} else {
		return s
	}
}
