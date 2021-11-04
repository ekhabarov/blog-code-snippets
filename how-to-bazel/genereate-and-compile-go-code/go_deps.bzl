load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_deps():
    go_repository(
        name = "com_github_ekhabarov_helloworld_generator",
        importpath = "github.com/ekhabarov/helloworld-generator",
        sum = "h1:MrREQgX6I0/4cstUhbuqfALzUF3W2Nz8kVZRq6A4q+E=",
        version = "v0.0.1",
    )
