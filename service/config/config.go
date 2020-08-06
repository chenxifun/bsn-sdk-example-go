package config

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/astaxie/beego"
)

var (
	fabricConf *config.Config
)

func GetFabricConfig() *config.Config {
	if fabricConf == nil {
		InitFabricConfig()
	}
	return fabricConf
}

func InitFabricConfig() error {
	appConf := beego.AppConfig

	api := appConf.String("bsn_Api")
	userCode := appConf.String("bsn_UserCode")
	appCode := appConf.String("bsn_AppCode")

	mspDir := appConf.String("bsn_MspPath")
	httpsCert := appConf.String("bsn_HttpsCert")
	privK := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcRniHqapLZ4dwFpJ
Zo3ExKJfnRrYlOzHtLgWYEtiOy2hRANCAAQfFo0cjWXm9Fe1F/vKeYZM+5xIGAa8
pgvb1+c+s8bRqw+9xWvSoQv8AuP2TFJe4iTxZJE1tLxHVsREfH0mOH1p
-----END PRIVATE KEY-----`

	conf, err := config.NewConfig(api, userCode, appCode, config.PubK_R1_test, privK, mspDir, httpsCert)
	if err != nil {
		return err
	}

	fabricConf = conf
	return nil
}
