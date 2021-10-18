package pkg

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)


func ConsoleQrCode(uuid string) {
	login := fmt.Sprintf(wxqrcode+uuid)
	q, _ := qrcode.New(login, qrcode.Low)
	fmt.Println(q.ToString(true))
	fmt.Println("访问下面网址扫描二维码登录\n"+login)
}

// HotLogin 热登录，需要接受一个热存储对象
func HotLogin(bot *openwechat.Bot) error {
	file := "storage.json"
	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage(file)
	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	//if strings.Contains(err.Error(), "not login") {
	//	if _, err = os.Stat(file); ! os.IsNotExist(err) {
	//		os.Remove(file)
	//	}
	//}
	return err
}


func Login(bot *openwechat.Bot) error {
	return bot.Login()
}


