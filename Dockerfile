FROM alpine:latest
COPY dist/ocictl /bin/ocictl

FROM abegolli/ocibuilder-base
COPY --from=ocictl /bin/ocictl /bin
