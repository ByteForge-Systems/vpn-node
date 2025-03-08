package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ByteForge-Systems/vpn-node/utils"
	"github.com/google/uuid"
)

// Логика для управления Xray

type Client struct {
	ID   string `json:"id"`
	Flow string `json:"flow"`
}

type Config struct {
	Inbounds []struct {
		Settings struct {
			Clients []Client `json:"clients"`
		} `json:"settings"`
	} `json:"inbounds"`
}

// Загрузка конфигурации Xray
func loadConfig() (*Config, error) {
	data, err := os.ReadFile(utils.GetEnv("CONFIG_PATH"))
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// Сохранение конфигурации Xray
func saveConfig(clients []Client) error {
	data, err := os.ReadFile(utils.GetEnv("CONFIG_PATH"))
	if err != nil {
		return err
	}
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	if inbounds, ok := raw["inbounds"].([]interface{}); ok && len(inbounds) > 0 {
		if inbound, ok := inbounds[0].(map[string]interface{}); ok {
			if settings, ok := inbound["settings"].(map[string]interface{}); ok {
				settings["clients"] = clients
			}
		}
	}
	newData, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(utils.GetEnv("CONFIG_PATH"), newData, 0644)
}

// Запуск Xray
func StartXray() error {
	cmd := exec.Command("xray", "-config", utils.GetEnv("CONFIG_PATH"))
	return cmd.Start()
}

// Остановка Xray
func StopXray() error {
	cmd := exec.Command("pkill", "-f", "xray")
	return cmd.Run()
}

// Перезапуск Xray
func RestartXray() error {
	if err := StopXray(); err != nil {
		return err
	}
	return StartXray()
}

// Генерация нового пользователя
func GenerateUser() (string, error) {
	newUUID := uuid.New().String()
	config, err := loadConfig()
	if err != nil {
		return "", err
	}
	newClient := Client{
		ID:   newUUID,
		Flow: "xtls-rprx-vision",
	}
	config.Inbounds[0].Settings.Clients = append(config.Inbounds[0].Settings.Clients, newClient)
	err = saveConfig(config.Inbounds[0].Settings.Clients)
	if err != nil {
		return "", err
	}
	err = RestartXray()
	if err != nil {
		return "", err
	}
	return newUUID, nil
}

// Удаление пользователя
func RemoveUser(userID string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}
	clients := config.Inbounds[0].Settings.Clients
	newClients := []Client{}
	for _, client := range clients {
		if client.ID != userID {
			newClients = append(newClients, client)
		}
	}
	if len(newClients) == len(clients) {
		return fmt.Errorf("❌ Пользователь не найден")
	}
	err = saveConfig(newClients)
	if err != nil {
		return err
	}
	err = RestartXray()
	if err != nil {
		return err
	}
	return nil
}

// Генерация VLESS-ссылки
func GenerateVLESSLink(userID string) (string, error) {
	serverIP, err := exec.Command("curl", "-s", "ifconfig.me").Output()
	if err != nil {
		return "", err
	}
	publicKey := utils.GetEnv("PUBLIC_KEY")
	vlessLink := fmt.Sprintf("vless://%s@%s:445?security=reality&encryption=none&pbk=%s&fp=chrome&type=tcp&flow=xtls-rprx-vision-udp443&sni=www.cloudflare.com#XrayVPN",
		userID, strings.TrimSpace(string(serverIP)), publicKey)
	return vlessLink, nil
}

// Список всех пользователей
func ListUsers() ([]Client, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return config.Inbounds[0].Settings.Clients, nil
}

// Проверка статуса Xray
func GetXrayStatus() (string, error) {
	cmd := exec.Command("pgrep", "-f", "xray")
	output, err := cmd.Output()
	if err != nil || len(output) == 0 {
		return "Xray is not running", nil
	}
	return "Xray is running", nil
}

// Сбор метрик сервера
func GetServerMetrics() (string, error) {
	return "Server metrics", nil
}
