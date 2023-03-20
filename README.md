# programme.lv

`programme.lv` is a modern latvian programming education platform.

The frontend is served by a stateless `next.js` service. The backend is written in `go` and consists mainly of the `controller` and the `scheduler`. The `controller` routes incoming requests and communicates with the `postgres` database. The `scheduler` publishes jobs to `worker`s.

Communication between the frontend and backend is done via REST API. Communication between the `scheduler` and the `worker` is done via gRPC and defined in the `protofiles`.

Live updates to frontend users will be supplied using websockets.

## starting reverse proxy

```
sudo caddy run --config ./caddy.conf --adapter caddyfile
```

## starting frontend

```
cd website
yarn install
yarn run dev
```

## starting backend

Besides having go installed an .env file should be located in the backend folder and contain `db_conn_string`.

```
cd backend
go run scripts/migration/migrate.go
go run .
```

## start worker
```
cd backend
cd worker
go run .
``
