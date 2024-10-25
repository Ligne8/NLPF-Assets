# Utiliser l'image Go officielle
FROM golang:1.23.2

# Répertoire de travail
WORKDIR /app

# Copier les fichiers de l'application
COPY . .

# Compiler l'application
RUN go mod tidy
RUN go build -o app

# Exposer le port sur lequel l'application Go va écouter
EXPOSE 8080

# Commande de démarrage
CMD ["./app"]