FROM golang:1.6
MAINTAINER tu.ue.et@gmail.com

ENV HOSTING=github.com \
    VCS_USER=hhlab \
    PROJECT=react_hypercard

ENV WORK_DIR=/go/src/${HOSTING}/${VCS_USER}/${PROJECT}

RUN mkdir -p $WORK_DIR
WORKDIR $WORK_DIR
# WORKDIR = /go/src/github.com/hhlab/react_hypercard

# Install tools
RUN go get golang.org/x/tools/cmd/stringer \
    github.com/codegangsta/gin \
    github.com/gin-gonic/gin



# godo hotloader 
ADD entrypoint.sh ${WORK_DIR}/entrypoint.sh
# ADD consumer_entrypoint.sh ${WORK_DIR}/consumer_entrypoint.sh

ADD . $WORK_DIR


