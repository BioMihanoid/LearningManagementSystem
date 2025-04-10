# Learning Management System API (LMS API)

## Описание
**LMS API (система управления обучением) — это REST API, разработанное на Go для управления учебными курсами, ролями пользователей, материалами, тестированиями и логированием активности.**

## Запуск
### Запускать только после пункта Docker
```sh
go run cmd/main.go
```
### Docker
```sh
docker-compose up -d
```

## Функции
- 🔑 **Аутентификация и авторизация** (включая JWT, refresh tokens)
- 🎓 **Курсы и учебные материалы**
- ✏️ **Тестирование студентов**
- 📒 **Логирование активности**

## Основные API-маршруты

### Аутентификация и пользователи
- `POST /auth/register` — Регистрация пользователя
- `POST /auth/login` — Вход в систему
- `GET /auth/profile` — Получение профиля
- `PUT /auth/profile` — Обновление профиля

### Курсы и уроки
- `GET /auth/courses` — Получение списка курсов
- `GET /auth/courses/:id` — Получение информации о курсе
- `GET /auth/courses/:id/lessons` — Получение уроков курса
- `GET /auth/lessons/:id` — Получение информации об уроке

### Тестирование
- `GET /auth/test/:id` — Получение теста
- `POST /auth/test/id/submit` — Отправка результатов теста

### Управление курсами и уроками (Преподаватель)
- `POST /auth/teacher/courses` — Создание курса
- `PUT /auth/teacher/courses/:id` — Обновление курса
- `DELETE /auth/teacher/courses/:id` — Удаление курса
- `POST /auth/teacher/lessons` — Создание урока
- `PUT /auth/teacher/lessons/:id` — Обновление урока
- `DELETE /auth/teacher/lessons/:id` — Удаление урока

### Управление пользователями (Админ)
- `GET /auth/admin/users` — Получение списка пользователей
- `PUT /auth/admin/users/:user_id` — Изменение роли пользователя  

## Переменные окружения
- `CONFIG_PATH` — Строка c путем до файла `config.yaml`
- `SECRET_KEY` — Секретный ключ JWT 

## Автор
[BioMihanoid](https://github.com/BioMihanoid)