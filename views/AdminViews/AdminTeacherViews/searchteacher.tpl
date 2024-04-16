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
    </style>
</head>
<body>
    <div class="container-full-height columns is-gapless">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <ul class="menu-list">
                <li><a href="/admin/teacher/retrieve">查询教师</a></li>
                <li><a href="/admin/teacher/create">增加教师</a></li>
                <li><a href="/admin/teacher/update">更改教师</a></li>
                <li><a href="/admin/teacher/delete">删除教师</a></li>
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
        <main class="column">
            <h1 class="title">教师管理</h1>
            <!-- 表单 -->
            <form action="/TeacherList" method="post">
                <!-- 表单项 -->
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="teacherId">教师号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="teacherId" name="teacherId">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="teacherName">教师名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="teacherName" name="teacherName">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="phoneNumber">手机号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="phoneNumber" name="phoneNumber">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="email">邮箱：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="email" id="email" name="email">
                            </div>
                        </div>
                    </div>
                </div>

                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">查询教师</button>
                    </div>
                </div>
            </form>

            <!-- 表格 -->
            <table class="table is-fullwidth is-hoverable">
                <thead>
                    <tr>
                        <th class="has-text-centered has-text-left">教师号</th>
                        <th class="has-text-centered has-text-left">教师名</th>
                        <th class="has-text-centered has-text-left">手机号</th>
                        <th class="has-text-centered has-text-left">邮箱</th>
                        <th class="has-text-centered has-text-left">操作</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- 表行 -->
                    <tr>
                        <td>{{.Teacher.TeacherId}}</td>
                        <td>{{.Teacher.Name}}</td>
                        <td>{{.Teacher.Mobile}}</td>
                        <td>{{.Teacher.Email}}</td>
                        <td>
                            <div class="field is-grouped">
                                <div class="control">
                                    <form action="/TeacherDelete" method="post">
                                        <input type="hidden" name="teacherId" value="{{.ID}}">
                                        <button class="button is-danger" type="submit">删除</button>
                                    </form>
                                </div>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </table>
        </main>

    </div>
</body>
</html>
