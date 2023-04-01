# image-save

![architecture](https://img.shields.io/badge/architecture-amd64%2Carm64-blue)
![os](https://img.shields.io/badge/os-linux%2Cwindows-blue)
[![go-report](https://goreportcard.com/badge/github.com/DockerContainerService/image-save)](https://goreportcard.com/report/github.com/DockerContainerService/image-save)
![contributors](https://img.shields.io/github/contributors/DockerContainerService/image-save)
![size](https://img.shields.io/github/repo-size/DockerContainerService/image-save)
![languages](https://img.shields.io/github/languages/count/DockerContainerService/image-save)
![file](https://img.shields.io/github/directory-file-count/DockerContainerService/image-save)
![used-by](https://img.shields.io/sourcegraph/rrc/github.com/DockerContainerService/image-save)
[![license](https://img.shields.io/github/license/DockerContainerService/image-save)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![release](https://img.shields.io/github/v/release/DockerContainerService/image-save)](https://github.com/DockerContainerService/image-save/releases)
[![download](https://img.shields.io/github/downloads/DockerContainerService/image-save/total.svg)](https://api.github.com/repos/DockerContainerService/image-save/releases)
[![last-release](https://img.shields.io/github/release-date/DockerContainerService/image-save)](https://github.com/DockerContainerService/image-save/releases)

## Features
* Support save docker image to local independent of docker daemon
* Support for reading registry passwords in environment variables ``REGISTRY_PASSWORD``
* Supports multithreading layer download

## Usage
### Install image-save
you can download the latest binary release [here](https://github.com/DockerContainerService/image-save/releases)

### Install from source
```bash
go get github.com/DockerContainerService/image-save
cd ${GOPATH}/github.com/DockerContainerService/image-save
make all
```

### Get usage information
```bash
[root@tencent ~]# ./imsave -h
Save docker image to local without docker daemon
        Complete documentation is available at https://github.com/DockerContainerService/image-save

Usage:
  imsave [image] [flags]

Flags:
      --arch string     The architecture of the image you want to save (default "amd64")
  -d, --debug           Enable debug mode
  -h, --help            help for imsave
  -i, --insecure        Whether the registry is using http
      --os string       The osFilter of the image you want to save
  -o, --output string   Output file
  -p, --passwd string   Password of the registry
  -u, --user string     Username of the registry
  -v, --version         version for imsave
```
### Usage example
```bash
[root@tencent ~]# ./imsave alpine
Using default tag: latest
Using architecture: amd64
[f56be85fc22e]  ... done! [3.37MB in 2.913s; 1.13MB/s]
Output file: alpine_latest.tgz
```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=DockerContainerService/image-save&type=Date)](https://star-history.com/#DockerContainerService/image-save&Date)


