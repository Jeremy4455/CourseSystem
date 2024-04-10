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
            position: relative; /* 使得按钮的定位相对于 main 元素 */
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
    </style>
</head>
<body>
    <div class="container-full-height columns is-gapless">
        <!-- 左侧栏 -->
        <aside class="column is-one-fifth sidebar">
            <h6 class="menu-label">菜单</h6>
            <ul class="menu-list">
                <li><a href="/CourseManagement">课程管理</a></li>
                <li><a href="/TeacherManagement">教师管理</a></li>
                <li><a href="/StudentManagement">学生管理</a></li>
            </ul>
        </aside>

        <!-- 右侧内容 -->
        <main class="column">
            <h1 class="title">课程管理</h1>
            <!-- 表单 -->
            <form action="/CourseList" method="post">
                <!-- 表单项 -->
                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="courseNo">课程号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="courseNo" name="courseNo">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="teacherId">教师号：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="teacherId" name="teacherId">
                            </div>
                        </div>
                    </div>
                </div>

                <div class="field is-horizontal">
                    <div class="field-label is-normal">
                        <label class="label" for="courseName">课程名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="courseName" name="courseName">
                            </div>
                        </div>
                    </div>
                    <div class="field-label is-normal">
                        <label class="label" for="teacherName">教师名：</label>
                    </div>
                    <div class="field-body">
                        <div class="field">
                            <div class="control">
                                <input class="input is-small" type="text" id="teacherName" name="teacherName">
                            </div>
                        </div>
                    </div>
                </div>

                <!-- 按钮 -->
                <div class="field is-grouped" style="float: right">
                    <div class="control">
                        <button class="button is-primary" type="submit">查询课程</button>
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
                        <th>教师号</th>
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
                        <td>{{.TeacherId}}</td>
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
