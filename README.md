# miniweb

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kemokemo/miniweb)

Tiny web service to serve files. :tada:

## Usage

```sh
miniweb -p {your port} {your content directory}

# ex) miniweb -p 9000 ./my/content/directory
## open 'http://localhost:{your port}' via browser
```

For more detail, please check with the output of `miniweb -h` command.

## How to install

### Homebrew

```sh
brew install kemokemo/tap/miniweb
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install miniweb
```

### Build yourself

```sh
go install github.com/kemokemo/miniweb
```

### Binary

Download from [the release page](https://github.com/kemokemo/miniweb/releases/latest), unpack the archive and put the binary somewhere in your `PATH`.
