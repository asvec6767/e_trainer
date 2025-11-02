01.11.2025 старт проекта

## Технологии

- **Backend**: Go 1.23.2 + Gin
- **Frontend**: Vue 3 + PrimeVue 4.0 + Vite
- **Роутинг**: Vue Router

## Быстрый старт

### Backend

```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

### Frontend (разработка)

```bash
cd frontend
npm install
npm run dev
```

Frontend будет доступен на `http://localhost:5173`

### Сборка фронтенда для продакшена

```bash
cd frontend
npm run build
```

После сборки статические файлы будут в папке `static/`, и Go сервер будет их раздавать.

## Структура проекта

```
e_trainer/
├── frontend/          # Vue 3 приложение
│   ├── src/           # Исходный код
│   ├── package.json   # Зависимости фронтенда
│   └── vite.config.js # Конфигурация Vite
├── templates/         # HTML шаблоны (legacy)
├── static/            # Собранные статические файлы фронтенда
├── main.go           # Go сервер
└── go.mod            # Go зависимости
```

## Задачи

✅ Первые задачи:
- ✅ Hello world
- ✅ Базовая страница backend
- ✅ Базовая страница html шаблон
- ✅ Базовая страница frontend (Vue 3 + PrimeVue)

Вторые задачи:
- Определиться с дизайном
- Составление ТЗ и расписывание задач
- Расписать технолгии подробно

Следующие задачи:
- Развернуть через Докер
- Настроить Джиру и Конфлюенс
