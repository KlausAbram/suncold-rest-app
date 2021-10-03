FROM alpine:latest

RUN apk --no-cache add ca-certificates  
WORKDIR /root 

COPY ./.bin/app .
COPY ./configs ./configs 
COPY ./.sh ./script
COPY .env ./


RUN apk add postgresql-client

RUN chmod +x ./script/wait-for-postgres.sh

CMD ["./app"]






