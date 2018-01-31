<html>
<head>
    <title>
        登录
    </title>
</head>
<body>

<form action="login?username=astaxie" method="post">
    <input type="checkbox" name="interest" value="football">足球
    <input type="checkbox" name="interest" value="basketball">篮球
    <input type="checkbox" name="interest" value="tennis">网球
    <br/>
    用户名：<input type="text" name="username"/>
    <br/>
    密码：<input type="password" name="password"/>
    <br/>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登录"/>
</form>
</body>
</html>