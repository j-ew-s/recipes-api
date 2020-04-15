# PASSO 1
# Baixar a ultima versao de golang
FROM golang:latest

# PASSO 2
# Para montarmos nosso ambiente temos que criar um 
# diretorio dentro do $GOPATH como fazemos em DEV.

RUN mkdir -p /go/src/github.com/j-ew-s/recipes-api

# PASSO 3
# Direcionamos nosso WORKDIR para a pasta criada
WORKDIR /go/src/github.com/j-ew-s/recipes-api

# PASSO 4
# Copiamos TUDO . (usando ponto) que está no nosso diretorio do Dockerfile (arquivos e pastas)
#  para a estrutura de pasta criada. no passo dois
ADD . /go/src/github.com/j-ew-s/recipes-api

# PASSO 5
# Agora movemos para a pasta CMD que contem o MAIN.GO
WORKDIR /go/src/github.com/j-ew-s/recipes-api/cmd

# PASSO 6
# Baixar TODAS as dependencias necessárias do projeto
RUN go get -v

# PASSO 7
# Correr o projeto main
# correr projeto sem parametro : dev (8061)
# correr projeto com parametro : qa (8062)
# correr projeto com parametro : prod (8063)
RUN go build main.go 

# PASSO 8 
# Expoe a porta 8061 ja que o default é dev (pois o main.go nao usa parametro de qa ou prod)
EXPOSE 8061

CMD ["/go/src/github.com/j-ew-s/recipes-api/cmd/main"]
