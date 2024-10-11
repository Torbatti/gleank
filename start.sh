(cd ./views && templ generate)

(cd ./models/sqlc && sqlc generate)

go run ./pkgs/base/main.go serve

