package driver

import (
	"fmt"
	"github.com/fatih/color"
	src "github.com/mindwingx/go-clean-arch-boilerplate"
	"github.com/mindwingx/go-clean-arch-boilerplate/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

type SqlAbstraction interface {
	InitSql()
	Migrate()
	Sql() *Sql
	Close()
	//Seed()
}

type (
	Sql struct {
		config config
		locale LocaleAbstraction
		DB     *gorm.DB
	}

	config struct {
		Debug              bool
		Host               string
		Port               string
		Username           string
		Password           string
		Database           string
		Ssl                string
		MaxIdleConnections int
		MaxOpenConnections int
		MaxLifetimeSeconds int
	}
)

func NewSql(registry RegistryAbstraction, locale LocaleAbstraction) SqlAbstraction {
	database := new(Sql)
	registry.Parse(&database.config)
	database.locale = locale
	return database
}

func (g *Sql) InitSql() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		g.config.Host,
		g.config.Username,
		g.config.Password,
		g.config.Database,
		g.config.Port,
		g.config.Ssl,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 g.newGormLog(),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_connection_init_failure"), err)
	}

	sqlDatabase, err := database.DB()
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_connection_retrieve_failure"), err)
	}

	if g.config.MaxIdleConnections != 0 {
		sqlDatabase.SetMaxIdleConns(g.config.MaxIdleConnections)
	}

	if g.config.MaxOpenConnections != 0 {
		sqlDatabase.SetMaxOpenConns(g.config.MaxOpenConnections)
	}

	if g.config.MaxLifetimeSeconds != 0 {
		sqlDatabase.SetConnMaxLifetime(time.Second * time.Duration(g.config.MaxLifetimeSeconds))
	}

	if g.config.Debug {
		database = database.Debug()
		color.Yellow(g.locale.Get("db_debug_mode_enable"))
	}

	g.DB = database
}

func (g *Sql) Migrate() {
	path := fmt.Sprintf("%s/database/postgres", src.Root())

	// Open the directory
	dir, err := os.Open(path)
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_sql_file_load_failure"), err)
		return
	}

	defer dir.Close()

	// Read the directory contents
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_sql_dir_scan_failure"), err)
		return
	}

	// Sort the entries alphabetically by name - Sql file order by numeric(01, 02, etc)
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].Name() < fileInfos[j].Name()
	})

	// Iterate over the file info slice and print the file names
	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() {
			if err = g.DB.Exec(g.parseSqlFile(path, fileInfo)).Error; err != nil {
				helper.CustomPanic(g.locale.Get("db_migrate_failure"), err)
			}
		}
	}
}

func (g *Sql) Sql() *Sql {
	return g
}

func (g *Sql) Close() {
	sql, err := g.DB.DB()
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_connection_close_failure"), err)
		return
	}

	sql.Close()
}

// HELPER METHOD

func (g *Sql) newGormLog() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             helper.SlowSqlThreshold * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,                           // Log level
			IgnoreRecordNotFoundError: false,                                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                  // Disable color
		})
}

func (g *Sql) parseSqlFile(path string, fileInfo os.FileInfo) string {
	sqlFile := fmt.Sprintf("%s/%s", path, fileInfo.Name())
	sqlBytes, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		helper.CustomPanic(g.locale.Get("db_sql_file_parse_failure"), err)
	}
	// Convert SQL file contents to string
	return string(sqlBytes)
}
