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
	supervisor
# python3-dev \
# python-is-python3 \
# python3-setuptools \
# python3-pip \
# python3-mysql.connector \
# supervisor

RUN apt-get install -qqy mysql-server && \
    apt-get clean all && \
    mkdir -p /var/run/mysqld; \
    chown mysql /var/run/mysqld

RUN echo '''sql_mode = "STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION"''' >> /etc/mysql/mysql.conf.d/mysqld.cnf

COPY tools/docker/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY . ./root
WORKDIR /root

# COPY <<EOF /etc/supervisor/conf.d/supervisord.conf
# [supervisord]
# nodaemon=true

# [program:mysqld]
# command=/usr/bin/mysqld_safe
# autostart=true
# autorestart=true
# user=root
# EOF

# [program:cloudstack-ui]
# command=/bin/bash -c "npm run serve"
# directory=/root/ui
# stdout_logfile=/dev/stdout
# stdout_logfile_maxbytes=0
# user=root

COPY --from=maven . /root/.m2

RUN <<EOF
ls -la /root/.m2
mvn -pl developer -Pdeveloper -Dsimulator -DskipTests -T 4C clean install
EOF

RUN find /var/lib/mysql -type f -exec touch {} \; && \
    (/usr/bin/mysqld_safe &) && \
    sleep 5; \
    mysql -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password by ''" --connect-expired-password; \
    mvn -Pdeveloper -pl developer -Ddeploydb; \
    mvn -Pdeveloper -pl developer -Ddeploydb-simulator;

# RUN curl -fsSL https://deb.nodesource.com/setup_18.x -o nodesource_setup.sh \
# 	&& sudo -E bash nodesource_setup.sh \
# 	&& apt-get install -y nodejs make

# RUN cd ui && npm install

VOLUME /var/lib/mysql

EXPOSE 8096

CMD ["/usr/bin/mysqld_safe"]
