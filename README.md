# Internet topology

[![Go Report
Card](https://goreportcard.com/badge/github.com/lucasdc6/internet-topology)](https://goreportcard.com/report/github.com/lucasdc6/internet-topology)
[![GoDoc](https://godoc.org/github.com/lucasdc6/internet-topology?status.svg)](https://godoc.org/github.com/lucasdc6/internet-topology)


## Compilation

### Requeriments

- make
- go 1.16.2 or later
- Graphviz (optional)
  - Alternative https://dreampuf.github.io/GraphvizOnline

### Steps

```bash
make
# or
make small
```

### Build docker

```bash
make docker
```


## Usage

### binary

Compile or download from actions, then run

```console
./internet-topology --asn=5692 --deep=2 --ix --peers
dot -Tsvg /tmp/internet-topology.dot > unlp.svg
```


### Docker image

```console
docker run --rm -it \
  -v $PWD/data:/data \
  ghcr.io/lucasdc6/internet-topology:latest --asn=5692 --deep=2 --ix --peers -o /data/unlp
dot -Tsvg data/unlp.dot > unlp.svg
```
