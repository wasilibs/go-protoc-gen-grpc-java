FROM ghcr.io/wasilibs/wasix-sdk:sha-fc94d60

RUN apt-get update && apt-get install -y binaryen git patch

# Easiest way to build grpc-java plugin is add it to the protoc build
ADD buildtools/wasm/version.txt /version.txt
RUN git clone --recursive https://github.com/grpc/grpc-java.git --branch $(cat /version.txt | awk '{$1=$1};1') /grpcjava
RUN git clone --recursive https://github.com/protocolbuffers/protobuf.git --branch v26.1 /workspace
RUN cp -R /grpcjava/compiler/src/java_plugin/cpp /workspace/src/grpcjava
RUN echo "set(protoc-gen-grpc-java_files \${protobuf_SOURCE_DIR}/src/grpcjava/java_plugin.cpp \${protobuf_SOURCE_DIR}/src/grpcjava/java_generator.cpp) \n\
add_executable(protoc-gen-grpc-java \${protoc-gen-grpc-java_files} \${protobuf_version_rc_file}) \n\
target_link_libraries(protoc-gen-grpc-java libprotoc libprotobuf \${protobuf_ABSL_USED_TARGETS}) \n\
set_target_properties(protoc-gen-grpc-java PROPERTIES VERSION 1.60.0) \n\
" >> /workspace/CMakeLists.txt

WORKDIR /workspace

ADD buildtools/wasm/patch-absl.txt patch-absl.txt
RUN cd third_party/abseil-cpp && patch -p1 < ../../patch-absl.txt

ENV CXXFLAGS "${CXXFLAGS} -O3 -pthread -D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS -DABSL_HAVE_MMAP"
ENV LDFLAGS $LDFLAGS -lwasi-emulated-mman -lwasi-emulated-process-clocks -Wl,--max-memory=4294967296

RUN cmake -S . -B build
RUN cmake --build build --target protoc-gen-grpc-java

RUN wasm-opt -o build/protoc-gen-grpc-java.wasm --low-memory-unused --flatten --rereloop --converge -O3 build/protoc-gen-grpc-java

CMD ["cp",  "build/protoc-gen-grpc-java.wasm", "/out/"]
