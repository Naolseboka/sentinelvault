resource "docker_image" "postgres" {
  name = "postgres:16-alpine"
}

resource "docker_container" "postgres" {
  name  = "${var.project_name}-postgres"
  image = docker_image.postgres.image_id
  env = [
    "POSTGRES_USER=sentinel",
    "POSTGRES_PASSWORD=${var.postgres_password}",
    "POSTGRES_DB=${var.project_name}",
  ]
  ports {
    internal = 5432
    external = 5432
  }
}

output "container_name" {
  value = docker_container.postgres.name
}
