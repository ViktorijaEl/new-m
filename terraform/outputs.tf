##################################
# Create Storage Bucket - Output #
##################################

# output "gcp_bucket" {
#   value = module.bucket.google_storage_bucket.tf-bucket
# }

output "bucket_name" {
  value = module.bucket.bucket_name
}



output "google_compute_instance" {
  value = google_compute_instance.default
  sensitive = true
}
