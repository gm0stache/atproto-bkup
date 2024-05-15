# atproto-ipfs-bkup

util for backing up ATproto repos.

# usage

### general/help

```shell
atp-bkup help
```

### download your ATproto repo

structure:

```shell
atp-bkup dload [ATproto handle] -o [custom output path]
```

example:

```shell
atp-bkup dload mona.lisa.me -o ./nice/dir/mynicefilename.car
```

###

# installation

```shell
go install github.com/gm0stache/atproto-bkup/cmd@latest
```
