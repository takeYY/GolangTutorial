DROP TABLE IF EXISTS genre;

CREATE TABLE genre
(
  id int auto_increment,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  code varchar(255) not null,
  name varchar(255) not null,
  primary key (id)
);


INSERT INTO genre (id, code, name) VALUES
(1, 'action', 'アクション'),
(2, 'adventure', 'アドベンチャー'),
(3, 'animation', 'アニメーション'),
(4, 'comedy', 'コメディ'),
(5, 'crime', 'クライム'),
(6, 'documentary', 'ドキュメンタリー'),
(7, 'drama', 'ドラマ'),
(8, 'family', 'ファミリー'),
(9, 'fantasy', 'ファンタジー'),
(10, 'history', 'ヒストリー'),
(11, 'horror', 'ホラー'),
(12, 'musical', 'ミュージカル'),
(13, 'mystery', 'ミステリー'),
(14, 'romance', 'ロマンス'),
(15, 'sf', 'SF'),
(16, 'tv', 'TV'),
(17, 'thriller', 'スリラー'),
(18, 'war', '戦争'),
(19, 'western', '西部劇')
;
