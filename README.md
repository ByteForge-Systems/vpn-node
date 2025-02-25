Управление пользователями:
1. POST    /api/key/ - Добавить нового пользователя
2. DELETE  /api/key/{id} - Удалить пользователя по ID
3. GET     /api/key/{id}/link	- Сгенерировать ссылку для пользователя
4. GET	   /api/key/ - Получить список всех пользователей

Управление Xray:
1. POST /api/management/restart - Перезапустить Xray
2. GET	/api/management/status - Получить текущий статус Xray
3. POST /api/management/start - Запустить Xray
4. POST /api/management/stop - Остановить Xray