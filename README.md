# Microservices - cours CESI

## Prérequis

- `rust` (stable, 1.66.1 => stable-aarch64-apple-darwin)
- `go` (stable, 1.19.5 => darwin/arm64)
- `node` (lts 18.12.0)

## Lundi 16/01

`rust-sum`: api rest, calcul d'une somme

- utilisation de `rocket`

### Lancement

```sh
cd rust-sum
cargo run
```


## Mardi 17/01

`rust-rpc-todo`: implémentation microservice, d'une communication RPC pour une todolist

- utilisation de `tonic`/`tokio`/`prost`, expression du protobuf `proto/todo.proto`

### Lancement

```sh
cd rust-rpc
cargo run --bin grpc-server
cargo run --bin grpc-client
```

---

`node-rpc-mosquitto`: implémentation microservice, d'une communication RPC et d'un broker MQTT

- utilisation de `@grpc-js`/`mqtt`/, expression du protobuf `proto/transformer.proto`

### Initialisation


```sh
cd node-rpc-mosquitto
npm install # (pnpm install)
```

### Lancement

```sh
cd node-rpc-mosquitto
npm start # (pnpm start) # lancera les deux services en simultané
```

Appel du service RPC via `mosquitto`

```sh
mosquitto_pub -t transformer-uppercase -m "test de la transformation en majuscule"
```

## Mercredi 18/01

`node_red-multiservices`: implémentation multi-service avec `node-red`

### Initialisation

```sh
cd node_red-multiservices
npm install # (pnpm install)
```

### Lancement

```sh
cd node_red-multiservices
npm start # (pnpm start) # lancera les deux services en simultané
```

### Utilisation

- `http://localhost:1881/` : interface web de node-red server
- `http://localhost:1880/` : interface web de node-red service2

```sh
curl --request GET \
  --url 'http://localhost:1881/server?type=ok' # retourne une 200

curl --request GET \
  --url 'http://localhost:1881/server?type=ko' # retourne une 400

curl --request GET \
  --url 'http://localhost:1881/server?type=blabla' # retourne une 500

```

## Jeudi 19/01

`gokit-sum-svc1`: implémentation microservice, d'une somme avec `go-kit`

### Lancement

```sh
cd gokit-sum-svc1
go run .
```

---

`gokit-sum-svc2`: implémentation microservice, d'une somme avec `go-kit`, avec un logger

### Lancement

```sh
cd gokit-sum-svc1
go run .
```

### Utilisation

```sh
curl --request GET \
  --url http://localhost:8080/sum \
  --header 'Content-Type: application/json' \
  --data '{"num1": "1","num2": "3"}'
```
