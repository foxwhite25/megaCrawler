## Purpose
MegaCrawler is a scraper that is based in colly, and updates information in database periodically. 

## Feature
* Service powered auto restart, can also run in clt.
* Log to service.
* Host a webserver on service.
* Built in webclient on clt to run task and check task at runtime.

## Example
In this scraper it is intended to built plugins and do an empty import on them. Using `init()` to register website.

Then use `megaCrawler.Start()` to launch the crawler.

When the crawler is listening, you can use these flag to check or change the service:

* `--start string` Launch the selected website now.
* `--get string` Get the status of the selected website.
* `--list` List all current registered websites.
* `--debug` To enable verbose output
* `--test` Test Connection to all plugins
* `--update` Update the program to the latest version in GitHub Release Page.

Note: `megaCrawler.Start()` is a blocking call.