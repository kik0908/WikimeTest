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
	  Минимальный набор предоставляемой информации для привлечения пользователя. В начале странице находится баннер с коротким описанием сайта и кнопкой “Перейти к статьям” для перехода на страницу со списком статей, а также с постером и названием самого популярного на момент открытия сайта аниме, при нажатии на которые пользователь переходит на страницу этой статьи. В нижней части страницы представлены первые 16 популярных статей в виде карточек с постером и названием, на карточках по кнопке “Подробнее” можно перейти на страницу конкретной статьи.
  </p>
</details>

<details><summary><h3>Шапка сайта</h3></summary>
  <p> 
	  Доступна на всех страницах веб-приложения.
Для неавторизованных пользователей состоит из ссылки на главную страницу в виде логотипа “WIKIME” и на страницу со списком статей “Статьи”. Также в ней находится кнопка “Войти”, нажав на которую пользователь может авторизоваться через VK, перейдя на страницу от этого ресурса и введя свои данные.
	  
У авторизованных пользователей вместо кнопки входа находится ссылка на личную страницу в виде изображения своего аватара.
	  
Модераторам отображается ссылка на страницу добавления статьи в виде “Добавить”. 
Администраторам предоставляется ссылка на “Админ” страницу.

  </p>
</details>

<details><summary><h3>Список статей</h3></summary>
  <p> 
	 На странице пользователю предоставляются статьи об аниме в двух вариантах: списком или в виде таблицы, вид можно выбрать на панели перед карточками со статьями. Имеется возможность фильтровать информацию по “популярности”, “обновлению”, “рейтингу” и “дате выхода” при нажатии на соответствующую ссылку, а также выбирать интересующие статьи по жанрам, нажав на интересующие жанры на боковой панели. 
	  
Кроме того, пользователь имеет возможность найти конкретную статью с помощью поиска в начале странице, который ищет статьи по названию и описанию.
	  
Переход по страницам списка осуществляется внизу страницы путем нажатия на интересующую страницу или с помощью кнопок “Назад” или “Дальше”.

  </p>
</details>

<details><summary><h3>Статья</h3></summary>
  <p> 
Информация об аниме, которая включает в себя: название, оригинальное название, список жанров, режиссера, дату выхода, постер, рейтинг, описание и автора статьи, арты и кадры. Дополнительно на этой странице пользователь может ознакомиться с количеством людей, оценивших данную статью, а также отдельно с количеством людей, которые добавили это аниме в свой список избранных. Также на данной странице представлены комментарии авторизованных пользователей.
	  
Авторизованные пользователи имеют возможность на этой странице оценить статью, нажав на кнопку “Оценить” и выбрав соответствующую оценку по пятибалльной шкале, добавить в избранное, нажав на кнопку “Добавить в избранное”, и написать комментарий, написав текст в поле ввода и нажав кнопку “Отправить”.
	  
Если статью просматривает пользователь, добавивший её, то ему предоставляется возможность отредактировать информацию, нажав на кнопку “Редактировать”.
	  
	  Администраторы на данной странице могут перейти к редактированию статьи по кнопке “Редактировать”, а также удалить комментарии любых пользователей, нажав на “крестик” в карточке комментария.


  </p>
</details>

<details><summary><h3>Профиль пользователя</h3></summary>
  <p> 
	  На странице отображается имя и аватарка пользователя, которые при желании он может отредактировать, нажав на кнопки “Изменить никнейм” и “Изменить аватар”. Здесь же предоставляется список добавленных в избранное статей, просмотренных статей, которые открываются при нажатии на “стрелку-вниз”, закрытие возможно по этой же кнопке уже в виде “стрелки-вверх”. Просмотр также как и на странице со статьями доступен в двух вариантах и также реализована пагинация.
	  
Для модераторов и администраторов дополнительно отображается список добавленных ими статей.
	  
На этой же странице пользователь может выйти из своего аккаунта по кнопке “Выйти”.

  </p>
</details>

<details><summary><h3>Страница добавления статьи</h3></summary>
  <p> 
     1) Форма, состоящая из всех текстовых полей, необходимых для добавления новой статьи: название, оригинальное название, режиссер, список жанров, дата выхода и описание. Все поля являются обязательными. После заполнения пользователь переходит на следующую страницу, нажав на кнопку “Далее”.
	  
2) Добавление всех изображений для статьи: постер и арты, осуществляется нажатием на кнопки “Загрузить изображение”. Постер является обязательным полем, арты пользователь может добавлять в любом количестве или не добавлять вовсе. При этом во время создания новой статьи пользователь может редактировать загруженные файлы: заменять постер по кнопке “Заменить изображение”, удалять изображения для артов, нажав на соответствующую кнопку “крестик” у добавленного изображения, и добавлять дополнительные арты по кнопке “Загрузить изображение”. Также есть возможность для перехода на прошлую страницу с текстовой формой по кнопке “Назад”.
	  
