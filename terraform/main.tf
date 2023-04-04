provider "aws" {
  region = "us-west-2"
}

data "aws_s3_buckets" "all" {}

output "s3_bucket_names" {
  value = data.aws_s3_buckets.all.buckets[*].id
}
