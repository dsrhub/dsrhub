ARG UTASK_VERSION=v1.6.0
FROM ovhcom/utask:${UTASK_VERSION} AS builder

COPY .  /go/src/github.com/ovh/utask
WORKDIR /go/src/github.com/ovh/utask
RUN make

FROM ovhcom/utask:${UTASK_VERSION}

COPY templates  /app/templates
COPY --from=builder /go/src/github.com/ovh/utask/plugins /app/plugins
COPY --from=builder /go/src/github.com/ovh/utask/init    /app/init
COPY --from=builder /go/src/github.com/ovh/utask/utask   /app/utask
RUN chmod +x /app/utask
