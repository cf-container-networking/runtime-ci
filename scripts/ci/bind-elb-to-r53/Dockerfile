FROM ubuntu:trusty

RUN \
  apt-get update && \
  apt-get -y install \
    build-essential \
    mercurial \
    software-properties-common \
    wget \
    python-pip \
    jq

RUN \
  add-apt-repository ppa:git-core/ppa -y && \
  apt-get update && \
  apt-get install -y git && \
  apt-get remove -y --purge software-properties-common

ENV GOPATH $HOME/go
ENV PATH $HOME/go/bin:/usr/local/go/bin:$PATH
RUN \
  wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz -P /tmp && \
  tar xzvf /tmp/go1.6.2.linux-amd64.tar.gz -C /usr/local && \
  mkdir $GOPATH && \
  rm -rf /tmp/*

RUN \
  wget https://s3.amazonaws.com/bosh-init-artifacts/bosh-init-0.0.91-linux-amd64 -P /tmp && \
  mv /tmp/bosh-init-0.0.91-linux-amd64 /usr/local/bin/bosh-init && \
  chmod +x /usr/local/bin/bosh-init && \
  rm -rf /tmp/*

RUN \
  go get github.com/pivotal-cf-experimental/bosh-bootloader/bbl && \
  go install github.com/pivotal-cf-experimental/bosh-bootloader/bbl

RUN \
  wget https://github.com/postmodern/ruby-install/archive/v0.4.1.tar.gz -P /tmp && \
  tar zxf /tmp/v0.4.1.tar.gz -C /tmp && \
  cd /tmp/ruby-install-0.4.1 && make install && \
  ruby-install ruby 2.1.8 && \
  rm -rf /usr/local/src/ruby-2.1.8 && \
  rm -rf /tmp/*

ENV PATH /opt/rubies/ruby-2.1.8/bin:$PATH

RUN \
  pip install awscli==1.6.6