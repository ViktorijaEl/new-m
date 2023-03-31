################################
# Create Storage Bucket - Main #
################################

# Define the bucket module
module "bucket" {
  source = "./bucket"

  # Pass variables to the bucket module
  gcp_project   = var.gcp_project
  bucket-name   = var.bucket-name
  gcp_region    = var.gcp_region
  gcp_zone      = var.gcp_zone
  storage-class = var.storage-class 
}

################################
   
    
    
    
    
# # Define the vm module
# module "vm" {
#   source = "./vm"

#   # Pass variables to the vm module
#   instance_name = var.instance_name
#   machine_type = var.machine_type
#   zone = var.zone
#   network_name = var.network_name
#   subnet_name = var.subnet_name
#   disk_size = var.disk_size
#   disk_image = var.disk_image
# }



# Create a disk to attach to the VM
resource "google_compute_disk" "boot" {
  name = "boot-disk"
  type = "pd-standard"
  size = "10"
  zone = "europe-west3-c"

  disk_encryption_key {
    raw_key = var.csek_key
  }

  labels = {
    name = "boot-disk"
  }
}

# Create a Simple VM
resource "google_compute_instance" "default" {
  name         = var.vm-name
  machine_type = var.machine_type
  zone         = var.gcp_zone
  metadata = {
      block-project-ssh-keys = true
  }

  boot_disk {
    initialize_params {
      image = var.image
    }
    auto_delete = true
    disk_encryption_key_raw = var.csek_key
  }

  shielded_instance_config {
    enable_integrity_monitoring  = true
    enable_vtpm                  = true
  }

  service_account {
    email  = "dev-compute@dev-infra-380007.iam.gserviceaccount.com"
    scopes = ["cloud-platform"]
  }

  network_interface {
    network = "default"
  }
}
