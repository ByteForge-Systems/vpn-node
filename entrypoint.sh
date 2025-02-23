#!/bin/bash

if [ ! -f "/app/api/.env" ]; then
    echo "Генерируем ключи..."
    KEYS=$(xray x25519)
    PRIVATE_KEY=$(echo "$KEYS" | grep "Private key" | awk '{print $NF}')
    PUBLIC_KEY=$(echo "$KEYS" | grep "Public key" | awk '{print $NF}')

    echo "XRAY_PRIVATE_KEY=$PRIVATE_KEY" > /app/api/.env
    echo "XRAY_PUBLIC_KEY=$PUBLIC_KEY" >> /app/api/.env

    echo "Ключи сохранены в .env"
else
    echo ".env уже существует, пропускаем генерацию ключей."
    PRIVATE_KEY=$(grep XRAY_PRIVATE_KEY /app/api/.env | cut -d '=' -f2)
    PUBLIC_KEY=$(grep XRAY_PUBLIC_KEY /app/api/.env | cut -d '=' -f2)
fi

sed -i "s/REPLACE_WITH_PRIVATE_KEY/$PRIVATE_KEY/g" /usr/local/etc/xray/config.json
sed -i "s/REPLACE_WITH_PUBLIC_KEY/$PUBLIC_KEY/g" /usr/local/etc/xray/config.json

echo "Запускаем Xray..."
/usr/local/bin/xray -config /usr/local/etc/xray/config.json &

echo "Запускаем vpn-node..."
cd /app/api && go run main.go
