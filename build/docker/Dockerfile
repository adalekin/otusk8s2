FROM golang:1.16-alpine
RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN make build-linux

FROM scratch
WORKDIR /root/
COPY --from=0 /app/bin .
CMD ["./otusk8s2"]
