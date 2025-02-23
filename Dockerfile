FROM ubuntu:latest

# Устанавливаем нужные пакеты и добавляем поддержку Go
RUN apt-get update && \
    apt-get install -y curl uuid-runtime wget unzip software-properties-common && \
    add-apt-repository -y ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install -y golang-go && \
    rm -rf /var/lib/apt/lists/*

# Загружаем и устанавливаем Xray
RUN mkdir -p /xray && \
    curl -L -o /xray/Xray-linux.zip https://github.com/XTLS/Xray-core/releases/latest/download/Xray-linux-64.zip && \
    unzip /xray/Xray-linux.zip -d /xray && \
    chmod +x /xray/xray && \
    mv /xray/xray /usr/local/bin/xray || (ls -l /xray && exit 1)

# Создаём рабочую директорию
WORKDIR /app

# Копируем весь проект
COPY . .

# Настраиваем Go-зависимости
RUN rm -f go.mod go.sum && \
    go mod init vpn-node && \
    go get github.com/gin-gonic/gin && \
    find . -type f -name "*.go" -exec sed -i 's|github.com/ByteForge-Systems/vpn-node/api/routes|vpn-node/api/routes|g' {} + && \
    find . -type f -name "*.go" -exec sed -i 's|github.com/ByteForge-Systems/vpn-node/api/handlers|vpn-node/api/handlers|g' {} + && \
    find . -type f -name "*.go" -exec sed -i 's|github.com/ByteForge-Systems/vpn-node/utils|vpn-node/utils|g' {} + && \
    find . -type f -name "*.go" -exec sed -i 's|github.com/ByteForge-Systems/vpn-node/scripts|vpn-node/scripts|g' {} + && \
    go mod tidy

# Копируем конфиг Xray
COPY xray/config.json /usr/local/etc/xray/config.json

# Копируем скрипт запуска
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Открываем нужные порты
EXPOSE 443 3000 8080

# Запуск контейнера
CMD ["/entrypoint.sh"]
