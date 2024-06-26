FROM golang:1.22.4-bullseye

ARG APP_DIR="/usr/src/app"

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
