load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "github.com/ekhabarov/blog-code-snippets/bazel-k8s-envoy/service-one/server",
    deps = ["//service-one/pb"],
)
