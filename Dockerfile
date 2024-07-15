# Use a imagem base oficial do Golang
FROM golang:1.18-alpine

# Instale gcc, sqlite e outros pacotes necessários
RUN apk update && apk add --no-cache gcc g++ libc-dev git sqlite

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie o go.mod e go.sum para o diretório de trabalho
COPY go.mod ./
COPY go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código do aplicativo para o diretório de trabalho
COPY . .

# Compile o aplicativo Go
RUN go build -o app .

# Comando para rodar o aplicativo
CMD ["./app"]
