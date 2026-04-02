FROM ghcr.io/webassembly/wasi-sdk:wasi-sdk-32

RUN apt-get update && apt-get install -y binaryen git patch

# Easiest way to build grpc-java plugin is add it to the protoc build
ADD buildtools/wasm/version.txt /version.txt
RUN git clone --recursive https://github.com/grpc/grpc-java.git --branch $(cat /version.txt | awk '{$1=$1};1') /grpcjava
RUN git clone --recursive https://github.com/protocolbuffers/protobuf.git --branch v34.1 /workspace
RUN cp -R /grpcjava/compiler/src/java_plugin/cpp /workspace/src/grpcjava
RUN echo "set(protoc-gen-grpc-java_files \${protobuf_SOURCE_DIR}/src/grpcjava/java_plugin.cpp \${protobuf_SOURCE_DIR}/src/grpcjava/java_generator.cpp) \n\
add_executable(protoc-gen-grpc-java \${protoc-gen-grpc-java_files} \${protobuf_version_rc_file}) \n\
target_link_libraries(protoc-gen-grpc-java libprotoc libprotobuf \${protobuf_ABSL_USED_TARGETS}) \n\
set_target_properties(protoc-gen-grpc-java PROPERTIES VERSION $(cat /version.txt)) \n\
" >> /workspace/CMakeLists.txt

WORKDIR /workspace

ADD buildtools/wasm/patch-absl.txt patch-absl.txt

# libprotoc, the library that code generator plugins link against, includes the command line interface, including
# subprocess execution. Code generator plugins never need this, so we can remove the files to allow the project to
# build with wasi-sdk.
RUN rm src/google/protobuf/compiler/subprocess.* src/google/protobuf/compiler/command_line_interface.*
RUN sed -i '/src\/google\/protobuf\/compiler\/subprocess\./d' src/file_lists.cmake
RUN sed -i '/src\/google\/protobuf\/compiler\/command_line_interface\./d' src/file_lists.cmake

ENV CMAKE_TOOLCHAIN_FILE=/opt/wasi-sdk/share/cmake/wasi-sdk-pthread.cmake
ENV CFLAGS="-mllvm -wasm-enable-sjlj -DWASM_WAMR -D_WASI_EMULATED_MMAN -D_WASI_EMULATED_PROCESS_CLOCKS -D_WASI_EMULATED_SIGNAL -DABSL_HAVE_MMAP -DABSL_FORCE_THREAD_IDENTITY_MODE=1"
ENV CXXFLAGS="$CFLAGS -fno-exceptions"
ENV LDFLAGS="-lwasi-emulated-process-clocks -lwasi-emulated-mman -lwasi-emulated-signal -Wl,--max-memory=4294967296"

RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_C_FLAGS="$CFLAGS" -DCMAKE_CXX_FLAGS="$CXXFLAGS" -DCMAKE_EXE_LINKER_FLAGS="$LDFLAGS" -S . -B build

# cmake automatically downloads abseil in previous step so we patch here before then building
RUN cd build/_deps/absl-src/ && patch -p1 < ../../../patch-absl.txt

RUN cmake --build build --target protoc-gen-grpc-java -j 2

RUN wasm-opt -o build/protoc-gen-grpc-java.wasm --flatten --rereloop --converge -O3 build/protoc-gen-grpc-java

CMD ["cp", "build/protoc-gen-grpc-java.wasm", "/out/"]
