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


INSERT INTO genre (code, name) VALUES
('action', 'アクション'),
('adventure', 'アドベンチャー'),
('animation', 'アニメーション'),
('comedy', 'コメディ'),
('crime', 'クライム'),
('documentary', 'ドキュメンタリー'),
('drama', 'ドラマ'),
('family', 'ファミリー'),
('fantasy', 'ファンタジー'),
('history', 'ヒストリー'),
('horror', 'ホラー'),
('musical', 'ミュージカル'),
('mystery', 'ミステリー'),
('romance', 'ロマンス'),
('sf', 'SF'),
('tv', 'TV'),
('thriller', 'スリラー'),
('war', '戦争'),
('western', '西部劇')
;
