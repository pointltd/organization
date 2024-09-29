# Объявление переменных для конфиденциальных параметров

locals {
  zone             = "ru-central1-a"
  username         = "coded"
  ssh_key_path     = ".ssh/ya_coded_ubuntu.pub"
  target_folder_id = "b1g0k22us62vt6kut949"
  registry_name    = "point-registry"
  sa_name          = "sa-registry"
  network_name     = "docker-vm-network"
  subnet_name      = "docker-vm-network-subnet-a"
  vm_name          = "docker-vm"
  image_id         = "fd87tirk5i8vitv9uuo1"
  registry_id      = "crpbvv0uke53s3ief4mk"
}

# Настройка провайдера

terraform {
  required_providers {
    yandex    = {
      source  = "yandex-cloud/yandex"
      version = ">= 0.47.0"
    }
  }
}

provider "yandex" {
  zone = local.zone
  folder_id = local.target_folder_id
}

resource "yandex_compute_instance" "vm-1" {
  name = "from-terraform-vm"
  platform_id = "standard-v1"
  zone = "ru-central1-a"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = local.image_id
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.subnet-1.id
    nat       = true
  }

  metadata = {
    ssh-keys = "ubuntu:${file(".ssh/ya_coded_ubuntu.pub")}"
  }
}

resource "yandex_vpc_network" "network-1" {
  name = "from-terraform-network"
}

resource "yandex_vpc_subnet" "subnet-1" {
  name           = "from-terraform-subnet"
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.network-1.id
  v4_cidr_blocks = ["10.2.0.0/16"]
}

output "internal_ip_address_vm_1" {
  value = yandex_compute_instance.vm-1.network_interface.0.ip_address
}

output "external_ip_address_vm_1" {
  value = yandex_compute_instance.vm-1.network_interface.0.nat_ip_address
}