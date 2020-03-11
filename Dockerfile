FROM golang:1.14 as build

COPY . .

# -s -w strips debugging information
RUN go build -ldflags="-s -w" -o /bin/action

# Install upx (upx.github.io) to compress the compiled action
RUN apt-get update && apt-get -y install upx

# Compress the compiled action
RUN upx -q -9 /bin/action

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/action /bin/action

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/action"]