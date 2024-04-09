<!-- course.tpl -->
<!DOCTYPE html>
<html>
<head>
    <title>Course Management</title>
    <style>
        /* 设置左侧栏样式 */
        .sidebar {
            float: left;
            width: 15%;
            height: 100vh; 
            background-color: #f1f1f1;
            padding: 20px;
        }

        /* 右侧内容样式 */
        .content {
            float: left;
            width: 85%;
            padding: 20px;
        }

        /* 设置按钮样式 */
        .sidebar-button {
            display: block;
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <h1>Course Management System</h1>
    <!-- 左侧栏 -->
    <div class="sidebar">
        <a href="/CourseManagement" class="sidebar-button">课程管理</a>
        <a href="/TeacherManagement" class="sidebar-button">教师管理</a>
        <a href="/StudentManagement" class="sidebar-button">学生管理</a>
    </div>
    
    <!-- 右侧内容 -->
    <div class="content">
        <h1>Course Management System</h1>

        <!-- 原有表单和表格部分，请根据需要保留或调整 -->
        <form action="/CourseList" method="post">
            <label>课程号：</label>
            <input type="text" name="courseNo"><br><br>

            <label>课程名：</label>
            <input type="text" name="courseName"><br><br>

            <label>任课教师：</label>
            <input type="text" name="courseTeacher"><br><br>

            <label>课程时间：</label>
            <input type="text" name="courseTime"><br><br>

            <label>学分：</label>
            <input type="text" name="credit"><br><br>

            <label>上课地点：</label>
            <input type="text" name="classroom"><br><br>

            <button type="submit">查询课程</button>
            <button type="submit">新建课程</button>
        </form>

        <table>
            <tr>
                <th>课程号</th>
                <th>课程名</th>
                <th>课程时间</th>
                <th>学分</th>
                <th>上课地点</th>
                <th>任课教师</th>
                <th>操作</th>
            </tr>
            {{range .Courses}}
            <tr>
                <td>{{.CourseNo}}</td>
                <td>{{.CourseName}}</td>
                <td>{{.CourseTime}}</td>
                <td>{{.Credit}}</td>
                <td>{{.Classroom}}</td>
                <td>{{.Teacher}}</td>
                <td>
                    <form action="/CourseDelete" method="post">
                        <input type="hidden" name="courseId" value="{{.ID}}">
                        <button type="submit">删除</button>
                    </form>
                </td>
            </tr>
            {{end}}
        </table>
    </div>
    
</body>
</html>