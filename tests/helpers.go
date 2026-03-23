package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// GetServicesPath returns the path to the data/services directory.
func GetServicesPath() string {
	_, filename, _, _ := runtime.Caller(1)
	base := filepath.Dir(filename)
	return filepath.Join(base, "..", "data", "services")
}

// GetServicesDir returns the path to the data/services directory.
func GetServicesDir() string {
	return filepath.Join(GetServicesPath(), "dir")
}

// GetTelegramToken returns the path to the telegram token file.
func GetTelegramToken() string {
	return filepath.Join(GetServicesPath(), "token.txt")
}

// GetTelegramChatID returns the path to the telegram chat ID file.
func GetTelegramChatID() string {
	return filepath.Join(GetServicesPath(), "chat_id.txt")
}

// GetTelegramBot returns the telegram bot instance
func GetTelegramBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	return bot, nil
}

// GetProxyList returns a list of proxy servers from the proxy.txt file.
func GetProxyList(filePath string) []string {
	proxies := make([]string, 0)
	if _, err := os.Stat(filePath); os.IsExist(err) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.HasPrefix(line, "#") {
				proxies = append(proxies, line)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return proxies
}