# 初期テーブル構築
=> sql 配下に *.sql ファイルを置けば、DBコンテナ構築時にファイルを実行

# 起動
$ docker-compose up

# 停止
コンソール上で「ctrl + c」

# DB諸設定
- docker-compose.ymlの以下を編集
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
environment:
- MYSQL_ROOT_PASSWORD=secret #rootのパスワード
- MYSQL_DATABASE=revel #revelデータベースの作成
- MYSQL_USER=revel #revelユーザの作成
- MYSQL_PASSWORD=secret #revelユーザのパスワード
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# DBアクセス例
- DBコンテナに入って以下を実行  
$ mysql -urevel -psecret

# アクセス例
http://localhost:9000 # appコントローラー  
http://localhost:9000/users # userコントローラー

# テスト実行例（CUI）
- コンテナに入って以下を実行  
$ revel test myapp # apptest.go に記述のあるすべてのテスト  
$ revel test myapp AppTest.TestThatIndexPageWorks # apptest.go に記述のある特定のテスト  
=> テスト実行後は myapp 配下の test-results に結果が格納されます。  

# テスト実行例（GUI）
http://localhost:9000/@tests
=> テスト実行後は myapp 配下の test-results に結果が格納されます。

# テスト関連参考
https://qiita.com/shiwork/items/eeac5dd374a226e460e0  
https://revel.github.io/modules/testing.html
