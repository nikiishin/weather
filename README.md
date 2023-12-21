# Проект Погодного Сервера

Этот проект представляет собой простой веб-сервер на Go, который предоставляет информацию о погоде для заданного города.

## Основные функции

- **/hello**: Возвращает приветственное сообщение.
- **/weather/{city}**: Возвращает данные о погоде для указанного города.

## Установка и запуск

1. Клонируйте репозиторий в вашу локальную машину.
2. Установите необходимые зависимости, используя `go get`.
3. Запустите сервер, используя `go run main.go`.

## Конфигурация

Для работы с API OpenWeatherMap необходимо создать файл `.apiConfig` в корневом каталоге проекта. Файл должен содержать следующую структуру:

```json
{
    "openWeatherMapApiKey": "Ваш API-ключ"
}
