#!/bin/bash

# Обновляем систему и ставим нужные пакеты
echo "Обновляем систему и устанавливаем зависимости..."
sudo apt update
sudo apt install uuid-runtime curl -y -qq

# Устанавливаем Go 1.23.1
echo "Устанавливаем Go 1.23.1..."
GO_VERSION="1.23.1"
if ! command -v go &> /dev/null || ! /usr/local/go/bin/go version | grep -q "$GO_VERSION"; then
    echo "Скачиваем и устанавливаем Go $GO_VERSION..."
    wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    echo "export PATH=/usr/local/go/bin:\$PATH" | sudo tee -a /etc/profile
    export PATH="/usr/local/go/bin:$PATH"
    rm go${GO_VERSION}.linux-amd64.tar.gz
fi
echo "Проверяем версию Go..."
/usr/local/go/bin/go version

# Устанавливаем Xray
echo "Устанавливаем Xray..."
bash -c "$(curl -k -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install
if [ $? -ne 0 ]; then
    echo "Ошибка установки Xray. Выход."
    exit 1
fi

# Проверяем, что xray доступен
if ! command -v xray &> /dev/null; then
    echo "Xray не установлен. Проверяем путь..."
    if [ -f "/usr/local/bin/xray" ]; then
        sudo ln -sf /usr/local/bin/xray /usr/bin/xray
    else
        echo "Бинарник Xray не найден. Установка провалилась."
        exit 1
    fi
fi

# Перезагружаем конфигурацию systemd
echo "Перезагружаем конфигурацию systemd..."
sudo systemctl daemon-reload

# Проверяем и обновляем сервисный файл Xray
XRAY_SERVICE="/etc/systemd/system/xray.service"
if [ -f "$XRAY_SERVICE" ]; then
    echo "Обновляем лимиты в $XRAY_SERVICE..."
    sudo sed -i '/\[Service\]/a LimitNPROC=10000\nLimitNOFILE=1000000' $XRAY_SERVICE
    sudo systemctl daemon-reload
fi

# Генерируем ключи x25519
echo "Генерируем private и public ключи..."
KEYS=$(xray x25519)
PRIVATE_KEY=$(echo "$KEYS" | grep "Private key" | awk '{print $3}')
PUBLIC_KEY=$(echo "$KEYS" | grep "Public key" | awk '{print $3}')
echo "Private key: $PRIVATE_KEY"
echo "Public key: $PUBLIC_KEY"

# Генерируем UUID
echo "Генерируем UUID..."
UUID=$(uuidgen)
echo "UUID: $UUID"

# Путь к конфигу
CONFIG_PATH="/usr/local/etc/xray/config.json"

# Создаем конфиг
echo "Создаем конфигурационный файл в $CONFIG_PATH..."
sudo mkdir -p /usr/local/etc/xray
sudo tee "$CONFIG_PATH" > /dev/null <<EOF
{
  "api": {
    "services": ["HandlerService", "StatsService"],
    "tag": "api"
  },
  "stats": {},
  "policy": {
    "levels": {
      "0": {
        "statsUserUplink": true,
        "statsUserDownlink": true
      }
    }
  },
  "inbounds": [
    {
      "port": 443,
      "protocol": "vless",
      "settings": {
        "clients": [
          {
            "id": "$UUID",
            "flow": "xtls-rprx-vision"
          }
        ],
        "decryption": "none"
      },
      "streamSettings": {
        "network": "tcp",
        "security": "reality",
        "realitySettings": {
          "show": false,
          "dest": "www.cloudflare.com:443",
          "xver": 0,
          "serverNames": ["www.cloudflare.com"],
          "privateKey": "$PRIVATE_KEY",
          "shortIds": [""]
        }
      }
    },
    {
      "port": 10085,
      "listen": "127.0.0.1",
      "protocol": "dokodemo-door",
      "settings": {
        "address": "127.0.0.1"
      },
      "tag": "api"
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom",
      "settings": {}
    }
  ]
}
EOF

# Компилируем и запускаем main.go из папки api
API_DIR="./api"
GO_FILE="$API_DIR/main.go"
BINARY_NAME="api"

# Создаем или обновляем .env файл в папке api
ENV_FILE="$API_DIR/.env"
echo "Записываем переменные в $ENV_FILE..."
if [ ! -d "$API_DIR" ]; then
    echo "Папка $API_DIR не существует, создаем..."
    mkdir -p "$API_DIR"
fi
cat << EOF > "$ENV_FILE"
CONFIG_PATH=$CONFIG_PATH
PUBLIC_KEY=$PUBLIC_KEY
PRIVATE_KEY=$PRIVATE_KEY
EOF
if [ $? -eq 0 ]; then
    echo "Файл $ENV_FILE успешно создан."
    sudo chmod 644 "$ENV_FILE"
else
    echo "Ошибка при создании $ENV_FILE."
    exit 1
fi

# Перезапускаем и включаем Xray
echo "Запускаем Xray..."
sudo systemctl enable xray.service
sudo systemctl restart xray.service
if [ $? -ne 0 ]; then
    echo "Не удалось запустить xray.service. Проверяйте логи: journalctl -u xray.service"
else
    echo "Xray успешно запущен."
fi

if [ -f "$GO_FILE" ]; then
    echo "Компилируем $GO_FILE..."
    cd "$API_DIR" || { echo "Не удалось перейти в $API_DIR"; exit 1; }
    if [ -f "go.mod" ]; then
        echo "Найден go.mod в $API_DIR, проверяем версию..."
        cat go.mod
        if grep -q "go 1.23.5" go.mod; then
            sed -i 's/go 1.23.5/go 1.23/' go.mod
            echo "Исправлена версия Go в go.mod на 1.23"
        fi
        /usr/local/go/bin/go mod tidy
    fi
    /usr/local/go/bin/go build -o "$BINARY_NAME" main.go
    if [ $? -eq 0 ]; then
        echo "Компиляция прошла успешно, запускаем $BINARY_NAME..."
        ./$BINARY_NAME &
        echo "API запущен в фоновом режиме."
    else
        echo "Ошибка при компиляции $GO_FILE. Вывод ошибок:"
        /usr/local/go/bin/go build -v main.go
        exit 1
    fi
else
    echo "Файл $GO_FILE не найден. Пропускаем компиляцию и запуск."
fi

# Выводим инфу для пользователя
echo "Установка завершена!"
echo "Твой UUID: $UUID"
echo "Твой Private Key: $PRIVATE_KEY"
echo "Твой Public Key: $PUBLIC_KEY"
echo "Конфиг сохранен в: $CONFIG_PATH"
echo "Переменные записаны в: $ENV_FILE"
echo "Сервис Xray запущен и добавлен в автозагрузку."