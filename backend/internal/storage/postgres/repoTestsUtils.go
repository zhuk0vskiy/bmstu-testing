package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	DbName = "saladsTesting"
	DbUser = "user"
	DbPass = "pass"
	Port   = "5432/tcp"
	Image  = "postgres:14-alpine"
)

const (
	TestDataDir     = "../../../tests/test_data/"
	TearDownSQlsDir = "../../../tests/test_data_drop"
)

type TestDatabase struct {
	DbInstance *pgxpool.Pool
	DbAddress  string
	container  testcontainers.Container
}

func SetupTestDatabase(dbName string) *TestDatabase {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	container, dbInstance, dbAddr, err := createContainer(ctx, dbName)
	if err != nil {
		log.Fatal("failed to setup test: ", err)
	}

	err = migrateDb(dbAddr, dbName)
	if err != nil {
		log.Fatal("failed to perform db migration: ", err)
	}
	cancel()

	return &TestDatabase{
		container:  container,
		DbInstance: dbInstance,
		DbAddress:  dbAddr,
	}
}

func (tdb *TestDatabase) TearDown() {
	tdb.DbInstance.Close()
	_ = tdb.container.Terminate(context.Background())
}

func createContainer(ctx context.Context, dbName string) (testcontainers.Container, *pgxpool.Pool, string, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": DbPass,
		"POSTGRES_USER":     DbUser,
		//"POSTGRES_DB":       DbName,
		"POSTGRES_DB": dbName,
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        Image,
			ExposedPorts: []string{Port},
			Env:          env,
			WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}

	log.Println("postgres container ready and running at port: ", p.Port())
	time.Sleep(time.Second)

	dbAddr := fmt.Sprintf("localhost:%s", p.Port())
	//db, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, DbName))
	db, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, dbName))
	if err != nil {
		return container, db, dbAddr, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, dbAddr, nil
}

func migrateDb(dbAddr string, dbName string) error {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}
	_ = path
	pathToMigrationFiles := "/Users/maximhalitov/Desktop/IU7/ppo/backend/tests/migrations"

	//databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, DbName)
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", DbUser, DbPass, dbAddr, dbName)
	m, err := migrate.New(fmt.Sprintf("file:%s", pathToMigrationFiles), databaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func ExecuteSQLsFromDir(db *pgxpool.Pool, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error while reading dir with test data: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			err = executeSQL(db, filePath)
			if err != nil {
				return fmt.Errorf("executing sql %s: %w", file.Name(), err)
			}
		}
	}
	return nil
}

func executeSQL(db *pgxpool.Pool, filePath string) error {
	scriptContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading sql script: %w", err)
	}

	_, err = db.Exec(context.Background(), string(scriptContent))
	if err != nil {
		return fmt.Errorf("execution sql script: %w", err)
	}
	return nil
}
