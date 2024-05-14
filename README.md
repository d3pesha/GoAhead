# GoAhead

## Начало работы

Для запуска приложения вам потребуются Git, Docker. По умолчанию, бэкэнд работает на порту 8000, а Postgres - на порту 5432.

### Клонирование репозитория

```bash
git clone https://github.com/d3pesha/GoAhead
cd GoAhead
```

### Запуск сервиса

```bash
docker-compose up
```

Эта команда выполнит следующие действия:

1. Запустит контейнер Postgres.
2. Запустит контейнер с миграциями, выполнит их, а затем остановит.
3. Запустит контейнер бэкэнда.

Теперь сервис готов к использованию.

## Проверка работы

Get/currency принимает параметры val, date в формате 2006-01-02. Если дата не задана, то получаем курс по текущему дню.


Например: 
```
0.0.0.0:8000/currency?val=USD&date=2021-12-23
0.0.0.0:8000/currency?val=USD
```

### Доступные валюты
```
AUD, AZN, GBP, AMD, BYN, BGN, BRL, HUF, VND, HKD,
GEL, DKK, AED, USD, EUR, EGP, INR, IDR, KZT, CAD,
QAR, KGS, CNY, MDL, NZD, NOK, PLN, RON, XDR, SGD,
TJS, THB, TRY, TMT, UZS, UAH, CZK, SEK, CHF, RSD,
ZAR, KRW, JPY
```

