resource "docker_image" "redis" {
  name = "redis:7-alpine"
}

resource "docker_container" "redis" {
  name  = "${var.project_name}-redis"
  image = docker_image.redis.image_id
  ports {
    internal = 6379
    external = 6379
  }
}

output "container_name" {
  value = docker_container.redis.name
}
