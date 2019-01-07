FROM golang:alpine

ENV workspace /go/src/multichat/

RUN mkdir -p $workspace
COPY ./* /$workspace
WORKDIR $workspace

CMD ["/go/src/multichat/multichat_linux_amd64"]