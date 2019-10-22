FROM busybox AS binary
COPY dist/ocictl /bin/ocictl

FROM abegolli/ocibuilder-base:v0.1.0
COPY --from=binary /bin/ocictl /bin
