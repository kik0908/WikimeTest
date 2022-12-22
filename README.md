# Wikime
Проект для дисциплин: Web-разработка, Введение в базы данных.
***
Выполнили:
- Frontend: Таисия Кваша, 20.Б12пу
- Backend: Чередников Кирилл, 20.Б12пу
### Описание
Сайт для просмотра информации об аниме с возможностями добавления статей, их оцениванием, добавлением в списки просмотренного и/или избранного.

### Предметная область
Информация об Аниме

# Данные
## Коллекции

_<details><summary><h3>Anime</h3></summary>_
  <p> 
Коллекция для хранения наполнения статей.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------ | :---: | :-----------: | :--: | :----------------: |
| \_id | int64 | >0, not null|  + |     |
| Title| string | not null, len>0| | | |
| Origin Title | string | not null, len>0| | |
| Genres | string[], _*index*_ | not null, one of the _Genres.Geners_| | |
| Description | string | | | |
| Poster | string, path to img | must be valid, points to an existing file | | |
| Images | string[] | must be valid, points to an existing file | | |
| Director | string | | | 
| Release date | date | | | |
| Date added | date | | | |
| Author | int | >0, not null | | _Users.\_id_ |
| Rating | _Rating_ struct, index | not null
</p>
</details>

_<details><summary><h3>Genres</h3></summary>_
  <p> 
Коллекция для хранения жанров.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------ | :---: | :-----------: | :--: | :----------------: |
| \_id | string="Genres" | | + | |
| Genres | string[], _*index*_ | not null| | |
</p>
</details>

_<details><summary><h3>Vk</h3></summary>_
<p> 
Коллекция для сопоставления id пользователя с сайта  <a href="https://vk.com/">vk.com</a> с внутренним id в приложении.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int64 | >0, not null, valid vk user id | + |
| Inner Id | int | >0, not null | | _Users.id_
</p>
</details>

_<details><summary><h3>Users</h3></summary>_
<p> 
Коллекция для хранения информации о пользователях.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | int64 | >0, not null | + | |
| Nickname | string | len > 0, not null 
| Avatar | string, path to img | must be valid, points to an existing file
| Role  | string | not null, can only be one of: "admin", "moder", "user"
| Favorites | int[] | not null, length can be equal to 0 | | _Anime.\_id_
| Viewed | int[] | not null, length can be equal to 0 | | _Anime.\_id_
| Rated | struct{\_id: int64, Rate: int}[] | not null, length can be equal to 0 | | \_id -> _Anime.\_id_
	  
</p>
</details>

_<details><summary><h3>IdBase</h3></summary>_
<p> 
Коллекция для хранения id.

| Название атрибута | Тип | Ограничения | PR | Внешний ключ для |
| ------------------|:---:|:-----------:|:--:|:----------------:|
| \_id | string, {AnimeID, UserID} | not null | + | |
| LastId | int64 | not null | | _Anime.\_id_ or _Users.\_id_  |
	  
</p>
</details>

## Структуры

_<details><summary><h3>Rating</h3></summary>_
<p> 
В каждом поле хранится количество соответствующих оценок для данного аниме.

| Название атрибута | Тип | Ограничения | Внешний ключ для |
| ------------------|:---:|:-----------:|:----------------:|
| Five | int | >=0, not null | |
| Four | int | >=0, not null | |
| There | int | >=0, not null | |
| Two | int | >=0, not null | |
| One | int | >=0, not null | |
| InFavorites | int64 | >=0, not null | |
| Average | float | in range [0, 5], not null |  
| Watched | int64 | not null |  
</p>
</details>

