FROM ubuntu:latest

# Обновляем apt и устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    ca-certificates \
    build-essential \
    nano \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Go (версия 1.22.5, можно заменить на нужную)
ENV GO_VERSION=1.22.5
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Устанавливаем Xray из официального репозитория
RUN bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install

# Копируем исходный код приложения
WORKDIR /app
COPY api /app/api

# Копируем скрипт запуска
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Пробрасываем необходимые порты: 8080 для Go-приложения, 445 для Xray
EXPOSE 8080 445

# Запускаем скрипт entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
