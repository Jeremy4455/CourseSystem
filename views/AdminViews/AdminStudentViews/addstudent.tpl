<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Course Management</title>
    <!-- 引入Bulma CSS文件 -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
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
            position: relative; /* 使得按钮的定位相对于 main 元素 */
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
    </style>
</head>
<body>
    <div class="container-full-height columns is-gapless">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <ul class="menu-list">
                <li><a href="/admin/student/retrieve">查询学生</a></li>
                <li><a href="/admin/student/create">增加学生</a></li>
                <li><a href="/admin/student/update">更改学生</a></li>
                <li><a href="/admin/student/delete">删除学生</a></li>
            </ul>
            <!-- 返回上级目录按钮 -->
            <div class="field is-grouped" style="margin-top: 20px;">
                <div class="control">
                    <a class="button is-info is-small" href="/admin">返回上级目录</a>
                </div>
            </div>
            <!-- 登出按钮 -->
            <div class="field" style="margin-top: 20px;">
                <div class="control">
                    <a class="button is-danger is-small" href="/logout">登出</a>
                </div>
            </div>
        </aside>



    </div>
</body>
</html>
