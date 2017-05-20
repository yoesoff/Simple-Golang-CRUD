<html>
    <head>
    <title></title>
    </head>
    <body>
        <table>
            <tr>
                <th>
                    ID
                </th>

                <th>
                    Username
                </th>

                <th>
                   Departname
                </th>

                <th>
                   Created
                </th>
                <th>
                    actions
                    </th>
            </tr>
        {{ range  . }}
           <tr>
               <td>{{ .Uid }}</td>
               <td>{{ .Username }}</td>
               <td>{{ .Departname}}</td>
               <td>{{ .Created}}</td>
               <td> 
               
                   <a href="/create">Create</a>
                   <a href="/delete?id={{.Uid}}">delete</a>

               </td>
           </tr>
        {{ end }}
        </table>
        <a href="/create">Create</a>
    </body>
</html>
