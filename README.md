# ShortenerURL

Сервис для сокращения длинных URL-адресов, написанный на Go. Позволяет создавать короткие ссылки и перенаправлять по ним.

## Возможности

- 🔗 Сокращение длинных URL-адресов
- 🌐 REST API для интеграции
- 💾 Хранение данных в Redis
- 🐳 Docker контейнеризация
- 🔍 Поиск и управление ссылками

## Технологии

- **Бэкенд**: Go
- **Кэш**: Redis
- **Контейнеризация**: Docker, Docker Compose

## Настройка переменных окружения
Создайте в api/ файл .env, содержащий:
```env
API_QUOTA=10 - ограничние на количество запросов в 30 минут
DOMAIN="localhost:3000"
DB_ADDR="db:6379"
DB_PASS=""
APP_PORT=":3000"
```

## Установка и запуск

```bash
# Клонировать репозиторий
git clone https://github.com/TemniyKozhevnik/ShortenerURL.git
cd ShortenerURL

# Запустить все сервисы
docker-compose up -d
```

## API Endpoints

### Основные endpoints

#### 🔗 Создание короткой ссылки
```http
POST /api/v1/
```

#### Пример запроса
Запрос
```json
{
    "url": "https://github.com/TemniyKozhevnik"
}
```

Ответ
```json
{
    "url": "https://github.com/TemniyKozhevnik",
    "short": "localhost:3000/106214",
    "expiry": 24,
    "rate_limit": 7,
    "rate_limit_reset": 18
}
```

#### 🔗 Переход по короткой ссылке
```http
GET /<short>
```

#### Пример запроса
Запрос
http://localhost:3000/106214


Ответ
https://github.com/TemniyKozhevnik
