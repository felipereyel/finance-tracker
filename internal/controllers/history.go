package controllers

import (
	"fintracker/internal/models"
	"fintracker/internal/repositories/database"
	"sort"
	"strings"
)

type historyController struct {
	db database.Database
}

func (h historyController) GetHistory(userId string, aggregation string) (models.History, error) {
	histories, err := h.db.GetHistoricalPrices(userId)
	if err != nil {
		return models.EmptyHistory, err
	}

	wallets, err := h.db.ListWallets(userId)
	if err != nil {
		return models.EmptyHistory, err
	}

	history := models.History{
		Aggregation: aggregation,
		Series:      make([]models.HistorySeries, 0),
		Wallets:     make([][]string, 0),
	}

	for _, wallet := range wallets {
		history.Wallets = append(history.Wallets, []string{wallet.Id, wallet.Name})
	}

	allDates := make(map[string]bool)
	for _, assetHistory := range histories {
		for _, price := range assetHistory.Prices {
			allDates[price.LoggedAt] = true
		}
	}

	var sortedDates []string
	for date := range allDates {
		sortedDates = append(sortedDates, date)
	}
	sort.Strings(sortedDates)

	switch aggregation {
	case "total":
		series := h.buildTotalSeries(sortedDates, histories)
		history.Series = append(history.Series, series)
	case "wallet":
		series := h.buildWalletSeries(sortedDates, histories)
		history.Series = series
	case "tag":
		series := h.buildTagsSeries(sortedDates, histories)
		history.Series = series
	default:
		series := h.buildTotalSeries(sortedDates, histories)
		history.Series = append(history.Series, series)
	}

	return history, nil
}

func (h historyController) buildTotalSeries(dates []string, histories []models.AssetPriceHistory) models.HistorySeries {
	series := models.HistorySeries{
		Name:   "Total",
		Points: make([]models.HistoryDataPoint, 0),
	}

	for _, date := range dates {
		totalValue := float32(0)

		for _, assetHistory := range histories {
			latestPrice := h.getLatestPriceAtDate(assetHistory, date)
			if latestPrice != nil {
				totalValue += latestPrice.Value
			}
		}

		series.Points = append(series.Points, models.HistoryDataPoint{
			Date:  date,
			Value: totalValue,
		})
	}

	return series
}

func (h historyController) buildWalletSeries(dates []string, histories []models.AssetPriceHistory) []models.HistorySeries {
	walletHistories := make(map[string][]models.AssetPriceHistory)
	walletNames := make(map[string]string)

	for _, hist := range histories {
		walletHistories[hist.WalletId] = append(walletHistories[hist.WalletId], hist)
		walletNames[hist.WalletId] = hist.WalletName
	}

	var seriesList []models.HistorySeries

	for walletId, walletAssets := range walletHistories {
		series := models.HistorySeries{
			Name:   walletNames[walletId],
			Points: make([]models.HistoryDataPoint, 0),
		}

		for _, date := range dates {
			totalValue := float32(0)

			for _, assetHistory := range walletAssets {
				latestPrice := h.getLatestPriceAtDate(assetHistory, date)
				if latestPrice != nil {
					totalValue += latestPrice.Value
				}
			}

			series.Points = append(series.Points, models.HistoryDataPoint{
				Date:  date,
				Value: totalValue,
			})
		}

		seriesList = append(seriesList, series)
	}

	return seriesList
}

func (h historyController) buildTagsSeries(dates []string, histories []models.AssetPriceHistory) []models.HistorySeries {
	tagHistories := make(map[string][]models.AssetPriceHistory)

	for _, hist := range histories {
		tags := models.ParseTags(hist.Tag)
		for _, tag := range tags {
			tagHistories[tag] = append(tagHistories[tag], hist)
		}
	}

	var seriesList []models.HistorySeries

	for tagName, tagAssets := range tagHistories {
		series := models.HistorySeries{
			Name:   tagName,
			Points: make([]models.HistoryDataPoint, 0),
		}

		for _, date := range dates {
			totalValue := float32(0)

			for _, assetHistory := range tagAssets {
				latestPrice := h.getLatestPriceAtDate(assetHistory, date)
				if latestPrice != nil {
					totalValue += latestPrice.Value
				}
			}

			series.Points = append(series.Points, models.HistoryDataPoint{
				Date:  date,
				Value: totalValue,
			})
		}

		seriesList = append(seriesList, series)
	}

	// Sort by total value at the latest date (highest first)
	if len(dates) > 0 {
		lastDate := dates[len(dates)-1]
		sort.Slice(seriesList, func(i, j int) bool {
			var valI, valJ float32
			for _, p := range seriesList[i].Points {
				if p.Date == lastDate {
					valI = p.Value
					break
				}
			}
			for _, p := range seriesList[j].Points {
				if p.Date == lastDate {
					valJ = p.Value
					break
				}
			}
			return valI > valJ
		})
	}

	return seriesList
}

func (h historyController) getLatestPriceAtDate(assetHistory models.AssetPriceHistory, targetDate string) *models.Price {
	prices := assetHistory.Prices
	sellDate := assetHistory.SellDate

	// Truncate targetDate for comparison (take first 10 chars, trim space)
	targetDateClean := strings.TrimSpace(targetDate[:min(10, len(targetDate))])

	// Prepare sellDate for comparison
	sellDateClean := ""
	if sellDate != "" {
		sellDateClean = strings.TrimSpace(sellDate[:min(10, len(sellDate))])
	}

	// If asset was sold and target date is after sell date, return nil (not owned)
	if sellDateClean != "" && targetDateClean > sellDateClean {
		return nil
	}

	var latest *models.Price

	for i := range prices {
		priceDate := strings.TrimSpace(prices[i].LoggedAt[:min(10, len(prices[i].LoggedAt))])

		// Only consider prices up to the target date
		if priceDate <= targetDateClean {
			if latest == nil || prices[i].LoggedAt > latest.LoggedAt {
				latest = &prices[i]
			}
		}
	}

	return latest
}