## Общие ограничения целостности
  - В коллекции _Anime_ и _Users_ нельзя вставлять документы в обход получения индекса для соответствующей коллекции из _IdBase_
  - Если для поля указан внешний ключ, то должен существовать документ, на который указывает этот ключ
  - Для коллекции _Vk_ вставка происходит _*только при первой авторизации*_. Поле \_id должно соответствовать реальному id на ресурсе vk.com
  - Для каждого документа в коллекции _Users_ должен найтись документ в коллекции _Vk_ такой, что _Users.\_id_ == _Vk.\_id_. _Исключениями являются заранее добавленные для презентации боты_
  - В любой коллекции поле _\_id_ является индексом

# Пользовательские роли
1. **Неавторизорованный пользователь** - может просматривать статьи и использовать поиск. Безграничное количество
2. **Обычный пользователь** - может ставить оценки, добавлять в свой список избранного/просмотренногом, редактировать свой профиль, а именно: изменять аватарку и никнейм. Количество ограничено размером базы данных
3. **Модератор** - может добавлять статьи, редактировать созданные им статьи. Количество ограничено размером базы данных
4. **Администратор** - может управлять списком модераторов, может редактировать любые статьи. Количество ограничено размером базы данных
5. **Root** - может управлять списком администраторов.

Роли расположены в порядке возрастания приоритета. Каждый пользователь дополнительно имеет возможности пользователей с более низким приоритетом.

# Дополнительные требования

  - При изменении статьи должна быть возможность поменять каждый элемент статьи



# UI / API
## UI

<details><summary><h3>Главная страница</h3></summary>
  <p> 
	  Минимальный набор информации с красивым оформлением. Будет показан красивый банер с одной из статей и список популярнейших статей. Банер и список кликабельны.
  </p>
</details>

<details><summary><h3>Шапка</h3></summary>
  <p> 
	  В шапке будет представлена ссылка для перехода на главную страницу, кнопка для авторизации, а также, в зависимости от роли пользователя, кнопки для добавления статей и управления списками модераторов и администраторов.
  </p>
</details>

<details><summary><h3>Список статей</h3></summary>
  <p> 
	  Будет отображаться список статей с возможностью выборки статей в определенных жанрах их последующей сортировкой по рейтингу/дате обновления/дате выхода/популярности. Смотреть статьи можно в двух вариантах: таблицей или списком. Перед списком будет доступно поле для поиска аниме, текстовый поиск происходит по названию и описанию.
  </p>
</details>

<details><summary><h3>Статья</h3></summary>
  <p> 
	 На странице будут представлены: название, общая информация о тайтле, постер, средняя оценка, арты/кадры.
  </p>
</details>

<details><summary><h3>Профиль пользователя</h3></summary>
  <p> 
	  На этой странице будут отображаться никнейм и аватарка, а также списки избранного и просмотренно данного пользователя. Если пользователь на странице своего аккаунта, то будут отображаться кнопки для изменения аватарки и никнейма. Если пользователь добавил какую-то статью, то будет отображаться список добавленных статей.
  </p>
</details>

<details><summary><h3>Страница добавления статей</h3></summary>
  <p> 
      Будут отображаться поля для заполнения новой статьи.
  </p>
</details>

<details><summary><h3>Страница администраторов</h3></summary>
  <p> 
      Страница нужна для управления модераторским и администраторским составом. Для управления будут представлены два списка(список админов и список модераторов) с возможностью добавления пользователей в список и удаления неугодных из него. Доступна только для пользователей с ролью "admin" или "root".
  </p>
</details>

