# Finance Tracker

go (pocket base) as engine with templ and htmx for frontend

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

### Docker compose

```sh
docker compose up
```

### Go CLI

Go 1.25 and up required

```bash
# Install
go mod download

# Make templ and statics
make

# Run direct
go run main.go serve --http=0.0.0.0:8090

# or run with air for hot reload
go tool air serve --http=0.0.0.0:8090
```

## TODO:

### Must

- [x] Add hrefs
- [x] Add window titles
- [x] Data Migrations
- [x] wallets
- [x] Users
- [x] Fix Date Inputs (defaults)
- [x] add auth scope to routes
- [ ] add more charts to account summary
- [ ] add error fallback

### Later

- [ ] Add tests
- [ ] Add modeling for extra purchases for same asset
- [ ] currency
- [ ] % gain
