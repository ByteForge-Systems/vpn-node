Управление пользователями:
POST /api/user/ — добавить нового пользователя.
DELETE /api/user/{id} — удалить пользователя.
GET /api/user/{id}/link — сгенерировать VLESS-ссылку.

Управление Xray:
POST /api/management/restart — перезапустить Xray.
GET /api/management/status — получить статус Xray.
POST /api/management/start — запустить Xray.
POST /api/management/stop — остановить Xray.