# twittor

# Iniciando el proyecto crear el archivo go.mod

module github.com/nelsomar/twittor
go 1.12
require ()

# Instalando dependencias

go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/bson/primitive

go get golang.org/x/crypto/bcrypt

go get github.com/gorilla/mux
go get github.com/rs/cors
go get github.com/dgrijalva/jwt-go
