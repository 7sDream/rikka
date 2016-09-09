# Rikka - A simple image share system

![][badge-version-img] ![][badge-info-img]

[中文版][readme-zh]

Rikka is written in Golang primarily, and provide Docker image.

Rikka image has been published to [DockerHub][image-in-dockerhub], just try it!

Badges above shows latest version and size of Rikka image.

## Introduction

Rikka（`りっか` in Japanese, sound like `/ɾʲikka/`, not `/rikka/`）is a integral personal image share system, includes:

- A web application (See [Demo](#demo) section)
- A RESTful API server (See [API Doc][api-doc])
- A CLI tool named Rikkac based on the API (See [Rikkac Doc][rikkac-doc])
- image save plugins (See [Plugins Doc][plugins-doc])

Other part which not written in Golang (in plan):

- Andrild client
- iOS client

## Feature and Advantage

1. Simple and minimalist, no upload history
2. Image address can be copied to various formats
3. Image save part is plug-oriented
4. API provided
4. Modular Web server and API server 
5. CLI tool provided
6. **Only support latest verstion of Chrome/Firefox/Safari\***
7. Cute homepage image
8. An active maintainer

\*: Yes, It is advantage! If you can't copy url in view page or stock in "Please wait" message, it is certainly because your browser is too old to support javascript es6 syntax. Front-end is my weakness, I only learn javascript several days. I will be grateful if someone want help me to improve browser compatibility.

## Demo

There is a [Demo site][demo] built with Rikka, password is `rikka`, just try it.

homepage:

![homepage][home]

Click `Choose` button to choose an image.

Input password`rikka`.

Click `Upload` button.

If no error happened, you will be redirect to preview page:

![viewpage][view]

You will see a "Please wait" message If you uploaded a large file and save process is not finished, just wait a second.

When you see image url, you can click `Src`, `Markdown`, `HTML`, `rST`, `BBCode` button to copy image url in that format.

**But**: Once you close this page, you can't get it back except from browser history(Or you save the url).

This is intentional, Because main design concept is simple, just `Upload-Copy-Close-Patse`, then you can forget Rikka.

BTW: The preview image of Demo site is saved in Rikka too. 

## Plugins

Truly image save back-end of Rikka is written in plugin form, can be set by `-plugin` option.

Please see [Rikka Plugins Doc][plugins-doc] for alivaliable plugins.

## API

See [Rikka API Doc][api-doc].

## CLI - Rikkac

Rikkac is a CLI tool for Rikka based on Rikka RESTful API.

Build, install, configure and use guide can be found in [Rikkac Doc][rikkac-doc].

## Deploy

Want deploy Rikka system of you own? Check [Rikka Deploy Doc][deploy-doc] for deploy guide.

## Acknowledgements

- Thanks Golang and her developers
- Thanks Visual Studio Code and her developers
- Thanks open source

## License

All code of Rikka system are open source, based on  MIT license.

See [LICENSE][license].

[readme-zh]: https://github.com/7sDream/rikka/blob/master/Readme.zh.md

[badge-info-img]: https://images.microbadger.com/badges/image/7sdream/rikka.svg
[badge-version-img]: https://images.microbadger.com/badges/version/7sdream/rikka.svg

[image-in-dockerhub]: https://hub.docker.com/r/7sdream/rikka/

[demo]: http://7sdream-rikka-demo.daoapp.io/
[home]: http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-498160687
[view]: http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-457359417

[api-doc]: https://github.com/7sDream/rikka/tree/master/api
[rikkac-doc]: https://github.com/7sDream/rikka/tree/master/rikkac
[plugins-doc]: https://github.com/7sDream/rikka/tree/master/plugins
[deploy-doc]: https://github.com/7sDream/rikka/blob/master/deploy.md

[license]: https://github.com/7sDream/rikka/blob/master/LICENSE
