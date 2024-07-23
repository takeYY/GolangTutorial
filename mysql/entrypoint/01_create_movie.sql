DROP TABLE IF EXISTS movie;

CREATE TABLE movie
(
  id int auto_increment,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  title varchar(255) not null,
  overview text,

  primary key (id)
);


INSERT INTO movie (title, overview) VALUES
('ターミネーター', 'シュワちゃんカッコいい!!'),
('ターミネーター2', '液体金属'),
('ターミネーター3', 'T-X'),
('ターミネーター4', 'マーカス!!')
;
