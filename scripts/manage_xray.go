package scripts

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"os"
	"strings"
	"github.com/google/uuid"
)

// Логика для управления Xray

const (
	configPath = "/usr/local/etc/xray/config.json"
)

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
	data, err := os.ReadFile(configPath)
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
	data, err := os.ReadFile(configPath)
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
	return os.WriteFile(configPath, newData, 0644)
}

// Перезапуск Xray
func RestartXray() error {
	cmd := exec.Command("systemctl", "restart", "xray")
	return cmd.Run()
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
	publicKey := "Ivxj3onvH96bRecsbASk6MThhM0nQXinPQFamD3eeGM" // вынести?
	vlessLink := fmt.Sprintf("vless://%s@%s:443?security=reality&encryption=none&pbk=%s&fp=chrome&type=tcp&flow=xtls-rprx-vision-udp443&sni=www.cloudflare.com#XrayVPN",
		userID, strings.TrimSpace(string(serverIP)), publicKey) // ip сервера в константу?
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
	cmd := exec.Command("systemctl", "status", "xray")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	status := strings.Split(string(output), "\n")
	for _, line := range status {
		if strings.Contains(line, "Active:") {
			return line, nil
		}
	}
	return "", nil
}

// Сбор метрик сервера
func GetServerMetrics() (string, error) {
	// Реализация сбора метрик
	return "Server metrics", nil
}