# yunpian
云片网 GO SDK


##使用接口前要先注册APIKEY
``` go
RegisterKey(yourapikey)
```

##发送手机短信验证码
``` go
Sms_Send(SMSSendInfo{Mobile: "13250061802", Text: "【垣创科技】您的验证码是970702"})
```

##发送语音验证码
``` go
Voice_Send(VoiceSendInfo{Mobile: "13250061802", Text: "970702"})
```