provider "aws" {
  region = var.region
  access_key = var.access_key
  secret_key = var.secret_key
  assume_role {
    role_arn = var.role_arn
  }
}

resource "aws_sns_topic" "ghec_topic" {
  name = "ghec_topic"
  kms_master_key_id = var.kms_master_key_id
}

resource "aws_sns_topic_subscription" "ghec_subscription" {
  topic_arn    = aws_sns_topic.ghec_topic.arn
  protocol     = "email"
  endpoint     = "viktorija.springe@merck.com"
}


# DocumentDB
# Create a new VPC for the DocumentDB cluster
# resource "aws_vpc" "example" {
#   cidr_block = "10.0.0.0/16"
# }
data "aws_vpc" "example" {
  id = "vpc-06329d9e801d00819"
}
# Create a subnet in the VPC for the DocumentDB cluster
resource "aws_subnet" "example" {
  cidr_block = "10.0.0.0/24"
  vpc_id     = data.aws_vpc.example.id
}
# Create a new subnet group for the DocumentDB cluster
resource "aws_db_subnet_group" "example" {
  name       = "example-subnet-group"
  subnet_ids = [aws_subnet.example.id]
}
# Create security group for the DocumentDB cluster
resource "aws_security_group_rule" "http_ingress" {
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["10.0.0.0/24"]
  security_group_id = aws_security_group.example.id
}
resource "aws_security_group" "example" {
  name_prefix = "example-"
  vpc_id      = data.aws_vpc.example.id

  ingress {
    from_port   = 0
    to_port     = 27017
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow DocumentDB traffic from all IP addresses"
    }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }
}
# Create the DocumentDB cluster
resource "aws_docdb_cluster" "example" {
  cluster_identifier          = "example-cluster"
  engine                      = "docdb"
  master_username             = "admin"
  master_password             = "password"
  db_subnet_group_name        = aws_db_subnet_group.example.name
  vpc_security_group_ids      = [aws_security_group.example.id]
  skip_final_snapshot         = true
  backup_retention_period     = 7
  preferred_backup_window     = "01:00-02:00"
}
