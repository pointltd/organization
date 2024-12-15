locals {
  zone             = "ru-central1-a"
  username         = "coded"
  ssh_key_path     = ".ssh/ya_coded_ubuntu.pub"
  target_folder_id = "b1g0k22us62vt6kut949"
  registry_name    = "point-registry"
  network_name     = "docker-vm-network"
  subnet_name      = "docker-vm-network-subnet-a"
  container_name   = "organization-app"
  image_id         = "crppi5deo87qjhsgaf0c"
  registry_id      = "cr.yandex/crp4640u3tckkugq0upa"
  db_url_secret_id = "e6qdce4u6atkl8njrrol"
  jwt_secret_id    = "e6q0tnu3qnmhr4imib32"
  jwt_secret_version_id    = "e6qqunh4ra3ee11b24a8"
  db_url_secret_version_id = "e6qbk89fofussm10ksu5"
}

terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = ">= 0.90.0"
    }
  }

  backend "s3" {
    endpoints = {
      s3 = "https://storage.yandexcloud.net",
      #       dynamodb = "https://docapi.serverless.yandexcloud.net/ru-central1/b1gqteti3n0acn3o5mge/etnkuhk8r01g859r9fd2"
    }
    bucket = "terraform-state-s3-bucket-test"
    key    = "organization/terraform.tfstate"
    region = "ru-central1"

    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true
    skip_s3_checksum            = true

    #     dynamodb_table = "terraform-state-lock"
  }
}

provider "yandex" {
  zone      = local.zone
  folder_id = local.target_folder_id
}

resource "yandex_iam_service_account" "organization-sa" {
  name = "organization-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "registry_pull_permission" {
  folder_id = local.target_folder_id
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.organization-sa.id}"
}

resource "yandex_lockbox_secret_iam_binding" "lockbox_db_url_viewer_permission" {
  secret_id = local.db_url_secret_id
  role      = "lockbox.payloadViewer"
  members   = ["serviceAccount:${yandex_iam_service_account.organization-sa.id}"]
}

resource "yandex_lockbox_secret_iam_binding" "lockbox_jwt_secret_id_viewer_permission" {
  secret_id = local.jwt_secret_id
  role      = "lockbox.payloadViewer"
  members   = ["serviceAccount:${yandex_iam_service_account.organization-sa.id}"]
}

variable "ORGANIZATION_IMAGE_TAG" {
  type      = string
}

# module "organization-app" {
#   source = "../../modules/application"
#
#   container_name = local.container_name
#   registry_id    = local.registry_id
#   db_url_secret_id = local.db_url_secret_id
#   jwt_secret_id    = local.jwt_secret_id
#   db_url_secret_version_id = local.db_url_secret_version_id
#   jwt_secret_version_id    = local.jwt_secret_version_id
#   ORGANIZATION_IMAGE_TAG = var.ORGANIZATION_IMAGE_TAG
# }

resource "yandex_serverless_container" "organization-app-container" {
  name               = local.container_name
  service_account_id = yandex_iam_service_account.organization-sa.id
  memory             = 512  # Specify memory in MB
  cores              = 1

  secrets {
    environment_variable = "DATABASE_URL"
    id                   = local.db_url_secret_id
    key                  = "DATABASE_URL"
    version_id           = "e6qbk89fofussm10ksu5"
  }

  secrets {
    environment_variable = "JWT_SECRET"
    id                   = local.jwt_secret_id
    key                  = "JWT_SECRET"
    version_id           = "e6qqunh4ra3ee11b24a8"
  }

  image {
    url = "${local.registry_id}/organization-app:${var.ORGANIZATION_IMAGE_TAG}"
  }
}