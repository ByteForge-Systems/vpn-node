Управление пользователями:
1. POST    /api/user/ - Добавить нового пользователя
2. DELETE  /api/user/{id} - Удалить пользователя по ID
3. GET     /api/user/{id}/link	- Сгенерировать ссылку для пользователя
4. GET	   /api/user/all - Получить список всех пользователей

Управление Xray:
1. POST /api/management/restart - Перезапустить Xray
2. GET	/api/management/status - Получить текущий статус Xray
3. POST /api/management/start - Запустить Xray
4. POST /api/management/stop - Остановить Xray