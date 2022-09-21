# Wikime
Проект для дисциплин: Web-разработка, Введение в базы данных.
***
### Описание
Сайт для просмотра информации об аниме с возможностями добавления статей, их комментированием и оцениванием, добавлением в списки просмотренного и/или избранного.

### Предметная область
Информация об Аниме

# Данные
## Коллекции

_<details><summary><h3>ID base</h3></summary>_
  <p>
	  Коллекция нужна для хранения последнего индекса в коллекциях _Anime_ и _Users_.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------ | :---: | :-----------: | :--: | :----------------: |
| \_id | string| can only be one of: "anime", "user" | + | |
| seq | int | >0 | | _Anime.\_id_ OR _Users.\_id_ 
  </p>
</details>


_<details><summary><h3>Anime</h3></summary>_
  <p> 
Коллекция для хранения наполнения статей.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------ | :---: | :-----------: | :--: | :----------------: |
| \_id | int | >0, not null|  + | _Comments.\_id_    |
| Title| string | not null, len>0| | | |
| Origin Title | string | not null, len>0| | |
| Genres | string[], _*index*_ | not null, one of the _Genres.geners_| | |
| Description | string | | | |
| Poster | string, path to img | must be valid, points to an existing file | | |
| Images | string[] | must be valid, points to an existing file | | |
| URLs | string[] | | | |
| Director | string | | | |
| Date Added | date | not null | | |
| Release date | date | | | |
| Author | int | >0, not null | | _Users.\_id_ |
| Rating | _Rating_ struct, _*index*_ | not null | | |
</p>
</details>

_<details><summary><h3>Comments</h3></summary>_
<p> 
Коллекция для хранения комментариев.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int | >0, not null | + | 
| Comments | _Comment_ struct[] | not null
</p>
</details>

_<details><summary><h3>Vk</h3></summary>_
<p> 
Коллекция для сопоставления id пользователя с сайта  <a href="https://vk.com/">vk.com</a> с внутренним id в приложении.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int | >0, not null, valid vk user id | + |
| Inner Id | int | >0, not null | | _Users.\_id_
</p>
</details>
	  
_<details><summary><h3>Google</h3></summary>_
<p> 
Коллекция для сопоставления id пользователя с сайта  <a href="https://google.com/">google.com</a> с внутренним id в приложении.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int | >0, not null, valid vk user id | + |
| Inner Id | int | >0, not null | | _Users.\_id_
</p>
</details>

_<details><summary><h3>Users</h3></summary>_
<p> 
Коллекция для хранения информации о пользователях.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int | >0, not null | + | |
| Nickname | string | len > 0, not null 
| Photo | string, path to img | must be valid, points to an existing file
| Role  | string | not null, can only be one of: "admin", "moder", "user"
| Favorites | int[] | not null, length can be equal to 0 | | _Anime.\_id_
| Viewed | int[] | not null, length can be equal to 0 | | _Anime.\_id_
| Added | int[] | not null, length can be equal to 0 | | _Anime.\_id_
	  
</p>
</details>


## Структуры
_<details><summary><h3>Rating</h3></summary>_
<p> 
В каждом поле хранится количество соответствующих оценок.

| Название атрибута | Тип | Ограничения | Внешний ключ для |
| ------------------|:---:|:-----------:|:----------------:|
| Five | int | >=0, not null | |
| Four | int | >=0, not null | |
| There | int | >=0, not null | |
| Two | int | >=0, not null | |
| One | int | >=0, not null | |
| InFavorites | int | >=0, not null | | 
</p>
</details>

_<details><summary><h3>Comment</h3></summary>_
<p> 

| Название атрибута | Тип | Ограничения | Внешний ключ для |
| ------------------|:---:|:-----------:|:----------------:|
| User Id | int | >0, not null | _Users\_.id_
| Message | string | len > 0, not null | |
| DateTime | date | not null |
	  
</p>
</details>

