# collective_intelligence
## install

1. Install golang >= 1.13 . See https://golang.org/doc/install
2. `git clone https://github.com/sudnonk/collective_intelligence.git`
3. `mkdir collective_intelligence/{svgs,json}`

## run

1. Edit `main.go` and decide parameters in Config.
2. `go build` or `go run main.go` (Latter is useful when you tuning parameters.)
3. Run `collective_intelligence` executable file.
4. Results will be generated into `json/l`

If you encounter errors like `git fetch-pack: expected shallow list`, try upgrading git package.
(In my CentOS 7 server, yum provides git 1.8. But git fetch-pack sub-command works only git >= 2.0, so I installed new one via IUS repo.)