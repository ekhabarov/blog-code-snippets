load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/ekhabarov/helloworld-generator
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(name = "gazelle")

genrule(
    name = "generate_hw",
    outs = ["hw.go"],
    cmd = "$(location @com_github_ekhabarov_helloworld_generator//:helloworld-generator) > $(OUTS)" ,
    tools = [
        "@com_github_ekhabarov_helloworld_generator//:helloworld-generator",
    ],
)

go_library(
    name = "lib",
    srcs = [":generate_hw"],
    importpath = "github.com/ekhabarov/helloworld-generator",
)

go_binary(
    name = "bin",
    embed = ["lib"],
    importpath = "github.com/ekhabarov/helloworld-generator",
)
