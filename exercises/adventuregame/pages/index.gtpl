<html> 
<head>
    <title > Adventure Game </title>
    <h1 class="center" > {{(index . ).Title}} </h1>

    <style> 
    .center {
    margin: auto;
    width: 60%;
    border: 3px solid #73AD21;
    padding: 10px;
    }
    </style>
</head>
<body>
    <p class="center">
        {{(index . ).Story}}
    </p>

    <p> Options: </p>
    <p>

        {{ range  $value := .Options}}
        <li>{{ $value.Text  }}</li>
        <a href="/{{ $value.Arc  }}"> <button>{{ $value.Arc  }}</button></a>
        {{ end }}
    </p>
</body>
</html>