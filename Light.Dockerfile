FROM rustlang/rust:nightly-bullseye-slim
ARG ETH_RPC_URL
RUN apt-get update && apt-get install -y curl git pkg-config build-essential openssl libssl-dev
RUN rustup default stable
RUN git clone https://github.com/a16z/helios.git
RUN cd helios && cargo build --release
CMD /helios/target/release/helios --rpc-bind-ip "0.0.0.0" --execution-rpc $ETH_RPC_URL --rpc-port 7545 --checkpoint "0x4aef8b169a3ed574581828e80ca99428eecb31a3c204b301bcd5cc9e25352ab0"
