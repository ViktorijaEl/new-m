provider "google" {
  project = "dev-infra-380007"
  region  = "eu"
}

resource "google_storage_bucket" "tf-bucket" {
  project       = "dev-infra-380007"
  name          = "buketdfndjkfj"
  location      = "eu"
  force_destroy = true
  public_access_prevention = "enforced"
  uniform_bucket_level_access = true
  storage_class = "Standard" 

  versioning {
    enabled = true
  }

  logging {
    log_bucket = "dev-github-terraform"
  }
}
