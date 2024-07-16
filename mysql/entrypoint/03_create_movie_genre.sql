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
  foreign key (genre_id) references movie(id) on delete cascade
);
