FROM alpine
ADD entrypoint.sh /entrypoint.sh
RUN mkdir wasm-opt && chmod +x /entrypoint.sh
ADD wasm-opt /wasm-opt/wasm-opt 
ENV PATH $PATH:/wasm-opt
ENTRYPOINT ["/entrypoint.sh"]
