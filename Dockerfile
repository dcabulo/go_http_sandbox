FROM golang:1.19

WORKDIR /app

COPY . .

ENV GOOSE_DRIVER=postgres
ENV GOOSE_DBSTRING="user=postgres password=password host=changeHostname dbname=rssagg sslmode=disable"

RUN go install github.com/pressly/goose/v3/cmd/goose@latest 
RUN goose -dir "./sql/schema" up

RUN go build -o /rssagg

EXPOSE 808

CMD [ "/rssagg" ]