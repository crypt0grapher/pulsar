FROM golang:latest
ARG LIGHTNODE_RPC_URL
RUN mkdir /app
COPY . /app
COPY ./.etherlink /root/.etherlink
WORKDIR /app
RUN curl https://get.ignite.com/cli | bash
RUN ./ignite chain build
CMD /go/bin/etherlinkd start