FROM gcr.io/distroless/static:latest
COPY bin/hcl2json_linux_amd64 /hcl2json
ENTRYPOINT ["/hcl2json"]
