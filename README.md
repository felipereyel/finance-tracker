# Finance Tracker

go (pocket base) as backend and vue as frontend

## Build and Run

Build

```bash
docker build -t fintracker .
```


Run. You need to specify a pesistant volume to store the data (`./pb_data`)

```bash
docker run -p 8080:8080 -v ./pb_data:/pb_data fintracker
```

## Run locally (dev mode)

### Go Backend

```bash
go run main.go
# or
air
```

### Vue Frontend

```bash
VITE_POCKETBASE_URL=http://localhost:8090 npm run dev
```

## TODO:

### Must

- [x] Add hrefs
- [x] Add window titles
- [x] Data Migrations
- [x] wallets
- [ ] Users
- [ ] User add wallet
- [ ] Add template data
- [ ] Fix Date Inputs

### Later

- [ ] add other asset type not already on wallet
- [ ] Back
- [ ] Add tests
- [ ] Add modeling for extra purchases for same asset
- [ ] currency
- [ ] % gain
- [ ] extrapolation 