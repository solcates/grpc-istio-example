FROM golang:1.12 as builder
ENV GOOS=linux
ENV GOARC=amd64
ENV CGO_ENABLED=0
ARG GIT_COMMIT=unknown
ADD vendor /app/vendor
ADD Makefile /app/Makefile
ADD apis /app/apis
ADD pkg /app/pkg
ADD cmd/greeter /app/cmd/greeter
ADD go.* /app/
WORKDIR /app
RUN make build

FROM scratch
COPY --from=builder /app/greeter /greeter
ENTRYPOINT ["/greeter"]
CMD ["server"]
