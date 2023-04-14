load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

genrule(
    name = "generate_hello_go",
    outs = ["hello.go"],
    cmd = "$(execpath @com_github_ekhabarov_helloworld_generator//:helloworld-generator) > $@",
    tools = [
        "@com_github_ekhabarov_helloworld_generator//:helloworld-generator",
    ],
)

go_library(
    name = "hello-world_lib",
    srcs = ["hello.go"],
    importpath = "github.com/ekhabarov/helloworld-generator",
)

go_binary(
    name = "hello-world",
    embed = [":hello-world_lib"],
    importpath = "github.com/ekhabarov/helloworld-generator",
)
