FROM golang:1.22.4-bullseye

ARG APP_DIR="/usr/src/app"
ARG WITH_GOTEST="0"

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

# ローカルでテストする場合、gotestを使う
RUN \
  if [ "${WITH_GOTEST}" = "1" ] ; then \
    go install github.com/rakyll/gotest@latest ; \
  fi

RUN go mod tidy
