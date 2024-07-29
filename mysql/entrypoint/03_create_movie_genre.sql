DROP TABLE IF EXISTS movie_genre;

CREATE TABLE movie_genre
(
  id int auto_increment,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  movie_id int not null,
  genre_id int not null,

  primary key (id),
  foreign key (movie_id) references movie(id) on delete cascade,
  foreign key (genre_id) references genre(id) on delete cascade
);


INSERT INTO movie_genre (movie_id, genre_id) VALUES
(1, 1),
(1, 11),
(1, 15),
(4, 1),
(4, 7),
(4, 11),
(4, 15),
(4, 18)
;
