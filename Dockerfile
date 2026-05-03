FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY . /app