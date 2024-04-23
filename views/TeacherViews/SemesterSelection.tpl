<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>选择学期</title>
    <link rel="stylesheet" href="../../static/css/bulma.css">
    <style>
        html, body {
            height: 100%;
        }
        .hero-body {
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100%;
        }
        .title {
            text-align: center;
        }
    </style>
</head>
<body>
    <section class="hero is-fullheight">
        <div class="hero-body">
            <div class="container">
                <h1 class="title">请选择学期：</h1>
                <form action="/choose_course" method="post">
                    <div class="field">
                        <div class="control">
                            <ul class="list">
                                {{range .Semesters}}
                                <li>
                                    <button class="button is-link is-light is-fullwidth" type="submit" name="semester" value="{{.}}">
                                        {{.}}
                                    </button>
                                </li>
                                {{end}}
                            </ul>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </section>
</body>
</html>
