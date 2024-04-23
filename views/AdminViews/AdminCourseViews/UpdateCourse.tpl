<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>更改课程</title>
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
            <h1 class="title">更改课程</h1>
            <!-- 表格 -->

                <table class="table is-fullwidth is-hoverable">
                    <thead style="position: sticky; top: 0; z-index: 1;">
                        <tr>
                            <th class="has-text-centered has-text-left">课程号</th>
                            <th class="has-text-centered has-text-left">课程名</th>
                            <th class="has-text-centered has-text-left">学院</th>
                            <th class="has-text-centered has-text-left">学分</th>
                            <th class="has-text-centered has-text-left">操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <!-- 表行 -->
                        <!-- 使用range指令遍历所有课程 -->
                        {{range .Courses}}
                        <form action="/admin/course/update" method="post">
                        <tr>
                            <td>{{.CourseCode}}</td>
                            <td>
                                <input type="hidden" name="CourseCode" value="{{.CourseCode}}">
                                <input class="input is-small" type="text" name="Name" value="{{.Name}}">
                            </td>
                            <td>
                                <input class="input is-small" type="text" name="College" value="{{.College}}">
                            </td>
                            <td>
                                <input class="input is-small" type="number" name="Credit" value="{{.Credit}}" min="1" max="5" step="1">
                            </td>
                            <td>
                                <div class="field is-grouped">
                                    <div class="control">
                                        <button class="button is-info is-small" type="submit">保存更改</button>
                                    </div>
                                </div>
                            </td>
                        </tr>
                        </form>
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
            <script src="../../../static/js/UpdatePage.js"></script>
        </main>
    </div>
</body>
</html>