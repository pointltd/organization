resource "yandex_serverless_container" "organization-app-container" {
  name               = var.container_name
  service_account_id = yandex_iam_service_account.organization-sa.id
  memory             = 512  # Specify memory in MB
  cores              = 1

  secrets {
    environment_variable = "DATABASE_URL"
    id                   = var.db_url_secret_id
    key                  = "DATABASE_URL"
    version_id           = var.db_url_secret_version_id
  }

  secrets {
    environment_variable = "JWT_SECRET"
    id                   = var.jwt_secret_id
    key                  = "JWT_SECRET"
    version_id           = var.jwt_secret_version_id
  }

  image {
    url = "${var.registry_id}/organization-app:${var.ORGANIZATION_IMAGE_TAG}"
  }
}

provider "yandex" {
  zone      = "ru-central1-a"
  folder_id = "b1g0k22us62vt6kut949"
}
