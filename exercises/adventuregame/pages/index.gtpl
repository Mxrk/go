<html> 
<head>
    <title > Adventure Game </title>
    <h1 class="center" > {{(index . ).Title}} </h1>

    <style> 
    .center {
    margin: auto;
    width: 60%;
    border: 3px solid rgb(19, 128, 179);
    padding: 10px;
    }
    </style>
</head>
<body>
    <div class="center">
            {{ range  $value := .Story}}
            <p>{{$value }}</p>
            {{end}}
    </div>
    <p>
        {{ range  $value := .Options}}

        <div class="center">
            <a href="/{{ $value.Arc  }}"> {{ $value.Text  }}</a>

        </div>
        <p></p>
        {{ end }}
    </p>
</body>
</html>