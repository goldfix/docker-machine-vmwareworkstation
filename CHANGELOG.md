# Changelog

## v2.0.2

IMPROVEMENTS:

* Downgrade `.vmx` file to VmWare 12
* Refactoring on .vmx file
* fix creation mount folder
* build procedure:
  * removed "makefile" and substituted with [Task](https://taskfile.org/)
  * removed [Dep](https://golang.github.io/dep/) and substituted with [Go-Modules](https://github.com/golang/go/wiki/Modules)
* is possible use last version of [boot2docker](https://github.com/boot2docker/boot2docker/) ISO

## v2.0.1

IMPROVEMENTS:

* Update `.vmx` file to VmWare 15
* Removed option: `--vmwareworkstation-share-compat`
* Default VM setting changed to:
  * memory: 4GB
  * CPU: 2
  * Size Virtual Disk: 40GB
* Using iso [boot2docker v18.09.1-rc1](https://github.com/boot2docker/boot2docker/releases/tag/v18.09.1-rc1)

## v2.0.0

IMPROVEMENTS:

* Update github.com/docker/machine to `0.14.0`

---

## v1.1.0

IMPROVEMENTS:

* Updated docker-machine deps
* Updated testImport deps
* Fixes #32 (Unable to complete build when .docker folder is on a symlink)

## v1.0.10

IMPROVEMENTS:

* Mount on both /c/Users and /Users for compatibility
* Use glide for dependancy vendoring
* Update documentation

## v1.0.0

IMPROVEMENTS:

* Add appveyor auto builds
* Initial basic tests

Notes: Updated license files/notices

## v0.9.0

IMPROVEMENTS:

* Windows: dhcp lease file is automatically discoved and not attached to a fixed location
* Update files so it can BUILD against latest docker/machine
* Mimic updates from docker/machine PR docker/machine#2732 and docker/machine#2721

## v0.8.0 (2015-11-17)

Initial public release
