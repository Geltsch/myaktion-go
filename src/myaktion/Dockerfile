FROM golang:1.22
WORKDIR /go/src/app
COPY . . 
RUN go mod download
RUN go install
RUN chmod +x ./wait-for-it.sh ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["myaktion"] 
EXPOSE 8000