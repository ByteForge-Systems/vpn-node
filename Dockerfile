FROM ubuntu:latest

# Обновляем apt и устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    ca-certificates \
    build-essential \
    nano \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Go (версия 1.22.5)
ENV GO_VERSION=1.22.5
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Устанавливаем Xray вручную (скачиваем архив и распаковываем бинарник)
ENV XRAY_VERSION=1.9.2
RUN wget https://github.com/XTLS/Xray-core/releases/download/v${XRAY_VERSION}/Xray-linux-64.zip && \
    unzip Xray-linux-64.zip -d /usr/local/bin/ && \
    chmod +x /usr/local/bin/xray && \
    rm Xray-linux-64.zip

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
