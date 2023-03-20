# programme.lv

`programme.lv` is a modern latvian programming education platform. The frontend is a stateless `next.js` service, the backend is written in `go`. In front of the frontend and backend will be sitting a reverse proxy `caddy` that will route `/api/*` requests to backend and others to frontend. The backend can communicate with the `postgres` database. Backend consists of the `controller` that routes the api requests and works with database and the `scheduler` that publishes jobs to the `worker`s.
 
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