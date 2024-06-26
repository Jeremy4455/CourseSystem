<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>选课</title>
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
                <div></div>
                <div class="control" style="padding: 8px">
                    <a class="button is-danger is-small" href="/logout">登出</a>
                </div>
            </div>
        </aside>
        <!-- 右侧内容 -->
        <main class="column is-four-fifths">
            <h1 class="title">学生选课</h1>

            <!-- 搜索表单 -->
            <form action="/student/pick" method="get" style="margin-bottom: 20px;">
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label">课程号</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="courseCode" pattern="[0-9]{6}">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label">课程名</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="courseName">
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
                                <input class="input" type="text" name="courseTeacherId" pattern="[0-9]{6}">
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
                        <label class="label">上课时间</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input" type="text" name="courseTime">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">搜索</button>
                    </div>
                </div>
            </form>


            <!-- 课程列表 -->
            <table class="table is-fullwidth is-hoverable">
                <thead>
                    <tr>
                        <th>课程代码</th>
                        <th>课程名</th>
                        <th>教师号</th>
                        <th>教师名</th>
                        <th>时间</th>
                        <th>教室</th>
                        <th>已选人数</th>
                        <th>容量</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- 使用range指令遍历所有课程 -->
                    {{range .Classes}}
                    <tr>
                        <td>{{.CourseCode}}</td>
                        <td>{{.CourseName}}</td>
                        <td>{{.TeacherId}}</td>
                        <td>{{.Teacher}}</td>
                        <td>{{.Time}}</td>
                        <td>{{.Classroom}}</td>
                        <td>{{.PickedCount}}</td>
                        <td>{{.Capacity}}</td>
                        <td>
                            <form action="/student/pick" method="post">
                                <input type="hidden" name="courseCode" value="{{.CourseCode}}">
                                <input type="hidden" name="courseName" value="{{.CourseName}}">
                                <input type="hidden" name="courseTeacherId" value="{{.TeacherId}}">
                                <input type="hidden" name="CourseTeacherName" value="{{.Teacher}}">
                                <input type="hidden" name="courseSemester" value="{{.Semester}}">
                                <input type="hidden" name="courseTime" value="{{.Time}}">
                                <input type="hidden" name="classroom" value="{{.Classroom}}">
                                <div class="buttons">
                                    <button class="button is-primary" type="submit">选课</button>
                                </div>
                            </form>
                        </td>
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
            <script src="../../../static/js/RetrievePage.js"></script>
        </main>

    </div>
</body>
</html>
