##################################
# Create Storage Bucket - Output #
##################################

output "gcp_bucket" {
  value = google_storage_bucket.tf-bucket
}


output "google_compute_instance" {
  value = google_compute_instance.default
  sensitive = true
}
