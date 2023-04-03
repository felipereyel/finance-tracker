# Finance Tracker

pocket base as backend and vue as frontend

## Run

You can specify a pesistant volume to store the data (`./pb_data`)

```bash
docker run -p 8080:8080 -v ./pb_data:/pb_data --name vue-pocket felipereyel/finance-tracker:latest
```

## Run locally (dev mode)

```bash
docker compose up --build
```

### Faster run local

```bash
docker compose up --build --scale vuewatcher=0
VITE_POCKETBASE_URL=http://localhost:8080 npm run dev:local
```

## TODO:

- [ ] Add tests
- [ ] Data Migrations
- [ ] Add modeling for extra purchases for same asset
