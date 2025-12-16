<h1 align="center"> Привет! Я <a target="_blank"> Кармеев Артур из группы ЭФМО-01-25 </a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center"> Данная практика была выполнена супер поздно (16.12.2025)!</h3>


<h3 align="center"> Практика №5</h3>
<h3 align="center">Тема: Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)</h3>

## Структура работы:)

    └── pz5-db/
      ├── db.go
      ├── go.mod
      ├── go.sum
      ├── main.go
      ├── repository.go
      └── .idea/
          ├── .gitignore
          ├── modules.xml
          ├── pz5-db.iml
          └── workspace.xml


## 1. Установка постгри 15

<img width="717" height="566" alt="image" src="https://github.com/user-attachments/assets/b6032eaf-81c2-41ec-9564-d87f1422f464" />


рандомный скриншот установщика из интернета для вида


## 2. Создание БД, результаты запросов

В pgAdmin 4 создаем БД todo

<img width="480" height="510" alt="image" src="https://github.com/user-attachments/assets/1481a789-666f-44cd-91eb-32fa2855e6b2" />

Потом создаем таблицу tasks

<img width="1482" height="578" alt="image" src="https://github.com/user-attachments/assets/6b42a8b9-1d51-48d1-988a-2aa4a3a7b88e" />

Создаем первую задачу

<img width="1495" height="624" alt="image" src="https://github.com/user-attachments/assets/8743c6a6-5aea-4908-9579-458ad3d55eec" />

Псоле чего подключаемся к БД из Go с помощью database/sql и драйвера PostgreSQL. 3.	Выполняем параметризованные запросы INSERT и SELECT.

<img width="974" height="521" alt="image" src="https://github.com/user-attachments/assets/1c64a4d9-71f1-4d81-814b-873256b99bd0" />

Проверяем что все запросы внеслись в таблицу tasks

<img width="1495" height="923" alt="image" src="https://github.com/user-attachments/assets/ca3f17ac-ed3e-4c4b-bb02-5260dcafee43" />

## 3. Проверочные задания

1.	Реализуйте функцию, которая вернёт только выполненные (done=true) или невыполненные (done=false) задачи

<img width="954" height="850" alt="image" src="https://github.com/user-attachments/assets/b1e2616f-ed48-4075-bcd7-cfa7b305ec47" />

2.	Добавьте функцию:func (r *Repo) FindByID(ctx context.Context, id int) (*Task, error) и выведите в main подробности по задаче с конкретным id.

<img width="990" height="849" alt="image" src="https://github.com/user-attachments/assets/e5aecd4e-3c16-4e16-ace6-91c91b68107b" />

## 4. Ответики на вопросики

1) Что такое пул соединений *sql.DB и зачем его настраивать?	Что такое пул соединений *sql.DB и зачем его настраивать?

Кэш готовых подключений к БД для повторного использования, а не создания нового на каждый запрос.

Настраивается для:
Производительность — экономит 10-100мс на запрос (нет накладных расходов на установку соединения)

Защита БД — ограничивает одновременные подключения, чтобы не перегрузить СУБД

Стабильность — предотвращает утечки памяти и "протухшие" соединения


2) Почему используем плейсхолдеры $1, $2?

Автоматическая защита от SQL-инъекций. Данные пользователя никогда не подставляются в SQL напрямую.


3) Чем Query, QueryRow и Exec отличаются?

Query → много строк + итерация
QueryRow → одна строка + Scan
Exec → изменение данных (без SELECT)
