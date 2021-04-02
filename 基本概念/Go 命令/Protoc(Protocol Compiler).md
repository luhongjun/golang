# Protoc(Protocol Compiler)

- [Google官网 - protocol buffet文档](https://developers.google.cn/protocol-buffers/)
- [Go Generated Code](https://developers.google.cn/protocol-buffers/docs/reference/go-generated)
- [*关于Go proto 语法的官网文档](https://pkg.go.dev/google.golang.org/protobuf/proto#section-documentation)


## [安装](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)

在Ubuntu下可以通过`apt-get install protoc`进行安装；
通过go来安装`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

或者根据文档的介绍：

The protocol compiler is written in C++. If you are using C++, please follow the C++ Installation Instructions to install protoc along with the C++ runtime.

For non-C++ users, the simplest way to install the protocol compiler is to download a pre-built binary from our release page:

https://github.com/protocolbuffers/protobuf/releases

In the downloads section of each release, you can find pre-built binaries in zip packages: protoc-$VERSION-$PLATFORM.zip. It contains the protoc binary as well as a set of standard .proto files distributed along with protobuf.

If you are looking for an old version that is not available in the release page, check out the maven repo here:

https://repo1.maven.org/maven2/com/google/protobuf/protoc/

These pre-built binaries are only provided for released versions. If you want to use the github master version at HEAD, or you need to modify protobuf code, or you are using C++, it's recommended to build your own protoc binary from source.

If you would like to build protoc binary from source, see the C++ Installation Instructions.


## protoc 指令
``` 
$ protoc -help
Usage: protoc [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  -oFILE,                     Writes a FileDescriptorSet (a protocol buffer,
    --descriptor_set_out=FILE defined in descriptor.proto) containing all of
                              the input files to FILE.
  --include_imports           When using --descriptor_set_out, also include
                              all dependencies of the input files in the
                              set, so that the set is self-contained.
  --include_source_info       When using --descriptor_set_out, do not strip
                              SourceCodeInfo from the FileDescriptorProto.
                              This results in vastly larger descriptors that
                              include information about the original
                              location of each decl in the source file as
                              well as surrounding comments.
  --dependency_out=FILE       Write a dependency output file in the format
                              expected by make. This writes the transitive
                              set of input file paths to FILE
  --error_format=FORMAT       Set the format in which to print errors.
                              FORMAT may be 'gcc' (the default) or 'msvs'
                              (Microsoft Visual Studio format).
  --print_free_field_numbers  Print the free field numbers of the messages
                              defined in the given proto files. Groups share
                              the same field number space with the parent
                              message. Extension ranges are counted as
                              occupied fields numbers.

  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --javanano_out=OUT_DIR      Generate Java Nano source file.
  --js_out=OUT_DIR            Generate JavaScript source.
  --objc_out=OUT_DIR          Generate Objective C header and source.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.d
```

- 示例
``` 
//
protoc -I. -I./vendor -I./proto/ykproto -I./../../.. --go_out=plugins=grpc:. ./proto/message_center_proto/mc_email/*.proto
//
protoc -I. -I./vendor -I./proto/ykproto -I./../../.. --grpc-gateway_out=logtostderr=true:. ./proto/message_center_proto/mc_mini_program/*.proto
```

1、protoc指令的格式为 `protoc [参数] .proto文件路径`;

2、`-I[PATH], –proto_path=[PATH]`: 指定 import 修饰符扫描文件夹；可以指定多次，被指定的文件夹将按照先后制定顺序被扫描；如果没有指定，将使用当前文件夹作为扫描文件夹.

3、`–java_out=[PATH]`: 指定生成的java文件的输出文件夹

4、`–go_out=[PATH]`: 指定生成的go文件的输出文件夹