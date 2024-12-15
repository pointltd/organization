resource "yandex_iam_service_account" "organization-sa" {
  name = "organization-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "registry_pull_permission" {
  folder_id = var.folder_id
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.organization-sa.id}"
}

resource "yandex_lockbox_secret_iam_binding" "lockbox_db_url_viewer_permission" {
  secret_id = var.db_url_secret_id
  role      = "lockbox.payloadViewer"
  members   = ["serviceAccount:${yandex_iam_service_account.organization-sa.id}"]
}

resource "yandex_lockbox_secret_iam_binding" "lockbox_jwt_secret_id_viewer_permission" {
  secret_id = var.jwt_secret_id
  role      = "lockbox.payloadViewer"
  members   = ["serviceAccount:${yandex_iam_service_account.organization-sa.id}"]
}