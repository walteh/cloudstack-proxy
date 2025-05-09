# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# Build example: docker build -t <name> .

FROM node:14-bullseye AS build

LABEL Description="Apache CloudStack UI; Modern role-base progressive UI for Apache CloudStack"
LABEL Vendor="Apache.org"
LABEL License=ApacheV2
LABEL Version=0.5.0
LABEL Author="Apache CloudStack <dev@cloudstack.apache.org>"

WORKDIR /build

RUN apt-get -y update && apt-get -y upgrade

COPY . /build/
RUN npm install
RUN npm run build

FROM nginx:alpine AS runtime

LABEL org.opencontainers.image.title="Apache CloudStack UI" \
	org.opencontainers.image.description="A modern role-based progressive CloudStack UI" \
	org.opencontainers.image.authors="Apache CloudStack Contributors" \
	org.opencontainers.image.url="https://github.com/apache/cloudstack" \
	org.opencontainers.image.documentation="https://github.com/apache/cloudstack/blob/main/ui/README.md" \
	org.opencontainers.image.source="https://github.com/apache/cloudstack" \
	org.opencontainers.image.vendor="The Apache Software Foundation" \
	org.opencontainers.image.licenses="Apache-2.0" \
	org.opencontainers.image.ref.name="latest"

COPY <<EOF /etc/nginx/conf.d/default.conf
server {
    listen       80;
    server_name  localhost;
    location / {
        root   /usr/share/nginx/html;
        index  index.html;
    }
    location /client/ {
        proxy_pass   http://host.docker.internal:8080;
    }
}
EOF

COPY --from=build /build/dist/. /usr/share/nginx/html/
