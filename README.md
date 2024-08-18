# Pretty link

Simple short link app

## Layout
``` txt
pretty-link
├── app
│   └── app.go
├── auth
│   ├── api
│   │   ├── handler.go
│   │   └── register.go
│   ├── models
│   │   └── models.go
│   ├── repository
│   │   └── postgres.go
│   ├── usecase
│   │   └── usecase.go
│   └── usecase.go
├── cmd
│   └── main.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── migrations
│   └── 20240818142323_create_account_table.sql
└── README.md
```

## Migrations

```bash
$ export GOOSE_DRIVER=postgres
$ export GOOSE_DBSTRING=postgresql://%USER_NAME%:%USER_PASSWORD@HOST%:%PORT%/%DB_NAME%?sslmode=disable
```

```bash
# from project's root
$ goose -dir ./migrations up
```

for more [info](https://github.com/pressly/goose).
