# otus-ha-hw1
ДЗ Проблемы высоких нагрузок
Требуется разработать создание и просмотр анкет в социальной сети.

**Функциональные требования:**
1. Авторизация по паролю
2. Страница регистрации, где указывается следующая информация:
   * Имя
   * Фамилия
   * Возраст
   * Пол
   * Интересы
   * Город
3. Страницы с анкетой
4. Сделать инструкцию по локальному запуску приложения
5. Приложить Postman-коллекцию

**Нефункциональные требования:**
1. Любой язык программирования
2. В качестве базы данных использовать MySQL (PostgreSQL/MariaDB)
3. Не использовать ORM
4. Программа должна представлять из себя монолитное приложение.
5. Не рекомендуется использовать следующие технологии:
   * Репликация
   * Шардинг
   * Индексы
   * Кэширование
6. Фронт опционален

**Запуск**
1. make docker-start
2. make docker-stop
3. make newman
