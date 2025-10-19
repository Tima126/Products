
# Products API (CRUD для телефонов)

**Описание:**
Простое Go-приложение с PostgreSQL для управления таблицей **продуктов телефонов**.
Реализован **CRUD**: создание, чтение, обновление и удаление продуктов.

---

* **app/** — исходный код Go приложения
* **templates/** — HTML-шаблоны страниц
* **db/** — работа с PostgreSQL
* **migrations/** — SQL миграции для создания таблиц и начальных данных

---

## Порядок запуска

### 1️ Остановить и очистить контейнеры (если проект уже запускался)

```bash
docker compose down -v
```

* `-v` — удаляет все тома с данными, чтобы начать с чистой БД

---

### 2️ Запуск базы данных

```bash
docker compose up -d db
```

* Контейнер `db` запускает PostgreSQL
* Проверка логов:

```bash
docker logs product_postgres
```

* Должно быть:

```
database system is ready to accept connections
```

---

### 3️ Применение миграций

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/product_db?sslmode=disable" up
```

* `-path /migrations` — путь к папке с миграциями
* `-database` — строка подключения к PostgreSQL
* `up` — применяет все миграции (создаёт таблицы и добавляет начальные продукты)

**Примеры отката миграций:**

* Последняя миграция:

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/product_db?sslmode=disable" down 1
```

* Все миграции:

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/product_db?sslmode=disable" down
```

* Полное удаление всех таблиц:

```bash
docker-compose run migrate -path /migrations -database "postgres://admin:12345@db:5432/product_db?sslmode=disable" drop -f
```

---

### 4️ Сборка и запуск Go-сервера

```bash
docker compose build app
docker compose up -d app
```

* Сервис `app` запускает Go-сервер на порту **8080**
* Проверка логов:

```bash
docker logs product_app
```

* Должно быть:

```
Connection pool created successfully
Server started on port 8080
```

---

### 5️ Доступ к API

* **HTML страница со всеми продуктами:**

```
http://localhost:8080/
```

* **JSON API:**

```
GET /products          — список всех продуктов
GET /products/{id}     — продукт по ID
POST /products         — добавить продукт
PUT /products/{id}     — обновить продукт
DELETE /products/{id}  — удалить продукт
```

---

### 6️ Примеры запросов через curl

#### Добавить продукт

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"iPhone 15","description":"Новый iPhone","price":1200.0}'
```

#### Получить все продукты

```bash
curl http://localhost:8080/products
```

#### Получить продукт по ID

```bash
curl http://localhost:8080/products/1
```

#### Обновить продукт

```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"name":"iPhone 15 Pro","description":"Обновлённый","price":1400.0}'
```

#### Удалить продукт

```bash
curl -X DELETE http://localhost:8080/products/1
```

---

### 7️ Примечания

* HTML-шаблоны лежат в `templates/` и рендерятся для страницы `/`
* JSON API доступно через `/products` и `/products/{id}`
* Все данные хранятся в PostgreSQL (`db` контейнер)


