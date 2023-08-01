FROM golang:latest
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN curl https://get.ignite.com/cli | bash
RUN ./ignite chain build
CMD /go/bin/etherlinkd start