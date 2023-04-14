load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

genrule(
    name = "generate_hello_go",
    outs = ["hello.go"],
    cmd = """
    cat << EOF >> $@
package main

import "fmt"

func main() {
  fmt.Printf("Hello World!")
}
EOF
""",
)

go_library(
    name = "hello-world_lib",
    srcs = ["hello.go"],
    # srcs = [":generate_hello_go"], # genrule name also can be used here.
    importpath = "github.com/ekhabarov/hello-world",
)

go_binary(
    name = "hello-world",
    embed = [":hello-world_lib"],
    importpath = "github.com/ekhabarov/hello-world",
)
