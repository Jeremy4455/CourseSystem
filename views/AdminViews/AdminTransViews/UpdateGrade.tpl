<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>成绩更新</title>
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
            <h1 class="title">成绩更新</h1>

            <!-- 查询表单 -->
            <form action="/admin/trans/update" method="get">
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label">学生号</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="StudentId" required pattern="[0-9]{8}">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label">课程号</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="CourseCode" required pattern="[0-9]{6}">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label">教师号</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="TeacherId" required pattern="[0-9]{6}">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Semester">学期：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <div class="select is-small">
                                    <select name="Semester" required>
                                        <option value="">请选择学期</option>
                                        {{range .Semesters}}
                                        <option value="{{.}}">{{.}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-grouped" style="float: right; margin-top: 20px">
                    <div class="control">
                        <button class="button is-primary" type="submit">查询</button>
                    </div>
                </div>
            </form>

            <!-- 成绩更新表格 -->
            {{ if .ClassStudent }}
            <h2 class="subtitle">学生信息</h2>
            <table class="table is-fullwidth is-hoverable">
                <thead>
                    <tr>
                        <th>学生号</th>
                        <th>学生姓名</th>
                        <th>课程代码</th>
                        <th>课程名称</th>
                        <th>平时分</th>
                        <th>考试分</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>

                        <tr>
                            <form action="/admin/trans/update" method="post">
                            <td>{{ .ClassStudent.StudentId }}</td>
                            <td>{{ .ClassStudent.StudentName }}</td>
                            <td>{{ .ClassStudent.CourseCode }}</td>
                            <td>{{ .ClassStudent.CourseName }}</td>
                            <td>{{ .ClassStudent.Performance }}</td>
                            <td>
                                <input class="input" type="text" name="Score" value="{{ .ClassStudent.Score }}">
                            </td>
                            <td>
                                    <input type="hidden" name="StudentId" value="{{ .ClassStudent.StudentId }}">
                                    <input type="hidden" name="CourseCode" value="{{ .ClassStudent.CourseCode }}">
                                    <input type="hidden" name="TeacherId" value="{{ .ClassStudent.TeacherId }}">
                                    <input type="hidden" name="Semester" value="{{ .ClassStudent.Semester }}">
                                    <button class="button is-primary is-small" type="submit">更新</button>
                            </td>
                            </form>
                        </tr>

                </tbody>
            </table>
            {{ end }}
        </main>
    </div>
</body>
</html>
