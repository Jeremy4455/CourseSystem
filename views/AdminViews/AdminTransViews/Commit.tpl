<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>事务处理</title>
    <!-- 引入Bulma CSS文件 -->
    <link rel="stylesheet" href="../../../static/css/bulma.css">
    <style>
        /* 添加额外的自定义样式 */
        body, html {
            height: 100%;
            font-family: Arial, sans-serif;
        }
        .container-full-height {
            height: 100%;
        }
        .sidebar {
            background-color: #f1f1f1;
            padding: 20px;
        }
        .menu-label{
            text-align: center;
            font-size: 20px;
            margin-bottom: 20px;
        }
        main {
            padding: 20px;
            position: relative;
            margin-top: 20px;
        }
        .title {
            text-align: center;
            margin-bottom: 30px;
            font-size: 36px;
        }
        form {
            margin-bottom: 30px;
        }
        table {
            margin-top: 20px;
        }
        .container-button {
            margin-top: 20px;
            display: flex;
            align-items: center;
            flex-direction: column;
            justify-content: center;
            padding: 20px;
        }
    </style>
</head>
<body>
    <div class="container-full-height columns">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <ul class="menu-list" style="text-align: center; padding: 5px">
                <li><a href="/admin/trans/pick">学生选课</a></li>
                <li><a href="/admin/trans/drop">学生退课</a></li>
                <li><a href="/admin/trans/update">成绩更新</a></li>
                <li><a href="/admin/trans/commit">事务处理</a></li>
            </ul>
            <!-- 返回上级目录和登出按钮 -->
            <div class="container-button">
                <div class="control" style="padding: 8px">
                    <a class="button is-info is-small"  href="/admin">返回</a>
                </div>
                <div></div>
                <div class="control" style="padding: 8px">
                    <a class="button is-danger is-small" href="/logout">登出</a>
                </div>
            </div>
        </aside>
        <!-- 右侧内容 -->
        <main class="column is-four-fifths">
            <h1 class="title">事务处理</h1>

            <!-- 按钮组 -->
            <div class="buttons are-medium is-flex is-justify-content-space-between">
                <form action="/admin/trans/commit" method="post">
                    <button class="button is-info is-rounded" type="submit" name="Trans" value="newSemester">新增学期</button>
                </form>
                <form action="/admin/trans/commit" method="post">
                    <button class="button is-info is-rounded" type="submit" name="Trans" value="begin">开始选课</button>
                </form>
                <form action="/admin/trans/commit" method="post">
                    <button class="button is-info is-rounded" type="submit" name="Trans" value="end">结束选课</button>
                </form>
                <form action="/admin/trans/commit" method="post">
                    <button class="button is-info is-rounded" type="submit" name="Trans" value="done">敲定课程</button>
                </form>
                <form action="/admin/trans/commit" method="post">
                    <button class="button is-info is-rounded" type="submit" name="Trans" value="archive">成绩归档</button>
                </form>
            </div>
        </main>

    </div>
</body>
</html>
