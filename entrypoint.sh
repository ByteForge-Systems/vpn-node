#!/bin/sh
set -e

echo "Создание директории для конфигурации Xray..."
mkdir -p /usr/local/etc/xray

echo "Генерация конфигурационного файла Xray..."
cat <<EOF > /usr/local/etc/xray/config.json
{
  "inbounds": [
    {
      "port": 445,
      "protocol": "vless",
      "settings": {
        "clients": [],
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
          "privateKey": "${XRAY_PRIVATE_KEY}",
          "shortIds": [""]
        }
      }
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

echo "Запуск Xray..."
xray -config /usr/local/etc/xray/config.json &

echo "Ожидание запуска Xray..."
sleep 3

echo "Запуск Go-приложения..."
cd /app/api
exec go run main.go
