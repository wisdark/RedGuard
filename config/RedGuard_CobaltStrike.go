package config

var RedGuardConfig = `[cert]
# User Optional name
DNSName      = *.aliyun.com,manager.channel.aliyun.com,*.acs-internal.aliyuncs.com",*.connect.aliyun.com,aliyun.com,whois.www.net.cn,tianchi-global.com
# Cert User CommonName
CommonName   = *.aliyun.com
# Cert User Locality
Locality     = HangZhou
# Cert User Organization
Organization = Alibaba (China) Technology Co., Ltd.
# Cert User Country
Country      = CN
# Whether to use the certificate you have applied for true/false
HasCert      = true

[proxy]
# key   : Header Host value of the reverse proxy
# value : The actual address forwarded by the reverse proxy
HostTarget    = {"360.net":"http://127.0.0.1:8080","360.com":"https://127.0.0.1:4433"}
# HTTPS Reverse proxy port
Port_HTTPS    = :443
# HTTP Reverse proxy port
Port_HTTP     = :80
# RedGuard interception action: redirect / reset / proxy (Hijack HTTP Response)
drop_action   = redirect
# URL to redirect to
Redirect      = https://360.net
# IP address owning restrictions example:AllowLocation = 山东,上海,杭州 or shanghai,beijing
AllowLocation = *
# Whitelist list example: AllowIP = 172.16.1.1,192.168.1.1
AllowIP       = *
# Limit the  time of requests example: AllowTime = 8:00 - 16:00
AllowTime     = *
# C2 Malleable File Path
MalleableFile = *
# Edge Host Communication Domain
EdgeHost      = *
# Edge Host Proxy Target example: EdgeTarget = 360.com
EdgeTarget    = *
# Customize the header to be deleted example: Keep-Alive,Transfer-Encoding
DelHeader     = *

[SampleFinger]
# HTTP Request Header Field
FieldName   = *
# Sample Finger example:xxxxxx,xxxxxx
FieldFinger = *
`
