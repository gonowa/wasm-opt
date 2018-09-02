FROM alpine
ADD entrypoint.sh /entrypoint.sh
RUN mkdir wasm-opt && chmod +x /entrypoint.sh
RUN wget -qO- https://github.com/WebAssembly/binaryen/releases/download/version_50/binaryen-version_50-x86_64-linux.tar.gz | tar xvz -C /wasm-opt binaryen-version_50/wasm-opt --strip=1
ENV PATH $PATH:/wasm-opt
ENTRYPOINT ["/entrypoint.sh"]
