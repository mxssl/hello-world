# hello-world

## Роуты

`/` - отвечает кодом 200 и строкой `Hello World!`

`/health-check` - отвечает кодом 200

`/ready-check` - отвечает кодом 200

`/metrics` - отвечаем кодом 200, показывает метрики в формате prometheus. Кастомная метрика всего одна - `http_total_requests` - показывает общее кол-во запросов на `/` роут.

## Запуск локально

### Зависимости

В системе должны быть установлены:

- [docker](https://docs.docker.com/install)
- [docker-compose](https://docs.docker.com/compose/install)

### Запускаем docker-compose

```bash
make up
```

Или в detach режиме

```bash
docker-compose up -d
```

Далее приложение доступно на локальном порту `8080`

```bash
curl localhost:8080
Hello World!
```

Prometheus доступен на локальном порту `9090`:

[prometheus](http://localhost:9090)

Grafana доступна на локальном порту `3000`:

[grafana](http://localhost:3000)

```text
L: admin
P: admin
```

В grafana доступен дашборд App с несколькими графиками по метрикам приложения (процессор, память, горутины, общее кол-во http запросов).

## Kubernetes

Запуск приложения в Kubernetes

```bash
make kubernetes
```

Посмотреть Kubernetes манифесты можно в директории `kube`

## Запуск тестов

### Запуск юнит тестов приложения в golang контейнерe

```bash
make docker-test
```

### Запуск тестов контейнерезации

Зависимости:

- [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss)

Запуск `dgoss` тестов

```bash
make dgoss
```

Ознакомиться с конфигурацией `dgoss` можно в файле `goss.yaml`
