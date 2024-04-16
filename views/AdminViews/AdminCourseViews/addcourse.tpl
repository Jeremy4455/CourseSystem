<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Course Management</title>
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
        .center-box {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
        }
        .error-alert {
            position: fixed;
            top: 20px;
            left: 50%;
            transform: translateX(-50%);
            width: 60%;
            z-index: 1000;
            display: none;
        }
    </style>
</head>
<body>
    <div class="container-full-height columns is-gapless">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <ul class="menu-list">
                <li><a href="/admin/course/retrieve">查询课程</a></li>
                <li><a href="/admin/course/create">增加课程</a></li>
                <li><a href="/admin/course/update">更改课程</a></li>
                <li><a href="/admin/course/delete">删除课程</a></li>
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

        <!-- 右侧内容 -->
        <main class="column is-four-fifths">
            <div class="container">
                <div class="content center-box"> <!-- 修改这里 -->
                    <div style="width: 100%;"> <!-- 新增一个盒子 -->
                        <h1 class="title">添加课程</h1>
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
                                            <input class="input is-small" type="text" id="CourseCode" name="CourseCode" required>
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
                                            <input class="input is-small" type="text" id="Credit" name="Credit" required>
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
                    </div>
                    <script>
                        document.getElementById('courseForm').addEventListener('submit', function(event) {
                            var courseCode = document.getElementById('CourseCode').value;
                            var name = document.getElementById('Name').value;
                            var college = document.getElementById('College').value;
                            var credit = document.getElementById('Credit').value;

                            // 简单验证，确保所有字段都填写了
                            if (!courseCode || !name || !college || !credit) {
                                alert('所有字段都必须填写');
                                event.preventDefault(); // 阻止表单提交
                                return;
                            }

                            // 验证学分是否为数字
                            if (isNaN(credit)) {
                                alert('学分必须是一个数字');
                                event.preventDefault(); // 阻止表单提交
                                return;
                            }
                        });
                    </script>
                </div>
            </div>
        </main>
    </div>
</body>
</html>
