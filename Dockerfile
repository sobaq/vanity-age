FROM golang:1.19.2-bullseye
WORKDIR /go/src/github.com/seaofmars/vanity-age
COPY . ./
RUN go install
# ENTRYPOINT ["/bin/sh", "-c", "/go/bin/vanity-age | /usr/bin/tee"]
RUN \
    echo "#!/bin/sh" > /entrypoint.sh &&\
    echo "/go/bin/vanity-age \$1 | /usr/bin/tee /key.txt" >> /entrypoint.sh &&\
    # echo "echo $1" >> /entrypoint.sh &&\
    chmod 755 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

