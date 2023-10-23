go_library(
    name = "sandbox",
    srcs = [
        "sandbox_linux.go" if CONFIG.OS == "linux" else "sandbox_other.go",
    ],
    pgo_file = "//:pgo",
    visibility = ["//src/..."],
    deps = [
        "///third_party/go/golang.org_x_sys//unix" if CONFIG.OS == "linux" else "",
        "//src/core",
    ],
    
)
