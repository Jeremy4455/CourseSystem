<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>增加课程</title>
    <!-- 引入Bulma CSS文件 -->
    <link rel="stylesheet" href="/static/css/bulma.css">
    <link rel="stylesheet" href="/static/css/message-box.css">
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
                <li><a href="/admin/course/retrieve">查询课程</a></li>
                <li><a href="/admin/course/create">增加课程</a></li>
                <li><a href="/admin/course/update">更改课程</a></li>
                <li><a href="/admin/course/delete">删除课程</a></li>
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
            <h1 class="title">增加课程</h1>
            <!-- 表单 -->
            <form action="/admin/course/create" method="post">
                <!-- 表单项 -->
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="CourseCode">课程号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="CourseCode" name="CourseCode" required pattern="[0-9]{6}">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Name">课程名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="Name" name="Name" required>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="College">学院：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="College" name="College" required>
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Credit">学分：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="Credit" name="Credit" required min="1" max="5" step="1">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-horizontal">
                    <div class="field-label">
                        <label class="label" for="Proportion">占比：</label>
                    </div>
                    <div class="field-body">
                        <div class="field is-grouped">
                            <div class="control">
                                <input class="input is-small" type="text" id="Proportion" name="Proportion" min="0" max="1" step="0.1">
                            </div>
                        </div>
                    </div>
                </div>
                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">添加课程</button>
                    </div>
                </div>
            </form>

            {{ if .json.message }}
            <div id="success-notification" class="message-box is-success">
              <div class="message-header">
                <p>Success</p>
                <button class="delete" aria-label="delete" onclick="closeNotification('success-notification')"></button>
              </div>
              <div class="message-body">
                  {{.json.message}}
              </div>
            </div>
            {{ end }}

            {{ if .json.error }}
            <div id="error-notification" class="message-box is-danger">
              <div class="message-header">
                <p>Error</p>
                <button class="delete" aria-label="delete" onclick="closeNotification('error-notification')"></button>
              </div>
              <div class="message-body">
                  {{.json.error}}
              </div>
            </div>
            {{ end }}
            <script src="/static/js/closeNotification.js"></script>
        </main>
    </div>
</body>
</html>
