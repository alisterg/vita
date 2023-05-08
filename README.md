vīta (_latin - life_) - record anything about your life.

https://user-images.githubusercontent.com/8358022/236662869-1be3edb6-457e-417a-8ec7-1a279b4559ad.mp4

# CLI Usage

```sh
vita add movie # runs prompts to insert new movie (as defined as an EntryType)

vita find movie # lists all movies
vita find movie --num 10 # lists latest 10 movies
vita find movie --search "something" # finds movies with "something" in any field

vita update movie --search "something" # updates all movies with "something" in any field

vita routine weekly # runs the 'weekly' routine (as defined as a Routine)
```

# Development

## Prerequisites

- Go >= 1.20.3 `brew install go`
- Sam CLI `brew install aws-sam-cli`
- AWS CLI credentials setup

## Dev commands

- `make build` to build all targets
- `make run` to run the cli app without building
- `make runapi` to run the api locally without building
- `make test` to run automated tests

## Architecture

Using a ports/adapters pattern due to the nature of the app.

This allows us to easily:

- Change / add the data sources
- Test Core because it is completely decoupled
- Add new transport mechanisms (CLI / HTTP etc)

`ports` (repositories, usecases) are interfaces used by the core domain logic

`adapters` (datasources, transport) are plugins that allow us to change the functionality of the app at whim

## Types

`EntryType` is the basic schema for recording entries. Example:

```json
{ "book": ["Title", "Start date (YYYY-MM-DD)", "End date (YYYY-MM-DD)", "Rating", "Review"] }
```

`Entry` is a raw data entry. In the 'book' example, Entry.Data will be stored as a map of the EntryType array values to their result from user input. Example:

```json
{ "uuid": "blah", "date": 123, "data": {"Title": "slaughterhouse 5"...} }
```

`Routine` is a way to run sets of Entry Types without having to manually run each one.

## infra.yml

This is an AWS SAM template. It creates a CloudFormation stack containing all the infrastructure
necessary.

Ensure you have an AWS account and credentials to access it via the CLI.

1. `sam package -t infra.yml`
2. `sam deploy -t infra.yml --stack-name MyStackName --capabilities CAPABILITY_IAM`

`sam deploy -t infra.yml --guided` if the above does not work

## App details

Using Cobra for managing CLI commands and flags:

https://github.com/spf13/cobra/blob/main/user_guide.md
