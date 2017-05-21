<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/create" method="post">
            <input type="hidden" name="uid" value="{{.Uid}}" />
            Username: <input type="text" name="username" value="{{.Username}}"> <br>
            Departement: <input type="text" name="departname" value="{{.Departname}}"> <br>
            <input type="reset" value="Reset">
            <input type="submit" value="Insert">
        </form>
    </body>
</html>
