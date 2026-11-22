
job "springer-demo" {
  datacenters = ["dc1"]
  type        = "service"

  group "springer-demo-group" {
    count = 1

    network {
      port "http" {
        static = 8080
      }
    }

    task "springer-demo-app" {
      driver = "docker"

      config {
        image = "ghcr.io/shuayeb-x/springer-demo-app:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }

      service {
        name     = "springer-demo"
        provider = "nomad"
        port     = "http"

        check {
          name     = "application-health"
          type     = "http"
          path     = "/health"
          interval = "10s"
          timeout  = "3s"
        }
      }
    }
  }
}
