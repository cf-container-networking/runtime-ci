FROM ubuntu:trusty

RUN \
  apt-get update && \
  apt-get -y install \
    build-essential \
    wget \
    curl \
    openssh-client

ENV GOPATH $HOME/go
ENV PATH $HOME/go/bin:/usr/local/go/bin:$PATH
RUN \
  wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz -P /tmp && \
  tar xzvf /tmp/go1.6.2.linux-amd64.tar.gz -C /usr/local && \
  mkdir $GOPATH && \
  rm -rf /tmp/*

RUN wget -q -O cf.deb "https://cli.run.pivotal.io/stable?release=debian64&version=6.17.1&source=github-rel" && \
  dpkg -i cf.deb
