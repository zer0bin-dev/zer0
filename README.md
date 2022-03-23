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
$ echo "hello world" | zer0
$ cat /var/logs/prog.log | zer0 # Text only
$ echo "other instance" | zer0 --instance https://stepbro.voring.me
$ cat document.md | zer0 --md # Markdown mode
```

# Download:

Soon:tm:

For now...
```
git clone https://github.com/zer0bin-dev/zer0
cd zer0
go build
sudo mv ./zer0 /usr/bin
sudo mv ./zer0.1.gz /usr/share/man/man1/zer0.1.gz
mandb
```

# Technologies used:

<a href="https://www.rust-lang.org/"><img src="https://github.com/tandpfun/skill-icons/raw/main/icons/GoLang.svg" height=40/></a>
