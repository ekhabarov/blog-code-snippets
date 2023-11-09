load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cst",
    srcs = [
        "cst.go",
        "cst_stringer.go",
    ],
    importpath = "github.com/ekhabarov/blog-code-snippets/how-to-bazel/generate-code-with-stringer/cst",
    visibility = ["//visibility:public"],
)

# Source https://stackoverflow.com/questions/66499040/bazel-build-gogenerate-stringer-stringer-cant-happen-constant-is-not-an-i
genrule(
    name = "stringer",
    srcs = ["cst.go"],
    outs = ["cst_stringer.go"],
    cmd = """
      GO_REL_PATH=`dirname $(location @go_sdk//:bin/go)`
      GO_ABS_PATH=`cd $$GO_REL_PATH && pwd`
      env PATH=$$GO_ABS_PATH HOME=$(GENDIR) \
      $(location @org_golang_x_tools//cmd/stringer) \
          -output=$(OUTS) \
          -type=MyType $(location cst.go)
    """,
    tools = [
        "@go_sdk//:bin/go",
        "@org_golang_x_tools//cmd/stringer",
    ],
)
