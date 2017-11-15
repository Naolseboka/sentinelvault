resource "docker_image" "minio" {
  name = "minio/minio:latest"
}

resource "docker_container" "minio" {
  name  = "${var.project_name}-minio"
  image = docker_image.minio.image_id
  command = ["server", "/data", "--console-address", ":9001"]
  env = [
    "MINIO_ROOT_USER=sentinel",
    "MINIO_ROOT_PASSWORD=${var.minio_root_password}",
  ]
  ports {
    internal = 9000
    external = 9000
  }
}

output "endpoint" {
  value = "http://localhost:9000"
}
