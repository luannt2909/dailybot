package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"time"
)

/*
	{
	    "name": "spaces/AAAA_4XQBbM/messages/vNLlCgVtqWg.vNLlCgVtqWg",
	    "sender": {
	        "name": "users/114022495153014004089",
	        "displayName": "dailybot",
	        "avatarUrl": "",
	        "email": "",
	        "domainId": "",
	        "type": "BOT",
	        "isAnonymous": false
	    },
	    "text": "HELLO",
	    "cards": [],
	    "cardsV2": [],
	    "previewText": "",
	    "annotations": [],
	    "thread": {
	        "name": "spaces/AAAA_4XQBbM/threads/vNLlCgVtqWg",
	        "retentionSettings": {
	            "state": "PERMANENT",
	            "expiryTimestamp": "0"
	        },
	        "threadKey": ""
	    },
	    "space": {
	        "name": "spaces/AAAA_4XQBbM",
	        "type": "ROOM",
	        "singleUserBotDm": false,
	        "threaded": false,
	        "displayName": "Daily Task report",
	        "spaceThreadingState": "THREADED_MESSAGES",
	        "legacyGroupChat": false
	    },
	    "fallbackText": "",
	    "argumentText": "HELLO",
	    "attachment": [],
	    "threadReply": false,
	    "retentionSettings": {
	        "state": "PERMANENT",
	        "expiryTimestamp": "0"
	    },
	    "clientAssignedMessageId": "",
	    "createTime": "2023-01-17T11:30:57.778655Z"
	}
*/
var webhook = "https://chat.googleapis.com/v1/spaces/AAAA_4XQBbM/messages?key=AIzaSyDdI0hCZtE6vySjMm-WEfRq3CPzqKqqsHI&token=345O9nUjumKOP_qj8S8ttiyACqXCT5h7KzgkNV9BBxE%3D"

const remindDailyTaskSchedule = "30 2 * * 1,2,3,4,5"
const reportDailyTaskSchedule = "0 10 * * 1,2,3,4,5"

func main() {
	app := fx.New(fx.Invoke(setupCronJob))
	app.Run()
}

func setupCronJob(lc fx.Lifecycle) {
	//pushMsgReportDailyTask()
	t := time.Now().String()
	fmt.Println("current time: ", t)
	cr := cron.New()
	cr.AddFunc(remindDailyTaskSchedule, func() {
		pushMsgReportDailyTask()
	})
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			cr.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cr.Stop()
			return nil
		},
	})
}

func reportDailyTask() {
	client := resty.New()
	message := buildPushMessage()
	body := map[string]interface{}{
		"text": message,
	}
	rsp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(webhook)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.String())
}

func pushMsgReportDailyTask() {
	client := resty.New()
	message := buildPushMessage()
	body := map[string]interface{}{
		"text": message,
	}
	rsp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(webhook)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.String())
}

func buildPushMessage() string {
	today := time.Now().Format("02/01/2006")
	fmt.Println(today)
	return fmt.Sprintf(
		`=== *(%s)* - *Update c??ng vi???c ng??y h??m nay n??o c??c c???u ??i <3* ===
1. H??m qua ???? l??m nh???ng g???
2. H??m nay d??? ?????nh s??? l??m g???
3. C?? kh?? kh??n g?? trong c??ng vi???c kh??ng?
<users/all> _Reply tin nh???n n??y gi??p m??nh nh?? <3_
`, today)
}
