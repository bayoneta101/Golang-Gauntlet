group "default" {
  targets = ["client", "server"]
}

target "client" {
  context    = "."
  dockerfile = "Dockerfile"
  args = {
    SERVICE = "client"
  }
  tags = ["gauntlet-client:latest"]
}

target "server" {
  context    = "."
  dockerfile = "Dockerfile"
  args = {
    SERVICE = "server"
  }
  tags = ["gauntlet-server:latest"]
}
