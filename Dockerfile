FROM golang:1.22-alpine AS builder
WORKDIR /app
# Copia os arquivos necessários
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Compila o binário com flags de segurança e otimização
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o vex ./main.go

# Imagem final com git
FROM alpine:3.20
RUN apk add --no-cache git curl jq \
    && git config --system user.email "vex@kube.labs" \
    && git config --system user.name "Vex"

ENV PATH="/usr/local/bin:${PATH}"
COPY --from=builder /app/vex /usr/local/bin/vex
ENTRYPOINT ["vex"]
