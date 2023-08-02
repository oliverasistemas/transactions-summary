## Transaction summary mailer

### Program Features
1. It can process financial transactions information read from CSV file.
2. The CSV input file can be read from local file system or from AWS S3 bucket.
3. It persists transactions and accounts into a relational database (MySQL).
4. It sends a summary email listing transactions aggregated data. 
5. It can be easily branded.
6. It can be executed on local, container or lambda environments.

## Run locally on macOS
### Install Go:
Make sure you have Go installed on your system. You can download it from the official Go website (https://golang.org/) and follow the installation instructions.

### Clone the repository: Clone this repository to your local machine using git clone.

## Set up environment variables

### CSV file source
#### AWS S3
Set the following environment variables:
```shell
S3_BUCKET_NAME=[bucket]
S3_FILE_KEY=[file_name].csv
AWS_REGION=[region] # Optional
```
or:

#### Local file system
Set the following environment variables:
```shell
LOCAL_CSV_FILE_PATH=/Users/[your_user]/code/stori/fixtures/sample_input.csv
```

## Mailer
#### The mailer is set up with Gmail, but it can be easily changed modifying `mailer.NewMailer` 

Before setting up the environment make sure you count with an application specific password
https://support.google.com/accounts/answer/185833?hl=en
```shell
EMAIL_FROM=[your_account]@gmail.com #Replace with your Gmail account
EMAIL_PASSWORD=[application_password] #Replace with your account application_password
```

## AWS authentication and configuration
### Configure you AWS region
```shell
AWS_REGION=[region] # Optional
AWS_ACCESS_KEY_ID=[KEY_ID]
AWS_SECRET_ACCESS_KEY=[ACCESS_KEY]
```

## Database
 You can run a local container [db service listed on docker-compose]. 

 Or deploy an RDS instance with the Terraform main.tf .

#### Regardless of your choice, please, set the environment variables for MySQL.
```shell
MYSQL_HOST=db
MYSQL_USER=root
MYSQL_PASSWORD=stori123
MYSQL_DB=stori
MYSQL_PORT=3306
```

## Install dependencies:
This project uses AWS SDKs and MySQL driver. Use go get to install the required dependencies.
```shell
go get
```
### Program flags
The program receives and address flag
```shell
summary-mailer -address="you@example.com"
```

### Sample data
There is a sample file in `fixtures/sample_input.csv`

### Build and Run 
You can make a copy of `run.example.sh` and replace the environment variables and program flags

```shell
cp run.example.sh run.sh
```

### Build and Run
You can make a copy of `run.example.sh` and replace the environment variables and program flags

```shell
cp run.example.sh run.sh
```

### Testing
Run the unit tests 
```shell
go test -cover ./...
```

### Run with Docker compose
Replace your environment variables in docker-compose.yml
```shell
docker compose up
```

### Run as AWS Lambda

#### Compile and zip lambda function handler
```shell
GOOS=linux GOARCH=amd64 go build -o ./build/main ./aws_lambda/main.go
zip build/lambda_function.zip build/main
```

#### Create AWS artifacts with Terraform

You can use the Terraform files included to create lambda and an RDS instance to hold data
```shell
terraform init
terraform apply
```

