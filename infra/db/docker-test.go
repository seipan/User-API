package db

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ory/dockertest"
)

var (
	user     = "root"
	password = "hoge"
	dbName   = "hoge"
	port     = "5433"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	// Dockerコンテナへのファイルマウント時に絶対パスが必要
	pwd, _ := os.Getwd()
	log.Println(pwd)
	strings.Replace(pwd, "/infra/db", "", -1)

	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	// テーブル定義などはここで流し込むのが良さそう
	runOptions := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.0",
		Env: []string{
			" POSTGRES_USER=root",
			"POSTGRES_PASSWORD=hoge",
			"POSTGRES_DB=hoge",
		},
		Mounts: []string{
			pwd + "/db/init:/docker-entrypoint-initdb.d", // コンテナ起動時に実行したいSQL
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		// DBコンテナが立ち上がってから疎通可能になるまで少しかかるのでちょっと待ったほうが良さそう
		time.Sleep(time.Second * 10)

		var err error
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}
