# Sobike - 搜索附近的单车

```text
███████╗ ██████╗ ██████╗ ██╗██╗  ██╗███████╗
██╔════╝██╔═══██╗██╔══██╗██║██║ ██╔╝██╔════╝
███████╗██║   ██║██████╔╝██║█████╔╝ █████╗  
╚════██║██║   ██║██╔══██╗██║██╔═██╗ ██╔══╝  
███████║╚██████╔╝██████╔╝██║██║  ██╗███████╗
╚══════╝ ╚═════╝ ╚═════╝ ╚═╝╚═╝  ╚═╝╚══════╝
```

[![Release][1]][2] [![MIT licensed][3]][4]

[1]: https://img.shields.io/badge/release-v0.1-brightgreen.svg
[2]: https://github.com/playniuniu/go-sobike/releases
[3]: https://img.shields.io/dub/l/vibe-d.svg
[4]: LICENSE

### 使用方法

```bash
sobike 你的地址
```

或者

```
sobike -c 城市 你的地址
```

### 开发者

GO Dep

```bash
go get github.com/fatih/color
go get github.com/sirupsen/logrus
go get github.com/urfave/cli
```

Build & Release

```bash
make build
```

**Note:** Install [upx](https://upx.github.io/) before make release

```bash
make relase
```