## Общие ограничения целостности
  - Максимальная длина для типа string - 3072, если не указано иного
  - В коллекции _Anime_ и _Users_ нельзя вставлять документы в обход получения индекса для соответствующей коллекции из _ID base_
  - Если для поля указан внешний ключ, то должен существовать документ, на который указывает этот ключ
  - Для коллекций _Google_ и _Vk_ вставка происходит _*только при первой авторизации*_. В этих коллекциях поле \_id должно соответствовать реальному id на соответствующем ресурсе
  - Для каждого документа в коллекции _Users_ должен найтись документ или в коллекции _Vk_, или в коллекции _Google_ такой, что _Users.\_id_ == _Vk.\_id_ или _Users.\_id_ == _Google.\_id_ соответственно. _Исключениями являются заранее добавленные для презентации боты_
  - В любой коллекции поле _\_id_ является индексом

# Пользовательские роли
1. **Неавторизорованный пользователь** - может просматривать статьи и комментарии к ним, профили пользователей. Безграничное количество
2. **Обычный пользователь** - может ставить оценки, добавлять в свой список избранного/просмотренного, добавлять комментарии к существующим статьям, редактировать свой профиль, а именно: изменять аватарку и никнейм. Количество ограничено размером базы данных
3. **Модератор** - может добавлять статьи, редактировать созданные им статьи. Количество ограничено размером базы данных
4. **Администратор** - может изменять роли пользователей, может редактировать любые статьи. Количество ограничено размером базы данных

Роли расположены в порядке возрастания приоритета. Каждый пользователь дополнительно имеет возможности пользователей с более низким приоритетом.


# UI / API
## UI

<details><summary><h3>Главная страница</h3></summary>
  <p> 
	  Минимальный набор информации с красивым оформлением. Будет показан красивый банер с одной из статей(статья раз в 24 часа случайно выбирается из популярных статей) и список популярнейших статей. Банер и список кликабельны.
  </p>
</details>

<details><summary><h3>Шапка и футер</h3></summary>
  <p> 
	  В шапке будет представлена ссылка для перехода на главную страницу, строка поиска статьи по определенному аниме, кнопка для авторизации, а также, в зависимости от роли пользователя, кнопки для добавления статей и управления списками модераторов и администраторов.<br> В футере будут разнообразные ссылки.
	  
  </p>
</details>

<details><summary><h3>Список статей</h3></summary>
  <p> 
	  Будет отображаться список статей с возможностью выборки статей в определенных жанрах их последующей сортировкой по рейтингу/дате обновления/дате выхода/популярности. Смотреть статьи можно в двух вариантах: таблицей или списком.
  </p>
</details>

<details><summary><h3>Статья</h3></summary>
  <p> 
	 На странице будут представлены: название, общая информация о тайтле, постер, оценки, средняя оценка, арты/кадры, комментарии.
  </p>
</details>

<details><summary><h3>Профиль пользователя</h3></summary>
  <p> 
	  На этой странице будут отображаться никнейм и аватарка, а также списки избранного и просмотренно данного пользователя. Если пользователь на странице своего аккаунта, то будут отображаться кнопки для изменения аватарки и никнейма. 
  </p>
</details>

<details><summary><h3>Страница добавления статей</h3></summary>
  <p> 
      Будут отображаться поля для заполнения новой статьи.
  </p>
</details>

<details><summary><h3>Страница администраторов</h3></summary>
  <p> 
      Страница нужна для управления модераторским и администраторским составом. Для управления будут представлены два списка(список админов и список модераторов) с возможностью добавления пользователей в список и удаления неугодных из него. Доступна только для пользователей с ролью "admin".
  </p>
</details>

## API
  - Будет реализовано API в стиле REST
  - Будет реализована страничка с документацией по каждому возможному запросу
  - Авторизация будет производиться с помощью протокола OAuth 2.0 через сторонние сервисы(например [vk.com](https://vk.com/)) 
  - Доступ к фотографиям будет осуществляться через GET запрос по URL, который возвращает сервер в качестве пути к фото

# Технологии разработки
#### Frontend
  - HTML, CSS
  - JavaScript
  - React 18.2.0

#### Backend
  - Golang 1.18
  - Golang standard library, net/http для обработки входящих соединений и отправки ответов
  - Gorilla/mux для маршрутизации
  - Gorilla/sessions для хранения сессий

#### СУБД
  - MongoDB

# Тестирование
#### Frontend
  - Jest
#### Backend
  - Gotest
