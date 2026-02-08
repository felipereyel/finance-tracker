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

## Application Screens

### Assets List Screen
- **Template**: `AssetSummaryPage`
- **Route**: `/assets`
- **Features**:
  - Displays all assets in a sortable table
  - Filters by asset type and wallet
  - Shows name, wallet, type, buy date, buy price, last date, and last price
  - Navigation to summary page and create new asset modal

### Summary Dashboard
- **Template**: `SummaryPage`
- **Route**: `/summary`
- **Features**:
  - Displays financial summary charts and analytics
  - Embedded iframe for comprehensive chart visualization
  - Overview of portfolio performance

### Asset Details Screen
- **Template**: `AssetDetailsPage`
- **Route**: `/assets/{id}`
- **Features**:
  - Detailed view of individual asset
  - Price history chart visualization
  - Editable fields for sell date and comments (auto-save)
  - Read-only display of asset information
  - Links to price history and add new price (if not sold)

### Asset Prices History Screen
- **Template**: `AssetPricesPage`
- **Route**: `/assets/{id}/prices`
- **Features**:
  - Complete price history for an asset
  - Table showing date, value, and comments
  - Link to individual price details
  - Option to add new price entries

### New Asset Modal
- **Template**: `NewAsset`
- **Features**:
  - Modal form for creating new assets
  - Fields: name, wallet selection, asset type, initial price, buy date, comment
  - Form validation and submission

### New Price Modal
- **Template**: `NewPrice`
- **Features**:
  - Modal form for adding price updates to existing assets
  - Fields: price, log date, comment
  - Pre-fills with current asset price

### Price Details Screen
- **Template**: `PricePage`
- **Route**: `/prices/{id}`
- **Features**:
  - Detailed view of individual price entry
  - Shows value, gain calculation, and logged date
  - Editable comment field (auto-save)
  - Read-only display of price information

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
- [ ] refactor withControllerClousure

### Later

- [ ] Add tests
- [ ] Add modeling for extra purchases for same asset
- [ ] currency
- [ ] % gain