## API
  - Будет реализовано API в стиле REST
  - Будет реализована страничка с документацией по каждому возможному запросу
  - Авторизация будет производиться с помощью протокола OAuth 2.0 через сторонние сервисы(например [vk.com](https://vk.com/)) 
  - Доступ к фотографиям будет осуществляться через GET запрос по URL, который возвращает сервер в качестве пути к фото
### Методы API
  1) /anime/{anime_id}/images
      * Access: админ или модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id аниме, в запись о котором добавляется фотография
      * Params: None
      * Method: POST
      * Content-type: form-data
      * Body: file = {file for uploadind}, формат файла jpg или png
      * Description: Загрузить новую фотографию в статью
  1) /anime/{anime_id}/poster
      * Access: админ или модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id аниме, у которого меняется постер
      * Params: None
      * Method: POST
      * Content-type: form-data
      * Body: file = {file for uploadind}, формат файла jpg или png
      * Description: Изменение постера у статьи
  1) /anime/{anime_id}/images/{img_name}
      * Access: админ или модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id аниме, из статьи про которое удаляется фотография; img_name - название фото для удаления
      * Params: None
      * Method: DELETE
      * Content-type: none
      * Body: none
      * Description: Удалить фотографию
  1) /users/current/avatar
      * Access: Текущий пользователь
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: form-data
      * Body: file = {file for uploadind}, формат файла jpg или png
      * Description: Изменить аватар пользователя
  1) /users/{user_id}
      * Access: Все
      * Authorization: False
      * Variables: user_id - id пользователя
      * Params: None
      * Method: GET
      * Content-type: none
      * Body: none
      * Description: Получить информацию о пользователе
  1) /users/current
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: none
      * Body: none
      * Description: Получить информацию об аутентифицированном пользователе
  1) /users/admins
      * Access: Админы
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: none
      * Body: none
      * Description: Получить список администраторов
  1) /users/moderators
      * Access: Админы
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: none
      * Body: none
      * Description: Получить список модераторов
  1) /users/current/nickname
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: PUT
      * Content-type: json
      * Body: nickname - новый никнейм
      * Description: Изменить никнейм
  1) /users/current/favorites
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: animeId - id аниме
      * Description: Добавить аниме в избранное
  1) /users/current/favorites
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: DELETE
      * Content-type: json
      * Body: animeId - id аниме
      * Description: Удалить аниме из избранного
  1) /users/current/watched
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: animeId - id аниме
      * Description: Добавить аниме в список просмотренного
  1) /users/current/watched
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: DELETE
      * Content-type: json
      * Body: animeId - id аниме
      * Description: Удалить аниме из списка просмотренного
  1) /users/{user_id}/role
      * Access: Админы
      * Authorization: True
      * Variables: user_id - id пользователя
      * Params: roleName - присваиваемая роль, может быть moderator, admin, user
      * Method: PUT
      * Content-type: None
      * Body: None
      * Description: Изменить роль пользователя
  1) /users/{user_id}/role
      * Access: Админы
      * Authorization: True
      * Variables: user_id - id пользователя
      * Params: None
      * Method: DELETE
      * Content-type: None
      * Body: None
      * Description: Сбросить роль до базовой
  1) /comments
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: message - текст комментария; anime_id - id аниме, в статье о котором создается комментарий
      * Description: Создать комментарий
  1) /anime/{anime_id}/comments
      * Access: Все
      * Authorization: False
      * Variables: anime_id - id статьи для получения комментариев
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить комментарии
  1) /comments/{comment_id}
      * Access: Админ, модератор, который создал статью, или пользователь, который написал этот комментарий.	
      * Authorization: True
      * Variables: comment_id - id удаляемого комментария
      * Params: None
      * Method: DELETE
      * Content-type: None
      * Body: None
      * Description: Удалить комментарий
  1) /auth/vk
      * Access: Все
      * Authorization: False
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: AuthToken - токен от vk api
      * Description: Аутентификация пользователя
  1) /anime
      * Access: Админ или модератор
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: title - название аниме; originTitle - оригинальное название; description - описание аниме; director - режиссер или студия; genres - список жанров; releaseDate - дата выпуска аниме в unix формате
      * Description: Создание аниме
  
  
# Технологии разработки
#### Frontend
  - HTML, CSS
  - JavaScript
  - React 18.2.0

#### Backend
  - Golang 1.18
  - Golang standard library, net/http для обработки входящих соединений и отправки ответов
  - Gorilla/mux для маршрутизации

#### СУБД
  - MongoDB
