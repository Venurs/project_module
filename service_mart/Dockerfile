FROM daocloud.io/library/ubuntu:16.04
FROM daocloud.io/library/python:3.6.2

ENV LANG C.UTF-8

WORKDIR app/

COPY sources.list /etc/apt

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 40976EAF437D05B5 3B4FE6ACC0B21F32 \
 && apt-get update && apt-get install --force-yes -y \
 python3-dev \
 python3-setuptools \
 libmysqlclient-dev \
 supervisor \
 libsm6 \
 libxrender1 \
 libxext-dev \
 && apt-get --force-yes -y autoremove \
 && apt-get --force-yes -y autoclean \
 && rm -rf /var/lib/apt/lists/*

COPY requirements.txt .

RUN pip install --upgrade pip -i http://mirrors.aliyun.com/pypi/simple/ \
 && pip install -r requirements.txt -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY src/ .
COPY run_app.sh .
RUN mkdir -p /django_file/
