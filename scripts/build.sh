#!/bin/sh --posix

Prefix=$1

Services=$(basename $(pwd))
Microservice=$(basename cmd/$2)
ServicesVersion=$(cat .version)

BINARY="bin/${Prefix}-${Microservice}"
SOURCE="cmd/${Microservice}/main.go"

ENV=$(uname -snr)
UPT=$(date +"%Y/%m/%d %H:%M:%S%z")
VER=$(cat internal/app/${Microservice}/.version)
TAG=$(echo $(git rev-parse --short HEAD)$([ -n "$(git status -s)" ] && echo "-dev" || echo ""))
BRS=$(echo $([ -n "$(git branch)" ] && echo "$(git symbolic-ref --short -q HEAD)" || echo "unknow"))

echo "Release ${Prefix}-${Microservice} :${ServicesVersion}.${VER}@${TAG}"

go build -tags netgo                                                  \
    -installsuffix 'static'                                           \
    -ldflags "                                                        \
    -s -w                                                             \
    -X '$(go list -m)/pkg/info.verStr=${VER}'                         \
    -X '$(go list -m)/pkg/info.brsStr=${BRS}'                         \
    -X '$(go list -m)/pkg/info.tagStr=${TAG}'                         \
    -X '$(go list -m)/pkg/info.uptStr=${UPT}'                         \
    -X '$(go list -m)/pkg/info.envStr=${ENV}'                         \
    -X '$(go list -m)/pkg/info.Prefix=${Prefix}'                      \
    -X '$(go list -m)/pkg/info.Services=${Services}'                  \
    -X '$(go list -m)/pkg/info.MicroService=${Microservice}'          \
    -X '$(go list -m)/pkg/info.ServicesVersion=${ServicesVersion}'    \
    "                                                                 \
    -o ${BINARY} ${SOURCE}
