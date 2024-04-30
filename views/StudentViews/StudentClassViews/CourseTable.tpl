<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>查看课表</title>
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
        .course-table {
            display: grid;
            margin-top: 30px;
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
        .time-column {
            width: 10%; /* 设置左侧节次信息列的宽度 */
        }
        .course-cell {
            width: 12%; /* 设置课程单元格的宽度 */
        }
        .table {
        text-align: center;
        }
        .table th{
            text-align: center;
        }
        .table td {
            text-align: center;
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
            <h1 class="title">课程表</h1>

            <!-- 课程表部分 -->

            <div class="course-table" id="timetable"></div>

            <!-- JavaScript 渲染课程表 -->
            <script>
                // 获取已选课程数据
                var classes = {{ .Classes }};

                // 在页面加载完成后执行
                document.addEventListener("DOMContentLoaded", function() {
                    renderTimetable();
                });

                // 渲染课程表
                function renderTimetable() {
                    var timetable = document.getElementById("timetable");

                    // 创建表格元素
                    var table = document.createElement("table");
                    table.classList.add("table", "is-bordered", "is-striped");

                    // 创建表头
                    var thead = document.createElement("thead");
                    var headerRow = document.createElement("tr");
                    // 新增一个空白表头单元格作为左侧空列
                    headerRow.appendChild(document.createElement("th"));
                    for (var i = 0; i < 7; i++) { // 一周七天
                        var th = document.createElement("th");
                        th.textContent = "周" + ["一", "二", "三", "四", "五", "六", "日"][i]; // 显示周几
                        headerRow.appendChild(th);
                    }
                    thead.appendChild(headerRow);
                    table.appendChild(thead);

                    // 创建课程表格
                    var tbody = document.createElement("tbody");
                    for (var j = 1; j <= 12; j++) { // 假设一天有 12 节课
                        var row = document.createElement("tr");
                        // 添加左侧节次信息列
                        var timeCell = document.createElement("td");
                        timeCell.textContent =  j ;
                        timeCell.classList.add("time-column"); // 添加自定义类以设置样式
                        row.appendChild(timeCell);
                        for (var k = 0; k < 7; k++) {
                            var cell = document.createElement("td");
                            cell.classList.add("course-cell"); // 添加自定义类以设置样式
                            var course = getCourseForTime(j, k); // 获取该时间段的课程
                            if (course) {
                                cell.textContent = course.Name;
                            }
                            row.appendChild(cell);
                        }
                        tbody.appendChild(row);
                    }
                    table.appendChild(tbody);

                    // 将表格添加到课程表容器中
                    timetable.appendChild(table);
                }


                // 根据时间和日期获取对应的课程
                function getCourseForTime(time, day) {
                    var dayOfWeek = ["一", "二", "三", "四", "五", "六", "日"][day]; // 映射到周几
                    for (var i = 0; i < classes.length; i++) {
                        var classTimes = classes[i].ClassTime.split(" "); // 根据空格分割不同的时间段
                        for (var j = 0; j < classTimes.length; j++) {
                            var classTime = classTimes[j].split(""); // 再次分割每个时间段
                            var dayStart = classTime[0]; // 时间段开始的周几
                            var timeRange = classTime.slice(1).join(""); // 时间段的具体时间范围，将数组合并为字符串
                            if (dayStart === dayOfWeek) {
                                // 检查时间范围是否包含连字符
                                if (timeRange.includes("-")) {
                                    // 如果包含连字符，解析开始和结束时间
                                    var rangeParts = timeRange.split("-");
                                    var start = parseInt(rangeParts[0]);
                                    var end = parseInt(rangeParts[1]);
                                } else {
                                    // 如果不包含连字符，开始和结束时间相同
                                    var start = parseInt(timeRange);
                                    var end = start;
                                }
                                if (time >= start && time <= end) {
                                    return classes[i].Course;
                                }
                            }
                        }
                    }
                    return null;
                }


            </script>

        </main>
    </div>
</body>
</html>
