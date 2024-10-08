# OpenTibiaBR - Login Server

[![Version](https://img.shields.io/github/v/release/tibia-oce/login-server)](https://github.com/tibia-oce/login-server/releases/latest)
[![Go](https://img.shields.io/github/go-mod/go-version/tibia-oce/login-server)](https://golang.org/doc/go1.16)
![GitHub repo size](https://img.shields.io/github/repo-size/tibia-oce/login-server)

[![Discord Channel](https://img.shields.io/discord/528117503952551936.svg?style=flat-square&logo=discord)](https://discord.gg/3NxYnyV)
[![GitHub pull request](https://img.shields.io/github/issues-pr/tibia-oce/login-server)](https://github.com/tibia-oce/login-server/pulls)
[![GitHub issues](https://img.shields.io/github/issues/tibia-oce/login-server)](https://github.com/tibia-oce/login-server/issues)


## Project

OpenTibiaBR - Login Server is a free open source login server developed in golang to enable cipclient and [otclient](https://github.com/tibia-oce/otclient) to connect and login to [canary server](https://github.com/tibia-oce/canary).

Current version supports only http login, through `/login` or `/login.php` routes.

The project is fully covered by tests and supports multi-platform build.
Every release is available with multi-platform applications for download.

## Builds
| Platform       | Build        |
| :------------- | :----------: |
| MacOS          | [![MacOS Build](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-macos.yml/badge.svg?branch=main)](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-macos.yml)   |
| Ubuntu         | [![Ubuntu Build](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-ubuntu.yml/badge.svg?branch=main)](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-ubuntu.yml) |
| Windows        | [![Windows Build](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-windows.yml/badge.svg?branch=main)](https://github.com/tibia-oce/login-server/actions/workflows/ci-build-windows.yml) |

[![Workflow](https://github.com/tibia-oce/login-server/actions/workflows/ci-multiplat-release.yml/badge.svg)](https://github.com/tibia-oce/login-server/actions/workflows/ci-multiplat-release.yml)

### Getting **Started**

To run it, simply download the latest release and define your environment variables.
You can set environment type as `dev` if you want to use a `.env` file (store it in the same folder of the login server).

You can also download our docker image and apply the environment variables to your container.

**Enviroment Variables**

|       NAME          |            HOW TO USE                |
| :------------------ | :----------------------------------  |
|`MYSQL_DBNAME`       | `database default database name`     |
|`MYSQL_HOST`         | `database host`                      |
|`MYSQL_PORT`         | `database port`                      |
|`MYSQL_PASS`         | `database password`                  |
|`MYSQL_USER`         | `database username`                  |
|`ENV_LOG_LEVEL`      | `logrus log level for verbose` [ref](https://pkg.go.dev/github.com/sirupsen/logrus#Level)   |
|`LOGIN_IP`           | `login ip address`                   |
|`LOGIN_HTTP_PORT`    | `login http port`                    |
|`LOGIN_GRPC_PORT`    | `login grpc port`                    |
|`RATE_LIMITER_BURST` | `rate limiter same request burst`    |
|`RATE_LIMITER_RATE`  | `rate limit request per sec per user`|
|`SERVER_IP`          | `game server IP address`             |
|`SERVER_LOCATION`    | `game server location`               |
|`SERVER_NAME`        | `game server name`                   |
|`SERVER_PORT`        | `game server game port`              |
|`VOCATIONS`          | `game vocation list csv (a,b,c)`     |

**Tests**  
`go test ./tests -v`

**Build**  
`RUN go build -o TARGET_NAME ./src/`

## Docker
`docker pull tibia-oce/login-server:latest`<br><br>
[![Automation](https://img.shields.io/docker/cloud/automated/tibia-oce/login-server)](https://hub.docker.com/r/tibia-oce/login-server)
[![Image Size](https://img.shields.io/docker/image-size/tibia-oce/login-server)](https://hub.docker.com/r/tibia-oce/login-server/tags?page=1&ordering=last_updated)
![Pulls](https://img.shields.io/docker/pulls/tibia-oce/login-server)
[![Build](https://img.shields.io/docker/cloud/build/tibia-oce/login-server)](https://hub.docker.com/r/tibia-oce/login-server/builds)

## Benchmark
There are a few known login versions available. The most common ones are a python login and login.php, from different websites versions.
We've performed a benchmark (code available in benchmark_test.go) where 1k valid requests were performed to the server, locally.
As you can see in the results below, we got up to 10x faster without decreasing the server availability (99.5% is still pretty good in any global standard).

![image](https://user-images.githubusercontent.com/34237492/118380499-7da2f500-b5e2-11eb-9025-eda180d501df.png)

Also, we performed a benchmark in google cloud, using cloud run and cloud sql database, both with lower possible specifications.
As you can see, we kept an average of 700 requests/s and a good availability, even if with the latency between my computed and cloud services being accounted in this graph.

![image](https://user-images.githubusercontent.com/34237492/118379403-64964600-b5da-11eb-9e11-25c92024986d.png)

Another great aspect is that, comparing with the python login, our docker image is almost 10x smaller (15Mb). 

## gRPC
From version 2.0.0 on, we start using gRPC protocol. 
The HTTP runs on top of the gRPC layer, using a reversed proxy. That lead to a small gain in the availability without any performance loss.

In the gRPC server we got a 10x performance boost, compared to the HTTP benchmarks:

![image](https://user-images.githubusercontent.com/34237492/118568814-e45a1700-b778-11eb-8b79-ddc26dde487c.png)

## Issues

We use the [issue tracker on GitHub](https://github.com/tibia-oce/login-server/issues). Everyone who is watching the repository gets notified by e-mail when there is an activity, so be mindful about comments that add no value (e.g. "+1"). 

We are willing to improve the login server with more features, so feel free to create issues with features requests and ideas, only bug fixes.

If you'd need an issue/feature to be prioritized, you should either do it yourself and submit a pull request, or place a bounty.

## Pull requests

Before [creating a pull request](https://github.com/tibia-oce/login-server/pulls) please keep in mind:

* Set one single scope in your pull request. Focus help us review and things to ship faster. Too many things on the same Pull Request make it harder to review, harder to test and hard to move on.
* Add tests. Pull Requests without tests **won't** be approved.
* Your code must follow go [standard golang format patterns](https://golang.org/doc/effective_go#formatting).
* There are people that doesn't play the game on the official server, so explain your changes to help understand what are you changing and why.
* Avoid opening a Pull Request to just update minor typo or comments. Try attaching those to other PRs with meaningful content.

## Special Thanks

* our partners
* our crew (majesty, gpedro, eduardo dantas, foot, lucas)

## **Sponsors**

If you want to sponsor here, join on discord and send a message for one of our administrators.

## Partners

[![Supported by OTServ Brasil](https://raw.githubusercontent.com/otbr/otserv-brasil/main/otbr.png)](https://forums.otserv.com.br)
