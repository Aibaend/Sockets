Необходимо реализовать сервер для Websocket-соединений для обмена сообщениями по каналам.

Процесс ожидаем следующий:
Клиент подключается. После подключения может отправлять команды на подключение к каналам.
Есть HTTP-endpoint для внутренних сервисов, по которым они могут публиковать сообщение в определенные каналы.

1) Для подключения к каналу пользователь отсылает следующий json:
{
  "type": "sub",
  "channel_id": "какой-то id канала"
}
2) Для отключения от канала отправляет:
{
  "type": "unsub",
  "channel_id": "какой-то id канала"
}

1) При поступлении запроса на endpoint:
{
  "type": "event",
  "channel_id": "какой-то id канала",
  "msg": {...} 
}
*в msg передается объект
2) Данное сообщение транслируется клиентам, которые подписаны на канал указанный в channel_id