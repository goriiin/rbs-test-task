
Сборка и запуск в корне репозитория выполнить команду `sudo make up`. 

Команда создаёт `.env`, если его нет, данные можно поменять. Собирает image `back_go`
поднимает `back_go`, `db_pg`, `nginx`.
Переменные БД берутся из`docker/.env`.

Проверка работы:
- curl 
  - `curl -s http://localhost:80/ping` → `{"status":"PONG"}`; 
  - `curl -s http://localhost:80/list` возвращает HTML‑таблицу с данными.
  - `curl http://localhost:80/health` - возвращает html с текстом healthy 
  - `curl -v --noproxy '*' http://localhost:80/add \
     -d "city=Moscow&temperature=15"` - добавляет запись в бд

- Браузер 
  - http://localhost:80/list - напрямую показывает ту же таблицу.
  - http://localhost:80/add - содержит форму для добавления
  - http://localhost:80/ping - > `{"status":"PONG"}`
  - http://localhost:80/health -> healthy


Проверка подключения к Postgres: контейнер `db_pg`. 
Заходим в контейнер в интерактивном режиме, вставляем запись и обновляем страницу http://localhost:80/list, должна появится запись
`sudo docker exec -it db_pg psql -U $MAIN_DB_USER -d $MAIN_DB_NAME -c "INSERT INTO weather (city, temperature) VALUES ('Kazan',14);"`

Если повторный запрос `/list` отражает новую строку, значит есть соединение между бекендом и бд.

Для очистки после теста выполнить 

`sudo make prune`

