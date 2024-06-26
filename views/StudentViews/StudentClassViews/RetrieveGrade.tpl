<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>成绩查询</title>
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
                <p>学号: {{.UserInfo.userId}}</p>
                <p>当前学期: {{.UserInfo.semester}}</p>
            </div>
            <ul class="menu-list" style="text-align: center; padding: 5px">
                <li><a href="/student/pick">学生选课</a></li>
                <li><a href="/student/drop">学生退课</a></li>
                <li><a href="/student/class">已选课程</a></li>
                <li><a href="/student/table">查看课表</a></li>
                <li><a href="/student/grade">成绩查询</a></li>
            </ul>
            <!-- 返回上级目录和登出按钮 -->
            <div class="container-button">
                <div class="control" style="padding: 8px">
                    <a class="button is-info is-small"  href="/student/index">返回</a>
                </div>
                <div class="control" style="padding: 8px">
                    <a class="button is-danger is-small" href="/logout">登出</a>
                </div>
            </div>
        </aside>

        <!-- 右侧内容 -->
        <main class="column is-four-fifths">
            <h1 class="title">成绩查询</h1>
            <!-- 成绩列表 -->
            <table class="table is-fullwidth is-hoverable">
                <thead>
                    <tr>
                        <th>课程号</th>
                        <th>课程名</th>
                        <th>授课教师</th>
                        <th>平时成绩</th>
                        <th>期末成绩</th>
                        <th>最终成绩</th>
                        <th>绩点</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Classes}}
                    <tr>
                        <td>{{.CourseCode}}</td>
                        <td>{{.CourseName}}</td>
                        <td>{{.Teacher}}</td>
                        <td>{{.Performance}}</td>
                        <td>{{.Score}}</td>
                        <td>{{.Grade}}</td>
                        <td>{{.Level}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </main>
    </div>
</body>
</html>
