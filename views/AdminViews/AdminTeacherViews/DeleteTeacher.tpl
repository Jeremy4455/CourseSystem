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
                <li><a href="/admin/teacher/retrieve">查询教师</a></li>
                <li><a href="/admin/teacher/create">增加教师</a></li>
                <li><a href="/admin/teacher/update">更改教师</a></li>
                <li><a href="/admin/teacher/delete">删除教师</a></li>
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
            <h1 class="title">删除教师</h1>
            <!-- 表单 -->
            <form action="/admin/teacher/delete" method="get">
                <!-- 表单项 -->
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="TeacherId">教师号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="TeacherId" name="TeacherId">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Name">教师名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="Name" name="Name">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="Mobile">手机号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="tel" id="Mobile" name="Mobile" pattern="[0-9]{11}">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Email">邮箱：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="email" id="Email" name="Email">
                            </div>
                        </div>
                    </div>
                </div>

                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">搜索教师</button>
                    </div>
                </div>
            </form>
            <!-- 表格 -->
            <table class="table is-fullwidth is-hoverable">
                <thead style="position: sticky; top: 0; z-index: 1;">
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
                    <!-- 使用range指令遍历所有课程 -->
                    {{range .Teachers}}
                    <tr>
                        <td>{{.TeacherId}}</td>
                        <td>{{.Name}}</td>
                        <td>{{.Mobile}}</td>
                        <td>{{.Email}}</td>
                        <td>
                            <div class="field is-grouped">
                                <div class="control">
                                    <form action="/admin/teacher/delete" method="post">
                                        <input type="hidden" name="TeacherId" value="{{.TeacherId}}">
                                        <button class="button is-danger" type="submit">删除</button>
                                    </form>
                                </div>
                            </div>
                        </td>
                    </tr>
                {{end}}
                </tbody>
            </table>
            <!-- 分页按钮 -->
            <div id="pagination-buttons-container" style="position: fixed; bottom: 20px; right: 20px;">
                <div class="field is-grouped">
                    <div class="control">
                        <button id="prev-page-btn" class="button is-primary" disabled>上一页</button>
                    </div>
                    <div class="control">
                        <button id="next-page-btn" class="button is-primary">下一页</button>
                    </div>
                </div>
            </div>
            <!-- 翻页功能 -->
            <script src="../../../static/js/RetrievePage.js"></script>
        </main>

    </div>
</body>
</html>
