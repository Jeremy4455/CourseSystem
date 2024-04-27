<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>学生退课</title>
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
                <li><a href="/admin/trans/upgrade">课程更新</a></li>
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
            <h1 class="title">学生退课</h1>
            <!-- 输入表单 -->
            <form action="/admin/trans/drop" method="post">
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
                    <div class="field-label is-normal">
                        <label class="label">课程名</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="CourseName" >
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
                        <label class="label">教师名</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="CourseTeacherName">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="Semester">开课学期：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <div class="select">
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
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-danger" type="submit">退课</button>
                    </div>
                </div>
            </form>

            <!-- 退课信息 -->
            {{if .ClassStudent}}
            <div class="notification is-info">
                <p>退课信息：</p>
                <p>课程号: {{.ClassStudent.CourseCode}}</p>
                <p>课程名: {{.ClassStudent.CourseName}}</p>
                <p>学分: {{.ClassStudent.Credit}}</p>
                <p>教师号: {{.ClassStudent.TeacherId}}</p>
                <p>教师名: {{.ClassStudent.TeacherName}}</p>
            </div>
            {{end}}
        </main>

    </div>
</body>
</html>
