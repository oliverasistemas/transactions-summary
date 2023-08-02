provider "aws" {
  region = "[your_region]"
}

resource "aws_iam_role" "lambda_role" {
  name = "stori-transactions-summary-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Effect    = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_lambda_function" "stori-transactions-summary" {
  filename      = "lambda_function.zip"
  function_name = "stori"
  role          = aws_iam_role.lambda_role.arn
  handler       = "main"
  runtime       = "go1.x"
  memory_size   = 512
  timeout       = 15

  environment {
    variables = {
      EMAIL_FROM     = "[your_account]" #Replace with your Gmail account
      EMAIL_PASSWORD=  "[application_password]" #Replace with your account application_password
      S3_BUCKET_NAME = "[bucket]"
      S3_FILE_KEY    = "[file_name].csv"
      MYSQL_HOST     = aws_db_instance.stori_rds.endpoint
      MYSQL_USER     = "root"
      MYSQL_PASSWORD = "stori123"
      MYSQL_DB       = "stori"
      MYSQL_PORT     = "3306"
    }
  }
}

resource "aws_db_subnet_group" "rds_subnet_group" {
  name       = "stori-rds-subnet-group"
  subnet_ids = ["subnet-xxxxxx", "subnet-yyyyyy"] # Replace with your actual subnet IDs
}

resource "aws_db_instance" "stori_rds" {
  allocated_storage    = 20
  engine               = "mysql"
  engine_version       = "8.0.23"
  instance_class       = "db.t2.micro"
  username             = "root"
  password             = "stori123"
  db_subnet_group_name = aws_db_subnet_group.rds_subnet_group.name

  tags = {
    Name = "Stori RDS"
  }
}
