<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>增加开课</title>
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
                <li><a href="/admin/class/retrieve">查询开课</a></li>
                <li><a href="/admin/class/create">增加开课</a></li>
                <li><a href="/admin/class/update">更改开课</a></li>
                <li><a href="/admin/class/delete">删除开课</a></li>
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
            <h1 class="title">增加开课</h1>
            <!-- 表单 -->
            <form action="/admin/class/create" method="post">
                <!-- 表单项 -->
                <table class="table is-fullwidth is-hoverable">
                    <thead style="position: sticky; top: 0; z-index: 1;">
                        <tr>
                            <th class="has-text-centered has-text-left">课程号</th>
                            <th class="has-text-centered has-text-left">课程名</th>
                            <th class="has-text-centered has-text-left">教师号</th>
                            <th class="has-text-centered has-text-left">教师名</th>
                            <th class="has-text-centered has-text-left">开课学期</th>
                            <th class="has-text-centered has-text-left">时间</th>
                            <th class="has-text-centered has-text-left">容量</th>
                            <th class="has-text-centered has-text-left">教室</th>
                        </tr>
                    </thead>
                    <tbody>
                        <!-- 表行 -->
                        <tr>
                            <td><input class="input is-small" type="text" id="CourseCode" name="CourseCode" required pattern="[0-9]{6}"></td>
                            <td><input class="input is-small" type="text" id="CourseName" name="CourseName"></td>
                            <td><input class="input is-small" type="text" id="CourseTeacherId" name="CourseTeacherId" required pattern="[0-9]{6}"></td>
                            <td><input class="input is-small" type="text" id="CourseTeacherName" name="CourseTeacherName"></td>
                            <td>
                                <select name="CourseSemester" required>
                                    <option value="">请选择学期</option>
                                    {{range .Semesters}}
                                    <option value="{{.}}">{{.}}</option>
                                    {{end}}
                                </select>
                            </td>
                            <td><input class="input is-small" type="text" id="ClassTime" name="ClassTime" required></td>
                            <td><input class="input is-small" type="text" id="Capacity" name="Capacity" required></td>
                            <td><input class="input is-small" type="text" id="Classroom" name="Classroom" required></td>
                        </tr>
                    </tbody>
                </table>

                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">添加开课</button>
                    </div>
                </div>
            </form>
            <!-- 表格 -->
            <table class="table is-fullwidth is-hoverable">
                <thead style="position: sticky; top: 0; z-index: 1;">
                    <tr>
                        <th class="has-text-centered has-text-left">课程号</th>
                        <th class="has-text-centered has-text-left">课程名</th>
                        <th class="has-text-centered has-text-left">教师号</th>
                        <th class="has-text-centered has-text-left">教师名</th>
                        <th class="has-text-centered has-text-left">开课学期</th>
                        <th class="has-text-centered has-text-left">时间</th>
                        <th class="has-text-centered has-text-left">容量</th>
                        <th class="has-text-centered has-text-left">教室</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- 表行 -->
                    <!-- 使用range指令遍历所有课程 -->
                    {{range .Classes}}
                    <tr>
                        <td>{{.Course.CourseCode}}</td>
                        <td>{{.Course.Name}}</td>
                        <td>{{.Teacher.TeacherId}}</td>
                        <td>{{.Teacher.Name}}</td>
                        <td>{{.Semester}}</td>
                        <td>{{.ClassTime}}</td>
                        <td>{{.Capacity}}</td>
                        <td>{{.Location}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
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
