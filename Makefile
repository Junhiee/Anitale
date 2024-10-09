mysql:
	goctl model mysql ddl -c --home ./.template/1.7.0/ --src ./apps/anime/rpc/desc/sql/anime.sql --dir ./apps/anime/rpc/model/ 

api:
	goctl api go -api ./apps/app/api/anime.api -dir ./apps/app/

rpc:
	goctl rpc protoc ./apps/anime/rpc/desc/proto/anime.proto --go_out=./apps/anime/rpc/ --go-grpc_out=./apps/anime/rpc/ --zrpc_out=./apps/anime/rpc/ 