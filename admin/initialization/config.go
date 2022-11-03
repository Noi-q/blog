package init

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()           // 获取当前工作目录
	viper.SetConfigName("application") // 设置读取的文件名
	viper.SetConfigType("yml")         // 设置读取的文件类型
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
