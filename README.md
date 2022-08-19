# ShotTower
<div align="center">

[![golangci-lint](https://github.com/DblK/shottower/actions/workflows/golangci-lint.yml/badge.svg?branch=master)](https://github.com/DblK/shottower/actions/workflows/golangci-lint.yml)
[![ginkgo](https://github.com/DblK/shottower/actions/workflows/ginkgo.yml/badge.svg?branch=master)](https://github.com/DblK/shottower/actions/workflows/ginkgo.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/DblK/shottower.svg)](https://github.com/DblK/shottower)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/DblK/shottower)
[![GoReportCard](https://goreportcard.com/badge/github.com/DblK/shottower)](https://goreportcard.com/report/github.com/DblK/shottower)
[![GitHub release](https://img.shields.io/github/release/DblK/shottower.svg)](https://GitHub.com/DblK/shottower/releases/)
</div>

An open source, self-hosted implementation of the Shotstack backend server.

Join our [Discord](https://discord.gg/GCXCNHRC) server for a chat.

## What is ShotStack

ShotStack is a video editing API that allow to build dynamic Video Applications at Scale.

Everything in ShotStack is Open Source, except the Dashboard and the backend server.

The backend server translate JSON to FFMPEG commands and serve the result.

## Design goal

`shottower` aims to implement a self-hosted, open source alternative to the Shotstack backend server. `shottower` has a narrower scope and are not built for heavy transcode (but pull request are welcome!).

`shottower` uses terms that maps to Shotstack's API, consult the glossary for explanations.
## Support

If you like `shottower` and find it useful, there is a sponsorship and donation buttons available in the repo.

If you would like to sponsor features, bugs or prioritization, reach out to one of the maintainers.

## Features (Basic + Exclusive 😎)

* "Basic" (`Track`, `Clip`, `VideoAsset`, ...) support of Shotstack's features (See bellow for progress)
    * Use only the `stage` endpoint value until all features are implemented (See https://github.com/DblK/shottower/issues/1 for multiple endpoint handling)
* 😎 Possible to burn subtitle into video clip
* 😎 Allow to use local file from `url` filed (`file:///Users/dblk/clips/my_asset`)
* 😎 Add an endpoint `/dl/{version}/renders/:id` to download renders (instead of cdn/s3)
* 😎 Add other values for resolution (`360`, `480`, `540`, `720`) all with default `25 fps`.
* 😎 Add other values for output quality (`highest`, `lowest`).
* [`Planned`] Allow to use ftp file from `url` filed (`ftp://user:password@dblk.org/mypath/my_asset`)
* [`Planned`] Add destination to Youtube

### Shotstack implementation progress

In the following matrix, there is only components than their implementation have been started.  
At the end of the road this section should either disappear or be full of `Yes` 😇.

| Component | Property | Implemented | Comment |
| ------------------- | -------- | ----------- | --------- |
| Timeline | soundtrack | Not yet |  |
| Timeline | background | Yes ✅ |  |
| Timeline | fonts | Not yet |  |
| Timeline | tracks | Yes ✅ |  |
| Timeline | cache | Yes ✅ |  |
| Track ✅ | all ✅ | Yes ✅ |  |
| Clip | asset | Partial 🛠 | Only `VideoAsset` are started |
| Clip | start | Yes ✅ |  |
| Clip | length | Yes ✅|  |
| Clip | fit | Not yet |  |
| Clip | scale | Yes ✅|  |
| Clip | position | Yes ✅ |  |
| Clip | offset | Not yet |  |
| Clip | transition | Not yet |  |
| Clip | effect | Not yet |  |
| Clip | filter | Not yet |  |
| Clip | opacity | Not yet |  |
| Clip | transform | Not yet |  |
| Clip [`VideoAsset`] | all ✅ | Yes ✅ |  |
| Clip [`ImageAsset`] | all ✅ | Yes ✅ |  |
| Clip [`TitleAsset`] | all | Not yet |  |
| Clip [`HTMLAsset`] | all | Not yet |  |
| Clip [`AudioAsset`] | src | Partial 🛠 | Download asset only |
| Clip [`AudioAsset`] | trim | Not yet |  |
| Clip [`AudioAsset`] | volume | Not yet |  |
| Clip [`AudioAsset`] | effect | Not yet |  |
| Clip [`LumaAsset`] | src | Partial 🛠 | Download asset only |
| Clip [`LumaAsset`] | trim | Not yet |  |
| Output | format | Partial 🛠 | Only `mp4` at the moment  |
| Output | resolution | Yes ✅ |  |
| Output | aspectRatio | Not yet |  |
| Output | size | Yes ✅ |  |
| Output | fps | Yes ✅ |  |
| Output | scaleTo | Not yet |  |
| Output | quality | Yes ✅ |  |
| Output | repeat | Not yet |  |
| Output | range | Not yet |  |
| Output | poster | Not yet |  |
| Output | thumbnail | Not yet |  |
| Output | destinations | Not yet |  |

#### Endpoint implementation

There are several endpoints that are available but not all of them are implemented.

| Category | Endpoint Description | Status | Comment |
| -------- | -------------------- | ------ | ------- |
| Edit     | Render asset | Yes ✅ | |
| Edit     | Get Render Status | Partial 🛠 | Missing real `owner` |
| Edit     | Create Template | Not yet | |
| Edit     | List Template | Not yet | |
| Edit     | Retrieve Template | Not yet | |
| Edit     | Update Template | Not yet | |
| Edit     | Delete Template | Not yet | |
| Edit     | Render Template | Not yet | |
| Edit     | Inspect Media | Not yet | |
| Serve    | Get Asset | Partial 🛠 | |
| Serve    | Delete Asset | Not yet | |
| Serve    | Get Asset by Render ID | Not yet | |
## Running shottower

Please have a look at the documentation under `docs/`.

## Disclaimer

1. We have nothing to do with shotstack
2. The purpose of ShotTower is maintaining a working, self-hosted Shotstack api compatible backend.

## Contributing

To contribute to shottower you would need the latest version of Go.

### Code style

To ensure we have some consistency with contributions, this project has adopted linting and style/formatting rules:

The **Go** code is linted with [`golangci-lint`](https://golangci-lint.run).

Check out the `.golangci.yml` to see the specific configuration.

### Commit message

To ensure, the changelog are well generated, please use prefix in your commit message as follow:
- `feat`, `features`, `feature`: For new features
- `fix`: For any fix
- `doc`, `docs`: For any update in documentation (README, etc...)

And if you `really` need to commit something that is not working `wip` ou `test` are your friends.

### Install development tools
* Go
* FFMPEG (v5 or up)
* (optional) I recommend using [gow](https://github.com/mitranim/gow)

### Testing and building

To run the tests (solo execution):
```golang
ginkgo -r --randomize-all --randomize-suites --race --trace -cover
```

To run the tests during development:
```golang
ginkgo watch -r --randomize-all --race --trace
```

To build the program:
```golang
go build
```

### Running the server

To run the server, follow these simple steps:

```
go run main.go
```

To run the server in a docker container
```
docker build --network=host -t shottower .
```

Once image is built use
```
docker run --rm -it shottower
```

## Contributors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/DblK>
            <img src=https://avatars.githubusercontent.com/u/832617?v=4 width="100;"  alt=Rémy Boulanouar/>
            <br />
            <sub style="font-size:14px"><b>Rémy Boulanouar</b></sub>
        </a>
    </td>
</tr>
</table>
