<html>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body {
            background-color: black;
            color: white;
        }
        .container {
            width: max-content;
            padding: 15px 30px;
        }
        pre {
            white-space: pre-wrap;
            word-wrap: break-word;
            font-size: 1.5em;
            font-smooth: auto;
        }
        .input {
            color: red;
        }
    </style>
    <body>
        <div class="container" id="container">
            {{if .Input}}
            <pre class="input" id="input">$ {{.Input}}</pre>
            {{end}}
            <pre class="output" id="output">{{.Output}}</pre>
        </div>
    </body>
</html>