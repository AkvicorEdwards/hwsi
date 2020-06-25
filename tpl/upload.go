package tpl

const upload string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
<form enctype="multipart/form-data" action="upload" method="post">
    <label>
        password:
        <input type="password" name="password"/>
    </label>
    <br />
    <br />
    <input type="file" name="filename"/>
    <input type="submit" value="upload"/>
</form>
</body>
</html>`