# DB schema
## searches
| Name       | Type         | Default           | Nullable | Comment    |
| ---------- | ------------ | ----------------- | -------- | ---------- |
| id         | char(36)     |                   | false    | ユーザーid |
| station    | varchar(100) | NULL              | true     | 検索駅名   |
| lat        | double(9,6)  | NULL              | true     | 検索緯度   |
| lng        | double(9,6)  | NULL              | true     | 検索経度   |
| created_at | datetime(6)  | CURRENT_TIMESTAMP | true     | 開始日時   |

## questions
| Name       | Type        | Default           | Nullable | Comment        |
| ---------- | ----------- | ----------------- | -------- | -------------- |
| id         | int(11)     |                   | false    | AUTO_INCREMENT |
| shop_id    | char(10)    |                   | false    | UNIQUE         |
| search_id  | char(36)    |                   | false    |                |
| number     | int         |                   | false    | 質問番号(1~7)  |
| created_at | datetime(6) | CURRENT_TIMESTAMP | true     | 作成日時       |

## shops
| Name          | Type          | Default           | Nullable | Comment      |
| ------------- | ------------- | ----------------- | -------- | ------------ |
| shop_id       | char(10)      |                   | false    |              |
| name          | varchar(200)  |                   | false    | お店の名前   |
| image         | varchar(2048) |                   | false    | 画像URL      |
| genre_code    | char(4)       |                   | false    | ジャンル     |
| subgenre_code | char(4)       | NULL              | true     | サブジャンル |
| price_code    | char(4)       |                   | false    | 価格帯       |
| created_at    | datetime      | CURRENT_TIMESTAMP | true     | 作成日時     |

## genres
| Name       | Type         | Default | Nullable | Comment |
| ---------- | ------------ | ------- | -------- | ------- |
| genre_code | char(4)      |         | false    |         |
| name       | varchar(100) |         | false    |         |

## prices
| Name       | Type         | Default | Nullable | Comment |
| ---------- | ------------ | ------- | -------- | ------- |
| price_code | char(4)      |         | false    |         |
| name       | varchar(100) |         | false    |         |

## gourmets
TODO: not use this table
| Name       | Type         | Default           | Nullable | Comment              |
| ---------- | ------------ | ----------------- | -------- | -------------------- |
| id         | int(11)      |                   | false    | AUTO_INCREMENT      |
| station    | varchar(100) | NULL              | true     | 駅名                 |
| lat        | double(9,6)  | NULL              | true     | 検索緯度             |
| lng        | double(9,6)  | NULL              | true     | 検索経度             |
| shops      | text         |                   | false    | HotPepperAPI(店一覧) |
| created_at | datetime     | CURRENT_TIMESTAMP | true     | 作成日時             |
