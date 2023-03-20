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
