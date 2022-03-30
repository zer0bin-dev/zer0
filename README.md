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

### AUR
```
yay -S zer0
```

### From source
```
go install github.com/zer0bin-dev/zer0@v1.0.0
```
(make sure you've [configured PATH correctly](./gopath.md))

### Binaries

https://github.com/zer0bin-dev/zer0/releases/tag/v1.0.0


### Manpage
```
sudo curl -fsSL https://raw.githubusercontent.com/zer0bin-dev/zer0/main/zer0.1.gz -o /usr/share/man/man1/zer0.1.gz && mandb
```

# Technologies used:

<a href="https://go.dev/"><img src="https://github.com/tandpfun/skill-icons/raw/main/icons/GoLang.svg" height=40/></a> <a href="https://aur.archlinux.org/packages/zer0"><img src="https://cdn.discordapp.com/attachments/810799100940255260/956349479320158308/AUR.svg" height=40/></a>
