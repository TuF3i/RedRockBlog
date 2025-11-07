package runc

import (
	"RedRock/core"
	"RedRock/core/api"
	"RedRock/core/dao/mySQL"
	"RedRock/core/utils/config"
	llog "RedRock/core/utils/log"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gitee.com/liumou_site/logger"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func GoWork() {
	logs := logger.NewLogger(1)
	logs.Modular = "init"
	env := RunServer{l: logs}

	env.initConfig()
	env.initLog()
	env.initMySQL()
	env.initBundle()

	env.RunGin()

	fmt.Println()
	c := color.New(color.FgGreen).Add(color.Underline)
	_, _ = c.Println("[GoWork Process Down! (Press Ctrl-C to exit)]")
	c = color.New(color.FgGreen)
	_, _ = c.Println("\t--Present System Log:")
}

func (root *RunServer) initLog() {
	root.l.Debug("[initLog] Started to load mod<LocalLog>")

	flog := log.Logger{}
	flog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flog.SetPrefix("[FileLog] ")

	clog := logger.NewLogger(1)

	date := time.Now().Format("20060102")

	core.Logger = &llog.Log{
		LogPath:  "data/log",
		LogLevel: "DEBUG",
		CLog:     clog,
		FLog:     &flog,
		ToDay:    date,
	}
	root.l.Info("[initLog] Successfully loaded mod<LocalLog>")
}

func (root *RunServer) initConfig() {
	root.l.Debug("[initConfig] Started to load mod<LocalConfig>")
	localConfig, err := config.InitConfig("data/config/config.json")
	if err != nil {
		root.l.Error("[initConfig] error<%v> in loading mod<LocalConfig>", err.Error())
		os.Exit(1)
	}

	core.GlobalConf = localConfig
	root.l.Info("[initConfig] Successfully loaded mod<LocalConfig>")
	c := color.New(color.FgCyan).Add(color.Underline)
	_, _ = c.Println("Configuration Details: ")
	color.Cyan("\tOidcProvider: %v\n", localConfig.OidcProvider)
	color.Cyan("\tClientID:     %v\n", localConfig.ClientID)
	color.Cyan("\tClientSecret: %v\n", localConfig.ClientSecret)
	color.Cyan("\tDomain:       %v\n", localConfig.Domain)
	color.Cyan("\tApiPort:      %v\n", localConfig.ApiListeningPort)
	color.Cyan("\tMySQLAddr:    %v\n", localConfig.MySQLAddr)
	color.Cyan("\tMySQLPort:    %v\n", localConfig.MySQLPort)
	color.Cyan("\tMySQLUser:    %v\n", localConfig.MySQLUser)
	color.Cyan("\tMySQLPassword:%v\n", localConfig.MySQLPassword)
	color.Cyan("\tMySQLDBName:  %v\n", localConfig.MySQLDBName)
	color.Cyan("\tI18nConfPath: %v\n", localConfig.I18nProfilePath)
	color.Cyan("\tMaxCommLevel: %v\n", localConfig.MaxCommentLevel)
}

func (root *RunServer) initMySQL() {
	root.l.Debug("[initMySQL] Started to load mod<Gorm>")

	db := mySQL.InitMySQL()

	err := db.GetConnection()
	if err != nil {
		root.l.Error("[initMySQL] error<%v> in connecting MySQL", err.Error())
		os.Exit(1)
	}

	err = db.MigrateDataBase()
	if err != nil {
		root.l.Error("[initMySQL] error<%v> in migrating DataBase", err.Error())
		os.Exit(1)
	}

	core.DataBase = db.DB
	root.l.Info("[initMySQL] Successfully loaded mod<Gorm>")
	c := color.New(color.FgCyan).Add(color.Underline)
	_, _ = c.Printf("\tGet MySQL Connection <%v:%v> at Addr <%v>\n",
		core.GlobalConf.MySQLAddr,
		core.GlobalConf.MySQLPort,
		&core.DataBase)
}

func (root *RunServer) initBundle() {
	root.l.Debug("[initBundle] Started to load mod<i18n>")

	var bundle *i18n.Bundle

	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, configurationFile := range core.GlobalConf.I18nProfilePath {
		_, err := bundle.LoadMessageFile(configurationFile)
		if err != nil {
			root.l.Error("[initBundle] error<%v> in loading mod<Bundle>", err.Error())
			os.Exit(1)
		}
	}

	core.Bundle = bundle
	root.l.Info("[initBundle] Successfully loaded mod<i18n>")
}

func (root *RunServer) RunGin() {
	root.l.Debug("[RunGin] Started to launch Gin Engine")
	router := api.InitGin().InitGinApi()

	go func(router *gin.Engine) {
		err := router.Run(fmt.Sprintf("%v:%v", core.GlobalConf.Domain, core.GlobalConf.ApiListeningPort))
		if err != nil {
			core.Logger.MyError(fmt.Sprintf("System Panic : %v", err.Error()))
			os.Exit(1)
		}
		root.l.Info("RedRockBlog Started Successfully!")
	}(router)
}
