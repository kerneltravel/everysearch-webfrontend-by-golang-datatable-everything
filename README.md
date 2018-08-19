# everysearch-webfrontend-by-golang-datatable-everything
Everything SDK is an powerful file search engiine.And this project -everysearch-web ,provides single-html-page as frontend for search ,with a customgolang daemon application providing search backend based on Everythin-SDK.

Install:
================================================================================
[get zipped file from release page](https://github.com/kerneltravel/everysearch-webfrontend-by-golang-datatable-everything/releases/)
and unzip .

Usage:
================================================================================
start backend by click&run deamon service file:
```
daemon-backend\everysearch-golang-backend.exe
```

open with browser the search page.
```
user-frontend\index.html 
```
or just click the single-run front application:
```
user-frontend\page-load-by-go-webview-demo.exe
```

Development:
================================================================================
1.on compiling:
need golang ENV,and build with build-backend.bat .
the backend relays on these 3 projects:
```
github.com/zserge/webview
```
and 
```
github.com/gin-gonic/gin
```
and
```
github.com/jof4002/Everything
```

2.More info on Everything-runtime files:
put Everything32.dll OR Everything64.dll  file under the same directory of your exe file.


License
================================================================================
New BSD License

