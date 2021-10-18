package main

import (
	"fmt"
	"strings"
	"wechatbot/pkg"

	"github.com/eatmoreapple/openwechat"
)

func main() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			fmt.Println(msg.Content)

			if msg.IsSendByGroup() {
				if strings.Contains(msg.Content, pkg.WECHATNAME) || msg.IsAt {
					if strings.Contains(msg.Content, "古诗词") {
						replyMsg := pkg.JinRiShiCiRenSheng()
						msg.ReplyText(replyMsg)
					} else {
						replyMsg := pkg.TianReboot(msg.Content)
						msg.ReplyText(replyMsg)
					}
					//_, err := msg.SenderInGroup()
					//if err != nil {
					//	fmt.Errorf("Get sender error")
					//}
				}
			}
			if !msg.IsSendByGroup() {
				_, err := msg.Sender()
				if err != nil {
					fmt.Errorf("Get sender error")
				}
				replyMsg := pkg.TianReboot(msg.Content)
				msg.ReplyText(replyMsg)
			}

		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = pkg.ConsoleQrCode

	// 登陆
	if err := pkg.HotLogin(bot); err != nil {
		fmt.Println("HotLogin error: ", err)
		return
	}

	//// 登陆
	//if err := pkg.Login(bot); err != nil {
	//	fmt.Println("Login error: ", err)
	//	return
	//}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println("*********** Friends **********")
	if err != nil {
		return
	}
	fmt.Println(friends)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println("*********** Groups **********")
	if err != nil {
		return
	}
	for _, g := range groups {
		fmt.Println(g.NickName)
		fmt.Println(g.Members())
	}

	fmt.Println(groups.Count())
	//seaGrps := groups.SearchByNickName(4, "上海")
	//fmt.Println(seaGrps)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	_ = bot.Block()
}
