сервисы

api-service - HTTP API (создание заказа + просмотр статистики)
 worker-service - Kafka consumer- обновляет Redis

Инфраструктура 
Kafka + Zookeeper
Kafka UI
Redis
RedisInsight



слои
domain - сущности
usecase — бизнес-логика + интерфейсы
integration - внешние зависимости 
transport (api-service) - HTTP handlers
consumer (worker-service) - Kafka consumer loop

Запуск инфраструктуры 

start_all.bat
