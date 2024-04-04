# Protos (AI Marketplace)

***тестовое задание***

Здесь хранится описание proto-контракта, а также актуальный скомпилированный gRPC Server & Client

Полный проект: [AI Marketplace](https://github.com/shamank/ai-marketplace)

### Использованные библиотеки:
- **protoc** - генерация pb-файлов по файлам `.proto`


### Запуск генерации

**Для старта микросервиса**:

```sh
make gen
```
или
```sh
protoc stats-service\stats-service.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
```

### Структура проекта

```
├───gen // сгенерированные proto-файлы
│   └───go
│       └───stats-service
└───stats-service // а здесь лежат сами .proto-файлы

```

*Примечание*: выделил proto-контракты и скомпилированные файлы в отдельный репозиторий для более удобной работы
