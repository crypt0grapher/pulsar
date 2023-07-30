FROM ubuntu:latest
ARG ETH_RPC_URL
RUN apt-get update && apt-get install -y curl
RUN curl https://raw.githubusercontent.com/a16z/helios/master/heliosup/install | bash
RUN /root/.helios/bin/heliosup
RUN /root/.helios/bin/helios --execution-rpc $ETH_RPC_URL -p 9545