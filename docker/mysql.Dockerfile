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
#
# CloudStack-simulator build

FROM ubuntu:22.04

LABEL Vendor="Apache.org" License="ApacheV2" Version="4.21.0.0-SNAPSHOT" Author="Apache CloudStack <dev@cloudstack.apache.org>"

ARG DEBIAN_FRONTEND=noninteractive

# Install dependencies in a single layer
RUN apt-get -y update && apt-get install -y \
	genisoimage \
	libffi-dev \
	libssl-dev \
	curl \
	gcc-10 \
	git \
	sudo \
	ipmitool \
	iproute2 \
	maven \
	openjdk-11-jdk \
	mysql-server \
	&& apt-get clean all \
	&& mkdir -p /var/run/mysqld \
	&& chown mysql /var/run/mysqld

# Configure MySQL
RUN echo 'sql_mode = "STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION"' >> /etc/mysql/mysql.conf.d/mysqld.cnf

# Copy required files
COPY . ./root/cloudstack-source
COPY --from=maven . /root/.m2

WORKDIR /root/cloudstack-source

# Build CloudStack and setup database in a single layer
RUN ls -la /root/.m2 \
	&& mvn -pl developer -Pdeveloper -Dsimulator -DskipTests -T 4C clean install \
	&& find /var/lib/mysql -type f -exec touch {} \; \
	&& (/usr/bin/mysqld_safe &) \
	&& sleep 5 \
	&& mysql -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password by ''" --connect-expired-password \
	&& mvn -Pdeveloper -pl developer -Ddeploydb \
	&& mvn -Pdeveloper -pl developer -Ddeploydb-simulator

WORKDIR /
# Cleanup to reduce image size
RUN rm -rf /root/.m2 \
	&& rm -rf /root/cloudstack-source

# \
# && apt-get remove -y gcc-10 git maven openjdk-11-jdk \
# && apt-get autoremove -y \
# && apt-get clean \
# && rm -rf /var/lib/apt/lists/* \
# && rm -rf /root/.cache \
# && rm -rf /tmp/* \
# && rm -rf /var/tmp/* \
# && find /root -type d -name "target" -exec rm -rf {} + 2> /dev/null || true

VOLUME /var/lib/mysql

EXPOSE 8096

CMD ["/usr/bin/mysqld_safe"]
