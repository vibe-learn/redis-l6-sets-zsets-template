        # redis — Множества и отсортированные множества

        Homework-шаблон для урока **l3_sets_and_zsets** (Множества и отсортированные множества) на платформе Vibe Learn.

        ## Что делать

        Реализуй leaderboard service: HTTP API с эндпойнтами PUT /score (обновить), GET /top?n=10,
GET /rank/:user. Внутри — zset. Тесты проверят: concurrent updates не теряются (1000 параллельных
ZINCRBY → сумма верна), top-N корректен, rank возвращает правильную позицию.

## Контекст (из transfer-задачи урока)

Социальная сеть. Нужно: (1) показать «друзья друзей моих друзей» для рекомендаций;
(2) trending hashtags за последний час; (3) топ-100 постов по лайкам. У тебя есть Redis.
Опиши какие типы и команды используешь для каждой задачи, и какие есть trade-off'ы по
памяти / времени.

## Recap из урока

- Set — уникальные элементы без порядка. SISMEMBER O(1). SINTER/SUNION/SDIFF для теории множеств.
- Sorted set — элементы со score, отсортированы. Skip list + hash внутри → ZADD/ZRANGE O(log N).
- Leaderboard в Redis: ZADD для апдейта, ZREVRANGE 0 9 WITHSCORES для топа, ZRANK для позиции. Всё O(log N), микросекунды.
- Sliding-window rate-limit через zset с timestamp как score — точнее чем INCR fixed-window, но дороже по памяти.
- SINTER на больших множествах — O(N×M), single-thread → может заморозить. Precompute через SINTERSTORE раз в N минут.

        ## Как работать

        1. Платформа Vibe Learn создаёт копию этого репо в твоём GitHub-аккаунте по клику «Начать домашку» на странице урока (через GitHub `/generate`, codecrafters-pattern).
        2. Склонируй копию локально, реализуй TODO в `main.go`, прогони тесты, запушь.
        3. CI (`.github/workflows/ci.yml`) запускает `go vet` + `go test ./...` на каждый push. Платформа слушает результат через webhook от GitHub Actions и обновляет статус домашки на странице урока.

        ## Локальное окружение

        - Go 1.22+
        - Docker + docker-compose — `docker compose up -d` поднимает single-node Redis 7 на `localhost:6379` (с включёнными keyspace-notifications и AOF). Адрес переопределяется через env `REDIS_ADDR`.

        ## Запуск

        ```bash
        # Поднять локальный Redis
        docker compose up -d

        # Прогнать тесты (интеграционный включается через REDIS_INTEGRATION=1)
        go test ./...
        REDIS_INTEGRATION=1 go test ./...

        # Запустить main (печатает marker; замени stub на реализацию)
        go run .
        ```

        ## Заметка автора

        Это baseline-шаблон, сгенерированный платформой. Бизнес-сущность задачи (что конкретно реализовать в `main.go`, какие тесты сделать строгими) расширяется по ходу итераций — параллельно с углублением теории урока.
