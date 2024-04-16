<!DOCTYPE html>
<html lang="en">
<head>
    <title>Login</title>
    <link rel="stylesheet" href="../static/css/bulma.css">
    <style>
        .center-box {
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            width: 500px;
            padding: 20px;
            background-color: #f5f5f5;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
    </style>
</head>
<body>
    <div class="center-box">
        <section class="section">
            <div class="container">
                <h2 class="title">Login</h2>
                <form method="post" action="/login">
                    <div class="field">
                        <label class="label" for="username">Username:</label>
                        <div class="control">
                            <input class="input" type="text" id="username" name="username" required>
                        </div>
                    </div>
                    <div class="field">
                        <label class="label" for="password">Password:</label>
                        <div class="control">
                            <input class="input" type="password" id="password" name="password" required>
                        </div>
                    </div>
                    <div class="field">
                        <div class="control">
                            <button class="button is-primary" type="submit">Login</button>
                        </div>
                    </div>
                </form>
            </div>
        </section>
    </div>
</body>
</html>
