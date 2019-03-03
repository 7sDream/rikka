# Rikka - A simple image share system

![][badge-version-img] ![][badge-info-img] ![][badge-license-img]

[中文版][readme-zh]

Rikka is written in Golang primarily, and provide Docker image.

Rikka image has been published to [DockerHub][image-in-dockerhub], just try it!

Badges above shows latest version and size of Rikka image.

## Introduction

Rikka（`りっか` in Japanese, sound like `/ɾʲikka/`, not `/rikka/`）is a integral personal image share system, includes:

- A web application (See [Demo](#demo) section)
- A REST API server (See [API Doc][api-doc])
- A CLI tool named Rikkac based on the API (See [Rikkac Doc][rikkac-doc])
- Image save plugins (See [Plugins Doc][plugins-doc] to get all available plugins)

Other parts not written in Golang (in plan):

- Android client
- iOS client

## Feature and Advantage

1. Simple and minimalist, no upload history
2. Image address can be copied to various formats
3. Many available image save plugins, such as weibo, QinNiu, UPai, Tencent Cloud, etc.
4. REST API provided
4. Modular Web server and API server
5. CLI tool provided
6. **Only guarantee support for recent versions of Chrome/Firefox/Safari**
7. Cute homepage image
8. An active maintainer :)

## Demo

There is a [Demo site][demo] built with Rikka, ~~password is `rikka`, just try it.~~ Because the free docker service provider I used stop it's free plan, the demo is on my personal VPS now. So the password is not given anymore, but you can also visit it and have a look :)

homepage:

![homepage][home]

Click `Choose` button to choose an image.

Input password`rikka`.

Click `Upload` button.

If no error happened, you will be redirect to preview page:

![view_page][view]

You will see a "Please wait" message If you uploaded a large file and save process is not finished, just wait a second.

When you see image url, you can click `Src`, `Markdown`, `HTML`, `RST`, `BBCode` button to copy image url in that format.

**But**: Once you close this page, you can't get it back except from browser history(Or you save the url).

This is intentional. The main design concept is simple, just `Upload-Copy-Close-Paste`, then you can forget Rikka.

BTW: The preview image of Demo site is saved in Rikka too. (But Github will put images which in Markdown files into its own CDN to accelerate access)

## Plugins

Truly image save back-end of Rikka is written as plugins, can be specified by `-plugin` option.

Please see [Rikka Plugins Doc][plugins-doc] for available plugins.

## API

See [Rikka API Doc][api-doc].

## CLI - Rikkac

Rikkac is a CLI tool for Rikka based on Rikka's REST API.

Build, install, configure and use guide can be found in [Rikkac Doc][rikkac-doc].

## Deploy

Want deploy Rikka system of you own? Check [Rikka Deploy Doc][deploy-doc] for deploy guide.

## Contribution

- Fork me
- Create a new branch from dev branch
- Add your code, comment, document and meaningful commit message
- Add yourself to CONTRIBUTION.md and describe your work
- PR to dev branch

Thanks all contributors!

You can see a list of contributors in [CONTRIBUTIONS.md][contributors].

## Acknowledgements

- Thanks Golang and her developers
- Thanks Visual Studio Code and her developers
- Thanks open source

## License

All code of Rikka system are open source, based on  MIT license.

See [LICENSE][license].

[readme-zh]: https://github.com/7sDream/rikka/blob/master/README.zh.md

[badge-info-img]: https://images.microbadger.com/badges/image/7sdream/rikka.svg
[badge-version-img]: https://images.microbadger.com/badges/version/7sdream/rikka.svg
[badge-license-img]: https://images.microbadger.com/badges/license/7sdream/rikka.svg

[image-in-dockerhub]: https://hub.docker.com/r/7sdream/rikka/

[demo]: https://rikka.7sdre.am/
[home]: https://rikka.7sdre.am/files/56c3ae9d-4d96-49c8-bc03-5104214a1ac8.png
[view]: https://rikka.7sdre.am/files/97bebf3b-9fb8-4b0c-a156-4b92b1951ae4.png

[api-doc]: https://github.com/7sDream/rikka/tree/master/api
[rikkac-doc]: https://github.com/7sDream/rikka/tree/master/rikkac
[plugins-doc]: https://github.com/7sDream/rikka/tree/master/plugins
[deploy-doc]: https://github.com/7sDream/rikka/blob/master/deploy.md

[contributors]: https://github.com/7sDream/rikka/blob/master/CONTRIBUTORS.md
[license]: https://github.com/7sDream/rikka/blob/master/LICENSE
