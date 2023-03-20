# programme.lv

Modern programming education platforma

![uzdevumi](./tasks-screenshot.png)
![diagramma](./diagram.svg)

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