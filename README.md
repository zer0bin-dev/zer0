<div align="center">
    <img src="https://raw.githubusercontent.com/zer0bin-dev/.github/main/zer0bin-client.svg" height="150px"/>
	<br>
    <img src="https://raw.githubusercontent.com/zer0bin-dev/.github/main/zer0bin-client-rainbow.svg" height="100"/>
	<br>
    CLI client for zer0bin
    <br>
</div>

# Usage:

```
$ echo "hello world" | zer0bin
$ cat /var/logs/prog.log | zer0bin # Text only
$ echo "other instance" | zer0bin --instance https://stepbro.voring.me
$ cat document.md | zer0bin --md # Markdown mode
```

# Download:

Soon:tm:

For now...
```
git clone https://github.com/zer0bin-dev/client zer0bin
cd zer0bin
go build -o zer0bin
sudo mv ./zer0bin /usr/bin
```

# Technologies used:

<a href="https://www.rust-lang.org/"><img src="https://github.com/tandpfun/skill-icons/raw/main/icons/GoLang.svg" height=40/></a>