После заполнения всех полей пользователь переходит на страницу добавленного аниме, нажав кнопку “Добавить”.

  </p>
</details>

<details><summary><h3>Админ-панель</h3></summary>
  <p> 
      Предоставляется только администраторам. На странице указаны списки администраторов и модераторов сайта с возможностью их изменения, а именно: изменение роли посредством удаления или добавления из\в списка.
	  
Удаление осуществляется путём нажатия на соответствующую кнопку “крестик” напротив карточки пользователя. Для добавления следует нажать на “плюс” у названия списка Администраторы/Модераторы и в появившейся форме ввести id пользователя, после чего нажать на кнопку “Добавить”. 
	  
Возможности на этой странице соответствуют роли пользователя.

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
      * Body: file - изображение, формат файла jpg или png
      * Description: Загрузить новую фотографию в статью
  1) /anime/{anime_id}/poster
      * Access: админ или модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id аниме, у которого меняется постер
      * Params: None
      * Method: POST
      * Content-type: form-data
      * Body: file - изображение, формат файла jpg или png
      * Description: Изменение постера у статьи	
  1) /anime/{anime_id}/images/{img_name}
      * Access: админ или модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id аниме, из статьи про которое удаляется фотография; img_name - название фото для удаления
      * Params: None
      * Method: DELETE
      * Content-type: None
      * Body: None
      * Description: Удалить фотографию
  1) /users/current/avatar
      * Access: Текущий пользователь
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: POST
      * Content-type: form-data
      * Body: file - изображение, формат файла jpg или png
      * Description: Изменить аватар пользователя
  1) /users/{user_id}
      * Access: Все
      * Authorization: False
      * Variables: user_id - id пользователя
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить информацию о пользователе
  1) /users/current
      * Access: Все
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить информацию об аутентифицированном пользователе
  1) /users/admins
      * Access: Админы
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить список администраторов
  1) /users/moderators
      * Access: Админы
      * Authorization: True
      * Variables: None
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
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
  1) /anime/{anime_id}
      * Access: Все
      * Authorization: False
      * Variables: anime_id - id аниме
      * Params: None
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить аниме по id
  1) /anime/list
      * Access: Все
      * Authorization: False
      * Variables: None
      * Params: id - id аниме, можно передать список таких параметров. Пример: id=1&id=2&id=3&id=4
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить аниме по id списком
  1) /anime
      * Access: Все
      * Authorization: False
      * Variables: None
      * Params: sortBy - сортировка аниме, может быть rating, dateAdded, favorites, releaseDate; order - порядок сортировки, прямой(1) или обратный(-1), необязательный; genres - массив жанров для посика, не должен указываться, если фильтр по жанрам не требуется, необязательный
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Получить отсортированные и отфильтрованные аниме
  1) /anime/{anime_id}/rating
      * Access: Все
      * Authorization: False
      * Variables: anime_id - id  аниме
      * Params: None
      * Method: POST
      * Content-type: json
      * Body: rating - оценка, может быть 1,2,3,4,5
      * Description: Оценить аниме
  1) /anime/popular
      * Access: Все
      * Authorization: False
      * Variables: None
      * Params: count - количество, максимум 30
      * Method: Get
      * Content-type: None
      * Body: None
      * Description: Получить список популярных аниме
  1) /anime
      * Access: Все
      * Authorization: False
      * Variables: None
      * Params: search - по какому тексту будет производится поиск
      * Method: GET
      * Content-type: None
      * Body: None
      * Description: Найти статью 
  1) /anime/{anime_id}
      * Access: Админ, модератор, который создал статью
      * Authorization: True
      * Variables: anime_id - id  аниме
      * Params: None
      * Method: PUT
      * Content-type: json
      * Body: title - название аниме; originTitle - оригинальное название; description - описание аниме; director - режиссер или студия; genres - список жанров; releaseDate - дата выпуска аниме в unix формате
      * Description: Изменить статью
  
  
# Технологии разработки
#### Frontend
  - HTML, CSS, TypeScript
  - [React](https://reactjs.org/) - JavaScript-библиотека для создания пользовательских интерфейсов
  - [React Router](https://reactrouter.com/en/6.5.0) - маршрутизация на стороне клиента
  - [Redux](https://redux.js.org/) - JavaScript-библиотека для управления состоянием приложения
  - [Redux Toolkit](https://redux-toolkit.js.org/) - библиотека для работы с хранилищем и асинхронной логикой

#### Backend
  - Golang 1.18
  - Golang standard library, net/http для обработки входящих соединений и отправки ответов
  - Gorilla/mux для маршрутизации
  - Gorilla/handlers для настройки CORS
  - mongo-driver - официальная библиотека для работы с Mongodb из Golang
  - github.com/go-playground/validator/v10 для валидации тел запросов
  - github.com/JeremyLoy/config для чтения конфигурации

#### СУБД
  - MongoDB
