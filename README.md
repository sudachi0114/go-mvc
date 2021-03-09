# Go で MVC/CRUD な web アプリケーションを作る

お勉強プロジェクトです。Go で (とても簡単な) webアプリケーションを作ります。
MVCフレームワークに則ったものにします。CRUD処理ができるようにします。

勉強したいこと
* Golang
  - Go で CRUD (DB 連携・処理)
  - Go で開発する時のパッケージング
* MVC (アプリケーションアーキテクチャ)

## 作るアプリケーション
* 『ユーザ』というモデルを持つ
* ユーザモデルの操作ができるコントローラーを持つ
* ユーザモデルのコントローラーは『登録 (Create)、表示 (Read)、修正(Update)、削除 (Delete)』ができる
* (コントローラーと、ルーティンは分離したい)
 

(余裕があれば API 化とかしたい。
あと、テストとかも勉強したいので組み込みたい。)


### インストールしたライブラリ

```shell
go get github.com/wcl48/valval      # UserModel のバリデーションのため
go get github.com/jinzhu/gorm       # Go ORM ライブラリ (DB とのやりとり用)
go get 	github.com/mattn/go-sqlite3 # Go で DB とやり取りする時のドライバ
```

注意: gorm を用いて DB に接続する際、**必ずドライバも import する** ようにする。

```go
import (
	"github.com/jinzhu/gorm"        // gorm
	_ "github.com/mattn/go-sqlite3" // gorm で使うドライバもインポートしないとエラーになる (この場合、sqlite3 のドライバを import )
)

func main() {
	db, err := gorm.Open("DBMS", "Connection")
	if err != nil {
		...
}

```

## 実行

### main

```shell
go run main.go
```

### DB migration

```shell
go run db/migrate.go
```

## Links
* [GoでCRUDでMVCなWEBアプリケーションを書く](https://qiita.com/masahikoofjoyto/items/b2e6c2cad447e48f91ee)
* [クラッド (CRUD)とは｜「分かりそう」で「分からない」](https://wa3.i-3-i.info/word123.html)