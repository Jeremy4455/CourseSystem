<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Course Management</title>
    <!-- 引入Bulma CSS文件 -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
    <style>
        /* 添加额外的自定义样式 */
        body, html {
            height: 100%;
        }
        .container-full-height {
            height: 100%;
        }
        .sidebar {
            background-color: #f1f1f1;
        }
        .content {
            padding: 20px;
        }
    </style>
</head>
<body>
    <div class="container-full-height columns is-gapless">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <p class="menu-label">菜单</p>
            <ul class="menu-list">
                <li><a href="/CourseManagement">课程管理</a></li>
                <li><a href="/TeacherManagement">教师管理</a></li>
                <li><a href="/StudentManagement">学生管理</a></li>
            </ul>
        </aside>

        <!-- 右侧内容 -->
        <main class="column">
            <h1 class="title" style="text-align: center">Course Management System</h1>
            <!-- 表单 -->
            <form action="/CourseList" method="post">
                <!-- 表单项 -->
                <div class="field">
                    <label class="label" for="courseNo">课程号：</label>
                    <div class="control">
                        <input class="input" type="text" id="courseNo" name="courseNo">
                    </div>
                </div>

                <!-- 其他表单项... -->

                <div class="field is-grouped">
                    <div class="control">
                        <button class="button is-primary" type="submit">查询课程</button>
                    </div>
                    <div class="control">
                        <button class="button is-link" type="submit">新建课程</button>
                    </div>
                </div>
            </form>

            <!-- 表格 -->
            <table class="table is-fullwidth is-hoverable">
                <thead>
                    <tr>
                        <th>课程号</th>
                        <th>课程名</th>
                        <th>课程时间</th>
                        <th>学分</th>
                        <th>上课地点</th>
                        <th>任课教师</th>
                        <th>操作</th>
                    </tr>
                </thead>
                <tbody>
                    <!-- 表行 -->
                    <tr>
                        <td>{{.CourseNo}}</td>
                        <td>{{.CourseName}}</td>
                        <td>{{.CourseTime}}</td>
                        <td>{{.Credit}}</td>
                        <td>{{.Classroom}}</td>
                        <td>{{.Teacher}}</td>
                        <td>
                            <div class="field is-grouped">
                                <div class="control">
                                    <form action="/CourseDelete" method="post">
                                        <input type="hidden" name="courseId" value="{{.ID}}">
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
