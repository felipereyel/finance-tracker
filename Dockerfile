# Download the latest version of PocketBase
FROM  --platform=linux/x86_64 alpine:latest as pocketbase
ARG PB_VERSION=0.13.4

RUN apk add --no-cache unzip ca-certificates
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

# Build the Vue app
FROM  --platform=linux/x86_64 node:18-alpine as vueapp

COPY package.json package-lock.json ./
RUN npm install

COPY ./src ./src
COPY ./public ./public
COPY index.html tsconfig.json tsconfig.node.json vite.config.ts ./
RUN npm run build


# Build the final image
FROM  --platform=linux/x86_64 alpine:latest as release

COPY --from=pocketbase /pb/pocketbase /pb
COPY --from=vueapp /dist /pb_public
# database data at /pb_data - mount as volume
# migrations at /pb_migrations - TODO

# start PocketBase instance
CMD ["/pb", "serve", "--http=0.0.0.0:8080"]