# Download the latest version of PocketBase
FROM  --platform=linux/x86_64 alpine:latest as pocketbase
ARG PB_VERSION=0.13.4

RUN apk add --no-cache unzip ca-certificates
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

# Build the final image
FROM  --platform=linux/x86_64 alpine:latest as release

COPY --from=pocketbase /pb/pocketbase /pb
# public files to serve at /pb_public - mount as volume
# database data at /pb_data - mount as volume
# migrations at /pb_migrations - TODO

# start PocketBase instance
CMD ["/pb", "serve", "--http=0.0.0.0:8080"]