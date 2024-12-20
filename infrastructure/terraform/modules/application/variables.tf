variable "folder_id" {
    description = "ID of the folder"
    type        = string
}

variable "container_name" {
    description = "Name of the container"
    type        = string
}

variable "db_url_secret_id" {
    description = "ID of the secret with database URL"
    type        = string
}

variable "db_url_secret_version_id" {
    description = "Version ID of the secret with database URL"
    type        = string
}

variable "jwt_secret_id" {
    description = "ID of the secret with JWT secret"
    type        = string
}

variable "jwt_secret_version_id" {
    description = "Version ID of the secret with JWT secret"
    type        = string
}

variable "registry_id" {
    description = "ID of the container registry"
    type        = string
}

variable "ORGANIZATION_IMAGE_TAG" {
  description = "Tag of the image. Set by the CI/CD system"
  type      = string
}
