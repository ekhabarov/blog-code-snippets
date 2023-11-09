load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/ekhabarov/blog-code-snippets/how-to-bazel/generate-code-with-stringer
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(name = "gazelle")

go_library(
    name = "generate-code-with-stringer_lib",
    srcs = ["main.go"],
    importpath = "github.com/ekhabarov/blog-code-snippets/how-to-bazel/generate-code-with-stringer",
    visibility = ["//visibility:private"],
    deps = ["//cst"],
)

go_binary(
    name = "generate-code-with-stringer",
    embed = [":generate-code-with-stringer_lib"],
    visibility = ["//visibility:public"],
)
