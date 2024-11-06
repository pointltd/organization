locals {
  zone             = "ru-central1-a"
  username         = "coded"
  ssh_key_path     = ".ssh/ya_coded_ubuntu.pub"
  target_folder_id = "b1g0k22us62vt6kut949"
  registry_name    = "point-registry"
  network_name     = "docker-vm-network"
  subnet_name      = "docker-vm-network-subnet-a"
  container_name   = "organization-container"
  image_id         = "crppi5deo87qjhsgaf0c"
  registry_id      = "cr.yandex/crp4640u3tckkugq0upa"
}

terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = ">= 0.47.0"
    }
  }
}

provider "yandex" {
  zone      = local.zone
  folder_id = local.target_folder_id
}

resource "yandex_iam_service_account" "organization-service-account" {
  name = "organization-service-account"
}

resource "yandex_resourcemanager_folder_iam_member" "registry_pull_permission" {
  folder_id = local.target_folder_id
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.organization-service-account.id}"
}

resource "yandex_serverless_container" "organization-app-container" {
  name               = local.container_name
  service_account_id = yandex_iam_service_account.organization-service-account.id
  memory             = 512  # Specify memory in MB
  cores              = 1

  image {
    url = "cr.yandex/crp4640u3tckkugq0upa/organization-app:organization-app"
  }
}

resource "yandex_vpc_network" "network-1" {
  name = "container-network"
}

resource "yandex_vpc_subnet" "subnet-1" {
  name           = "container-subnet"
  zone           = local.zone
  network_id     = yandex_vpc_network.network-1.id
  v4_cidr_blocks = ["10.2.0.0/16"]
}