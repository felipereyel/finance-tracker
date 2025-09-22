package utils

func FormatDateTime(datetime string) string {
	if len(datetime) >= 10 {
		return datetime[:10]
	}
	return datetime
}
