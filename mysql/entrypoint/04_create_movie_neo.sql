DROP TABLE IF EXISTS movie_neo;

CREATE TABLE movie_neo
(
  id int auto_increment,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  title varchar(255) not null,
  overview text,

  primary key (id)
);


INSERT INTO movie_neo (id, title, overview) VALUES
(1, 'ターミネーター', 'シュワちゃんカッコいい!!'),
(2, 'ターミネーター2', '液体金属'),
(3, 'ターミネーター3', 'T-X'),
(4, 'ターミネーター4', 'マーカス!!')
;
