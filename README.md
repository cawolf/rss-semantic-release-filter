[![test](https://github.com/cawolf/rss-semantic-release-filter/actions/workflows/test.yml/badge.svg)](https://github.com/cawolf/rss-semantic-release-filter/actions/workflows/test.yml)

# rss-semantic-release-filter

This application filters RSS feeds from GitHub project releases by semantic version levels and publishes it as a new RSS feed for you to consume. This way, you can specify for each project, which releases you are interested in, instead of searching through all releases for relevant changes interesting to you.

## installation

You can either build a binary for your system from this source code or use a [docker image](https://github.com/cawolf/rss-semantic-release-filter/pkgs/container/rss-semantic-release-filter).
An example `docker-compose.yml` is provided in the repository.

## usage

### create config.yaml

You need to configure the release feeds in a `config.yaml` file. Each feed needs a URL and a minimum semantic version level. A minimum configuration file looks like this:

```yaml
feeds:
  - url: https://github.com/kubernetes/kubernetes/releases.atom
    minimum_level: minor
```

This example would find only the latest Kubernetes releases, which are `minor` or above.

### refresh feed data

With a configuration file in the same directory as the binary, you can refresh the feed data:

```shell
./rss-semantic-release-filter refresh
```

This populates a local database with all matched feed items.

### generate filtered feed

When the local database was refreshed, you can serve your filtered feed:

```shell
./rss-semantic-release-filter feed
```

This opens an HTTP server on port 1323, serving your feed at the root URL. You can now consume the feed:

```shell
curl localhost:1323
```
