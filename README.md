vīta (_latin - life_)

Record everything about your life. Currently using DynamoDB for data storage / retrieval.

# USAGE

todo

# DEVELOPMENT

## Prerequisites

- Go >= 1.20.3 `brew install go`
- Sam CLI `brew install aws-sam-cli`
- AWS CLI credentials setup

## entry_types.json

The file `entry_types.json` contains your personal data types that the CLI application
will recognise.

Example:

```json
{ "location": ["City name"] }
```

This will allow you to enter a `location` entry into the application; it will prompt you for
the 'City name'.

## routines.json

The file `routines.json` contains routines that you can define to run sets of 'entries'.

Example:

```json
{ "weekly": ["location"] }
```

This looks for any key in `entry_types.json` that matches values in the array; in this example,
it will add a 'location' entry.

## Infrastructure

Ensure you have an AWS account and credentials to access it via the CLI.

1. `sam package -t infra.yml`
2. `sam deploy -t infra.yml --stack-name MyStackName --capabilities CAPABILITY_IAM`

This will create the CloudFormation stack containing a DynamoDB table and IAM role to access it.

## App details

Using Cobra for managing CLI commands and flags:

https://github.com/spf13/cobra/blob/main/user_guide.md
