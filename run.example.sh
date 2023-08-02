#!/bin/bash
go build -o build/summary-mailer .

# S3 bucket
export S3_BUCKET_NAME=[bucket]
export S3_FILE_KEY=[file_name].csv

# OR For Local file system:
export LOCAL_CSV_FILE_PATH=/Users/[your_user]/code/stori/fixtures/sample_input.csv

# Mailer
export EMAIL_FROM=[your_account]@gmail.com #Replace with your Gmail account
export EMAIL_PASSWORD=[application_password] #Replace with your account application_password

# AWS services
export AWS_REGION=[region] # Optional
export AWS_ACCESS_KEY_ID=[KEY_ID]
export AWS_SECRET_ACCESS_KEY=[ACCESS_KEY]

# Database (MySQL)
export MYSQL_HOST=db
export MYSQL_USER=root
export MYSQL_PASSWORD=stori123
export MYSQL_DB=stori
export MYSQL_PORT=3306

./build/summary-mailer -address=[to_address] #Replace with your recipient address