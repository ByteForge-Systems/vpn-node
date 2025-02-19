Управление пользователями:
1. POST /api/user/ — добавить нового пользователя.
2. DELETE /api/user/{id} — удалить пользователя.
3. GET /api/user/{id}/link — сгенерировать VLESS-ссылку.
4. GET /api/user/all/slice - вывод всех пользователей

Управление Xray:
1. POST /api/management/restart — перезапустить Xray.
2. GET /api/management/status — получить статус Xray.
3. POST /api/management/start — запустить Xray.
4. POST /api/management/stop — остановить Xray
