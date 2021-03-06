package cron

import (
	"fmt"
	"github.com/toolkits/net/httplib"
	"github.com/xsar/healthz/g"
	"log"
	"strings"
	"time"
)

func HealthCheck() {
	d := time.Duration(g.Config().Interval) * time.Second
	for {
		healthCheck()
		time.Sleep(d)
	}
}

func healthCheck() {
	urls := g.Config().Urls
	if urls == nil || len(urls) == 0 {
		return
	}

	ctimeout := time.Duration(g.Config().CTimeout) * time.Millisecond
	rwtimeout := time.Duration(g.Config().RWTimeout) * time.Millisecond

	okstrs := g.Config().OkStrs

	for _, url := range urls {
		req := httplib.Get(url).SetTimeout(ctimeout, rwtimeout)
		resp, err := req.String()
		if err != nil {
			// 如果仅仅是超时，直接忽略
			if !strings.Contains(err.Error(), "timeout") {
				Alert(fmt.Sprintf("curl %s fail %s", url, err.Error()))
			}
			continue
		}

		if okstrs == nil || len(okstrs) == 0 {
			// 说明用户不关心response body是否包含特定字符串
			continue
		}

		if InOkStr(resp, okstrs) {
			continue
		}

		Alert(fmt.Sprintf("curl %s respone: %s", url, resp))
	}
}

func Alert(content string) {
	req := httplib.Post(g.Config().Sender)
	req.Param("tos", g.Config().Tos)
	req.Param("content", fmt.Sprintf("%s %s", content, time.Now().Format("2006-01-02 15:04:05")))
	_, err := req.String()
	log.Println(content)
	if err != nil {
		log.Println("alert fail", err)
	}
}

func InOkStr(respone string, strs []string) bool {
	for _, s := range strs {
		if strings.Contains(respone, s) {
			return true
		}
	}
	return false
}
