FROM golang

USER root
WORKDIR /app

COPY . .

#RUN apt-get update && apt-get install -y docker.io

RUN cd backend && go mod tidy