mysql:
	goctl model mysql ddl -c --home ./.template/1.7.0/ --src ./apps/anime/rpc/desc/sql/anime.sql --dir ./apps/anime/rpc/model/ \
	goctl model mysql ddl -c --home ./.template/1.7.0/ --src ./apps/anime/rpc/desc/sql/tags.sql --dir ./apps/anime/rpc/model/ \
	goctl model mysql ddl -c --home ./.template/1.7.0/ --src ./apps/anime/rpc/desc/sql/anime_tags.sql --dir ./apps/anime/rpc/model/ 

	goctl model mysql ddl -c --home ./.template/1.7.0/ --src ./apps/stats/rpc/desc/sql/stats.sql --dir ./apps/stats/rpc/model/ 


api:
	goctl api go -api ./apps/app/anitale.api -dir ./apps/app/

rpc:
	goctl rpc protoc ./apps/anime/rpc/desc/proto/anime.proto --go_out=./apps/anime/rpc/ --go-grpc_out=./apps/anime/rpc/ --zrpc_out=./apps/anime/rpc/ --client=true \

	goctl rpc protoc ./apps/stats/rpc/desc/proto/stats.proto --go_out=./apps/stats/rpc/ --go-grpc_out=./apps/stats/rpc/ --zrpc_out=./apps/stats/rpc/ --client=true