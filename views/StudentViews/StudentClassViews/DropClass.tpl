<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>退课</title>
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
        .user-info {
            background-color: #eaeaea;
            padding: 10px;
            text-align: center;
            margin-bottom: 20px;
        }
    </style>
</head>
<body>
    <div class="container-full-height columns">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <!-- 用户信息 -->
            <div class="user-info">
                <p>姓名: {{.UserInfo.name}}</p>
                <p>学号: {{.UserInfo.studentID}}</p>
                <p>当前学期: {{.UserInfo.semester}}</p>
            </div>
            <ul class="menu-list" style="text-align: center; padding: 5px">
                <li><a href="/student/pick">学生选课</a></li>
                <li><a href="/student/drop">学生退课</a></li>
                <li><a href="/student/class">已选课程</a></li>
                <li><a href="/student/grade">成绩查询</a></li>
            </ul>
            <!-- 返回上级目录和登出按钮 -->
            <div class="container-button">
                <div class="control" style="padding: 8px">
                    <a class="button is-info is-small"  href="/student">返回</a>
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

            <!-- 已选课程列表 -->
            <div class="selected-courses">
                <h2 class="subtitle">已选课程</h2>
                <table class="table is-fullwidth is-hoverable">
                    <thead>
                        <tr>
                            <th>课程号</th>
                            <th>课程名</th>
                            <th>教师名</th>
                            <th>开课学期</th>
                            <th>上课时间</th>
                            <th>教室</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Classes}}
                        <tr>
                            <td>{{.Course.CourseCode}}</td>
                            <td>{{.Course.Name}}</td>
                            <td>{{.Teacher.Name}}</td>
                            <td>{{.Semester}}</td>
                            <td>{{.ClassTime}}</td>
                            <td>{{.Location}}</td>
                            <td>
                                <form action="/student/drop" method="post">
                                    <input type="hidden" name="courseCode" value="{{.Course.CourseCode}}">
                                    <input type="hidden" name="courseName" value="{{.Course.Name}}">
                                    <input type="hidden" name="courseTeacherId" value="{{.Teacher.TeacherId}}">
                                    <input type="hidden" name="CourseTeacherName" value="{{.Teacher.Name}}">
                                    <input type="hidden" name="courseSemester" value="{{.Semester}}">
                                    <input type="hidden" name="courseTime" value="{{.ClassTime}}">
                                    <input type="hidden" name="classroom" value="{{.Location}}">
                                    <button class="button is-danger is-small" type="submit">退课</button>
                                </form>
                            </td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </main>


    </div>
</body>
</html>
