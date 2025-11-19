package config

type ConfigRoot struct {
	conf       Config
	configPath string
}

type Config struct {
	RedirectURL string //OAuth2回调URL

	OidcProvider     string   `json:"OidcProvider"`     //Oauth2认证提供方
	ClientID         string   `json:"ClientID"`         //OAuth2 ClientID
	ClientSecret     string   `json:"ClientSecret"`     //OAuth2 ClientSecret
	Domain           string   `json:"Domain"`           //服务器的ip或域名
	ApiListeningPort string   `json:"apiListeningPort"` //API的监听段端口
	MySQLAddr        string   `json:"MySQLAddr"`        //MySQL服务器的地址
	MySQLPort        string   `json:"MySQLPort"`        //MySQL端口
	MySQLUser        string   `json:"MySQLUser"`        //MySQL用户名
	MySQLPassword    string   `json:"MySQLPassword"`    // MySQL密码
	MySQLDBName      string   `json:"MySQLDBName"`      //MySQL数据库
	I18nProfilePath  []string `json:"i18nProfilePath"`  //i18n多语言配置文件路径
	MaxCommentLevel  uint     `json:"MaxCommentLevel"`  //评论最大的嵌套层数
}
