terraform {
  required_version = ">= 1.5.0"
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

provider "docker" {}

module "postgres" {
  source           = "./modules/postgres"
  project_name     = var.project_name
  postgres_password = var.postgres_password
}

module "redis" {
  source       = "./modules/redis"
  project_name = var.project_name
}

module "minio" {
  source              = "./modules/minio"
  project_name        = var.project_name
  minio_root_password = var.minio_root_password
}

output "postgres_container" {
  value = module.postgres.container_name
}

output "redis_container" {
  value = module.redis.container_name
}

output "minio_endpoint" {
  value = module.minio.endpoint
}
