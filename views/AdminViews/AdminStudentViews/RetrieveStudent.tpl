<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>查询学生</title>
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
                <li><a href="/admin/student/retrieve">查询学生</a></li>
                <li><a href="/admin/student/create">增加学生</a></li>
                <li><a href="/admin/student/update">更改学生</a></li>
                <li><a href="/admin/student/delete">删除学生</a></li>
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
            <h1 class="title">查询学生</h1>
            <!-- 表单 -->
            <form action="/admin/student/retrieve" method="post">
                <!-- 表单项 -->
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="StudentId">学号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="StudentId" name="StudentId">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Name">姓名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="Name" name="Name">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Class">班级：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="Class" name="Class">
                            </div>
                        </div>
                    </div>
                </div>


                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">查询学生</button>
                    </div>
                </div>
            </form>
            <!-- 表格 -->
            <table class="table is-fullwidth is-hoverable">
                <thead style="position: sticky; top: 0; z-index: 1;">
                    <tr>
                        <th class="has-text-centered has-text-left">学号</th>
                        <th class="has-text-centered has-text-left">姓名</th>
                        <th class="has-text-centered has-text-left">班级</th>
                        <th class="has-text-centered has-text-left">绩点</th>
                        <th class="has-text-centered has-text-left">操作</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- 表行 -->
                    <!-- 使用range指令遍历所有课程 -->
                    {{range .Students}}
                    <tr>
                        <td>{{.StudentId}}</td>
                        <td>{{.Name}}</td>
                        <td>{{.Class}}</td>
                        <td>{{.Grade}}</td>
                        <td>
                            <div class="field is-grouped">
                                <div class="control">
                                    <form action="/admin/student/update" method="get">
                                        <input type="hidden" name="StudentId" value="{{.StudentId}}">
                                        <button class="button is-info" type="submit">更改</button>
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
