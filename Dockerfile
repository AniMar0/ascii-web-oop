FROM golang:1.22.3

WORKDIR /app

COPY . .


EXPOSE 8080


CMD ["go","run","main.go"]



#docker build -t ascii-art-web .
#docker run -d -p 8080:8080 ascii-art-web
