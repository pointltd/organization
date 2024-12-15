provider "yandex" {
  zone      = var.zone
  folder_id = var.folder_id
  max_retries = 3
}
