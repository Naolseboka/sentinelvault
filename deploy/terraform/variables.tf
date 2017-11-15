variable "project_name" {
  type    = string
  default = "sentinelvault"
}

variable "postgres_password" {
  type      = string
  sensitive = true
  default   = "sentinel"
}

variable "minio_root_password" {
  type      = string
  sensitive = true
  default   = "sentinelvault"
}
