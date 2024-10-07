FROM golang

USER root
WORKDIR /app

COPY . .

RUN cd backend && go mod tidy