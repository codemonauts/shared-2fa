# shared-2fa
[![Join Rocketchat Badge](https://codemonauts.rocket.chat/images/join-chat.svg)](https://codemonauts.rocket.chat)

Sometimes you have an account for a website which you share with your team so
everybody can use it, but still want to enable 2FA for enhanced security.
This tool helps you to share a virtual TOTP based MFA-device with your team
by saving the intial seed at the parameter store of AWS SecureSystemManager.

## Pricing
Using this tool will probably not produce any AWS costs for you.
Storing/Getting Parameters in the ParameterStore is
[free](https://aws.amazon.com/systems-manager/pricing/#Parameter_Store). The
only operations that will get charged by AWS is when you generate a token,
because in the background the tool has to retrieve the parameter from the
ParameterStore and because the parameters are encrypted this will make a call
to KMS to get the encryption key. The first 20.000 requests per month to KMS
are part of the *Free Tier* and after that 10.000 req/month are 0,03$. So
even with a big team and lots of entries, this will propably never apear on
your invoice.

## IAM permissions
With this example policy one can use all features of this tool. If you want
people to have only the ability to generate tokens, you can just remove the
`Delete` and `Create` actions.

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ssm:PutParameter",
                "ssm:DeleteParameter",
                "ssm:GetParameter"
            ],
            "Resource": "arn:aws:ssm:eu-central-1:<your-account-id>:parameter/2fa-*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": "ssm:DescribeParameters",
            "Resource": "*"
        }
    ]
}
```

## Usage
In order to use this tool, you need to have a set of AWS API Keys in the
[default configuration file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-where).
If you have used the aws-cli before, you are already good to go :)

### add
```
Create a new entry

Usage:
  shared-2fa add <name> <seed>
```
When enabling 2FA for an online service, you probably get an QR-Code which
you could scan e.g. with the Google Authenticator app on your smartphone.
Sometimes the website shows your the seed right next to the image (look for a
quite long, random, alphanumeric string). If not, you have to use a barcode
scanner app to get the content of the QR-Code and extract the seed out of
this [special
URI](https://github.com/google/google-authenticator/wiki/Key-Uri-Format).

### delete
```
Delete an entry

Usage:
  shared-2fa delete <name>
```

### generate
```
Generate a token for the given entry

Usage:
  shared-2fa generate <name>
```

### list
```
List all available entries

Usage:
  shared-2fa list
```


With ❤ by [codemonauts](https://codemonauts.com)
