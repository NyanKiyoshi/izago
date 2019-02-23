<div align='center'>
  <h1>Izago</h1>

  <p>
    <a href='https://godoc.org/github.com/NyanKiyoshi/izago'>
      <img src='https://godoc.org/github.com/NyanKiyoshi/izago?status.svg'
           alt='Godoc'/>
    </a>
    <a href='http://izago.rtfd.io/'>
      <img src='https://readthedocs.org/projects/izago/badge/?version=latest' alt='Documentation Status' />
    </a>
  </p>
  <p>
    <a href='https://travis-ci.org/NyanKiyoshi/izago'>
      <img src='https://travis-ci.org/NyanKiyoshi/izago.svg?branch=master'
           alt='Build Status'/>
    </a>
    <a href='https://codecov.io/gh/NyanKiyoshi/izago'>
      <img src='https://codecov.io/gh/NyanKiyoshi/izago/branch/master/graph/badge.svg'
           alt='Code coverage' />
    </a>
    <a href='https://codeclimate.com/github/NyanKiyoshi/izago/maintainability'>
      <img src='https://api.codeclimate.com/v1/badges/cbb26d85cdfd9c19d962/maintainability'
           alt='Maintainability' />
    </a>
  </p>
</div>

## Requirements
- Golang 1.10+ with modules (VGO) support;
- Make sure you correctly set `GOROOT` to your Go installation 
  and `GOPATH` to your Go cache.

## Installation
1. Install dependencies: `go get -v ./...`
1. Compile the bot: `go build -o compiledbot.out main.go`

## Usage
1. Configure the bot (example configuration: [configs/unittests.yaml](configs/unittests.yaml))
1. Optional: set your configuration path using the `CONFIG_PATH` environment variable
1. Run the bot and pass the configuration path 
   (if you didn't set `CONFIG_PATH`): `./compiledbot.out -c configs/yourConfig.yaml`

## Configuration
See the [Godoc reference](https://godoc.org/github.com/NyanKiyoshi/izago/izago/globals#Configuration).

---

<div align='center'>
  <a href='https://godoc.org/gopkg.in/NyanKiyoshi/izago.v0'>
    <sup>Development Godoc</sup>
  </a>
</div>
