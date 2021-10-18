package pkg

import "errors"

const (
	webwxinit            = "/cgi-bin/mmwebwx-bin/webwxinit"
	webwxstatusnotify    = "/cgi-bin/mmwebwx-bin/webwxstatusnotify"
	webwxsync            = "/cgi-bin/mmwebwx-bin/webwxsync"
	webwxsendmsg         = "/cgi-bin/mmwebwx-bin/webwxsendmsg"
	webwxgetcontact      = "/cgi-bin/mmwebwx-bin/webwxgetcontact"
	webwxsendmsgimg      = "/cgi-bin/mmwebwx-bin/webwxsendmsgimg"
	webwxsendappmsg      = "/cgi-bin/mmwebwx-bin/webwxsendappmsg"
	webwxbatchgetcontact = "/cgi-bin/mmwebwx-bin/webwxbatchgetcontact"
	webwxoplog           = "/cgi-bin/mmwebwx-bin/webwxoplog"
	webwxverifyuser      = "/cgi-bin/mmwebwx-bin/webwxverifyuser"
	synccheck            = "/cgi-bin/mmwebwx-bin/synccheck"
	webwxuploadmedia     = "/cgi-bin/mmwebwx-bin/webwxuploadmedia"
	webwxgetmsgimg       = "/cgi-bin/mmwebwx-bin/webwxgetmsgimg"
	webwxgetvoice        = "/cgi-bin/mmwebwx-bin/webwxgetvoice"
	webwxgetvideo        = "/cgi-bin/mmwebwx-bin/webwxgetvideo"
	webwxlogout          = "/cgi-bin/mmwebwx-bin/webwxlogout"
	webwxgetmedia        = "/cgi-bin/mmwebwx-bin/webwxgetmedia"
	webwxupdatechatroom  = "/cgi-bin/mmwebwx-bin/webwxupdatechatroom"
	webwxrevokemsg       = "/cgi-bin/mmwebwx-bin/webwxrevokemsg"
	webwxcheckupload     = "/cgi-bin/mmwebwx-bin/webwxcheckupload"

	webwxnewloginpage = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage"
	wxjslogin         = "https://login.wx.qq.com/jslogin"
	wxlogin           = "https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login"
	wxqrcode          = "https://login.weixin.qq.com/qrcode/"
	TIANAPIKEY        = "fdb2ce2ffbc97b4faace9f234a555936"
	WECHATNAME        = "糖糖猪"
)

var domainMap = map[string]*WechatDomain{
	"wx.qq.com":       {"https://wx.qq.com", "https://file.wx.qq.com", "https://webpush.wx.qq.com"},
	"wx2.qq.com":      {"https://wx2.qq.com", "https://file.wx2.qq.com", "https://webpush.wx2.qq.com"},
	"wx8.qq.com":      {"https://wx8.qq.com", "https://file.wx8.qq.com", "https://webpush.wx8.qq.com"},
	"web2.wechat.com": {"https://web2.wechat.com", "https://file.web2.wechat.com", "https://webpush.web2.wechat.com"},
	"wechat.com":      {"https://wechat.com", "https://file.web.wechat.com", "https://webpush.web.wechat.com"},
}

func getDomainByHost(host string) (*WechatDomain, error) {
	value, exist := domainMap[host]
	if !exist {
		return nil, errors.New("invalid host")
	}
	return value, nil
}

type WechatDomain struct {
	BaseHost, FileHost, SyncHost string
}
