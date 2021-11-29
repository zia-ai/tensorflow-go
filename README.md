# Tensorflow Go Bindings
This is a fork of tensorflow's repository (at version 2.7.0) with pre-generated protobuf definitions.

Specificially, this uses protobufs generated with protoc-gen-go v1.3.5 (before the api refactor)

# Matching libs
| TensorFlow | C library | URL |
| ---------- | --------- | --- |
| Linux || |
| Linux | CPU only	| https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.7.0.tar.gz |
| Linux | GPU support	| https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-2.7.0.tar.gz |
| macOS || |
| macOS | CPU only |	https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.7.0.tar.gz |
| Windows || |
| Windows | CPU only | https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-windows-x86_64-2.7.0.zip |
| Windows | GPU only | https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-windows-x86_64-2.7.0.zip |

# To generate on MacOS

1. Install dependencies

     `brew install bazel swig`

     `pip3 install --upgrade pip`

     `pip3 install numpy`

4. Get transforflow in go, build library
    a. `go get -d github.com/tensorflow/tensorflow/tensorflow/go`

    b. `cd ${GOPATH}/src/github.com/tensorflow/tensorflow`

    c. edit `.bazelversion`, change to value of `bazel --version`

    d. `./configure`

    e. `bazel build --config=macos_arm64 tensorflow/tools/lib_package:libtensorflow `

    f. `sudo tar -C /usr/local -xzf bazel-bin/tensorflow/tools/lib_package/libtensorflow.tar.gz`

5. Generate protos

    a. `go generate github.com/tensorflow/tensorflow/tensorflow/go/op`

    b. if doesn't work, `cd ${GOPATH}/src/github.com/tensorflow/tensorflow/go`
       then try to run `./genop/generate.sh`

    c. if failing, modify the file (may have to modify path to tf)

    d. may fail in status.proto, copy the go package to it like the other files in same directory

    e. copy generated files:

      `cp -R vendor/github.com/tensorflow/tensorflow/tensorflow/go/core core`

      `cp -R vendor/github.com/tensorflow/tensorflow/tensorflow/go/stream_executor stream_executor`

6. Checkout https://github.com/zia-ai/tensorflow-go

   `rsync -avz --delete ${GOPATH}/pkg/mod/github.com/tensorflow/tensorflow@v2.7.0+incompatible/tensorflow/go tensorflow/`

    `git checkout -b v2.7.0`

    commit & push

7. In your project
    `go get -d github.com/zia-ai/tensorflow-go@407003c`

    it will fail, take note of the version id (v1.0.1-0.20211126165055-407003c9e029)

    replace in go.mod

    replace github.com/tensorflow/tensorflow v2.7.0+incompatible => github.com/zia-ai/tensorflow-go v1.0.1-0.20211126165055-407003c9e029
