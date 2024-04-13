# Сервис orgchart

## Предназначение
Просмотр и редактирование данных об оргструктуре компании «Дороги и ямы».

## Автор
Константин Калинин

## Запуск сервиса/тестов

Для запуска сервиса нужно выполнить команды:
```bash
bash bin/run-codegen
brewkit build
docker-compose up --build -d
docker-compose down
```

Для запуска end2end тестов:
```bash
brewkit build
bin/run-end2end-tests
```