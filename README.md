# programme.lv

`programme.lv` is a modern latvian programming education platform.

The frontend is served by a stateless `next.js` service. The backend is written in `go` and consists mainly of the `controller` and the `scheduler`. The `controller` routes incoming requests and communicates with the `postgres` database. The `scheduler` publishes jobs to `worker`s.

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

```
cd backend
go run .
```