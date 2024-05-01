<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
   
</head>
<body>
    <h1>Razor Box: Portscan em Go</h1>
    <p>Esta é uma ferramenta simples escrita em Go para realizar scans de portas em um host específico.</p>
    
<h2>Instalação</h2>
<p>Antes de usar esta ferramenta, certifique-se de ter o Go instalado em seu sistema.</p>
<pre><code>go get github.com/gookit/color</code></pre>
<pre><code>go get github.com/gosuri/uiprogress</code></pre>

<h2>Uso</h2>
<p>Para usar a ferramenta, basta executar o seguinte comando:</p>
<pre><code>go run main.go</code></pre>
<p>A ferramenta solicitará o host alvo, o range de portas e o intervalo de tempo. Em seguida, iniciará o scan e mostrará os resultados.</p>

<h2>Compilação para Linux</h2>
<p>Para compilar o binário para Linux, use o seguinte comando:</p>
<pre><code>GOOS=linux GOARCH=amd64 go build -o razorbox-linux main.go</code></pre>

<h2>Compilação para Windows</h2>
<p>Para compilar o binário para Windows, use o seguinte comando:</p>
<pre><code>GOOS=windows GOARCH=amd64 go build -o razorbox-windows.exe main.go</code></pre>

<h2>Exemplo</h2>
<pre><code>Digite o host alvo: exemplo.com
Digite o range de portas (exemplo: 1-1000): 1-100
Escolha o intervalo (segundos): 1
Iniciando scan de portas no alvo: exemplo.com...
Porta 80 (HTTP) está aberta
Porta 443 (HTTPS) está aberta
Finalizado!</code></pre>

<h2>Contribuindo</h2>
<p>Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.</p>
</body>
</html>
