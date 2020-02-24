# shared-2fa

Sometimes you have an account for a website which you share with your team so
everybody can use it, but still want to enable 2FA for enhances security.
This tool helps you to share a virtual TOTP based MFA-device with your team
by saving the intial seed at AWS SecretsManager.

## Pricing
SecretsManager is charged both per secret and per API call. Every secret
costs 0.40\$/Month what will be the main operational costs of this tool. API
calls cost you 0.05$/10.000 requests what you will probably never reach in a
month even with a large team using this tool multiple times a day.

### Ways to recude the costs
Paying 0.40\$/month/secret can quickly sum up if you have a lot of accounts.
Because every value in AWS SecretsManager is a JSON object you could save all
seeds in a single key/value pair instead of one pair per account and cap the
monthly costs 0.40\$/month by this, the downside is, that you would loose the
way to limit access by IAM to single secrets. Therefore we didn't implemented
it this way.

## IAM permissions
With this example policy one can use all features of this tool. If you want
people to have only the ability to generate tokens, you can remove the
`Delete` and `Create` actions.
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:GetSecretValue",
                "secretsmanager:DeleteSecret"
            ],
            "Resource": "arn:aws:secretsmanager:eu-central-1:<your-account-id>:secret:2fa-*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "secretsmanager:CreateSecret",
                "secretsmanager:ListSecrets"
            ],
            "Resource": "*"
        }
    ]
}
```

## Usage
In order to use this tool, you need to have a set of AWS API Keys in the
[default configuration
file](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-where).
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
quite long, random, alphanueric string). If not, you have to use a barcode
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
