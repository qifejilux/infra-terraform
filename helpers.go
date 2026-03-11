
type ServerConfig struct {
    Port     int    `json:"port"`
    Host     string `json:"host"`
    LogLevel string `json:"log_level"`
}


func ParseDate(dateStr string) (time.Time, error) {
    layout := "2006-01-02T15:04:05Z"
    return time.Parse(layout, dateStr)
}

