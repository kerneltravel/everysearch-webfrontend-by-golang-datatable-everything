REM 编译得到无控制台界面，后台服务：
go build -ldflags "-H windowsgui" everysearch-golang-backend.go

REM 编译得到有控制台界面的，非隐藏式服务
REM  go build everysearch-golang-backend.go