DROP TABLE IF EXISTS genre_neo;

CREATE TABLE genre_neo
(
  id int auto_increment,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  code varchar(255) not null,
  name varchar(255) not null,
  movie_id int not null,
  primary key (id),
  foreign key (movie_id) references movie_neo(id) on delete cascade
);


INSERT INTO genre_neo (id, code, name, movie_id) VALUES
(1, 'action', 'アクション', 1),
(7, 'drama', 'ドラマ', 4),
(11, 'horror', 'ホラー', 1),
(12, 'horror', 'ホラー', 4),
(15, 'sf', 'SF', 1),
(16, 'sf', 'SF', 4),
(18, 'war', '戦争', 4)
;
