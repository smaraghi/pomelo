# Pomelo

[![Go Report Card](https://goreportcard.com/badge/github.com/smaraghi/pomelo)](https://goreportcard.com/report/github.com/smaraghi/pomelo)

#### About

A fruity CLI for managing stale files on your system. 

#### HOWTO

##### Install

Install Go if it is not already installed.

Then from the terminal run `go get github.com/smaraghi/pomelo`.

##### Use

Pomelo ships with two flags, `-dir` and `-term`. 

The `-dir` flag is utilized to specify a directory to recursively search. The default directory provided is the HOME directory. Run `-dir=x` to use in the command line. 

The `-term` flag is utilized to specify a term length for checking your files' access times. The default term provided is 30 days. Run `-term=x` to use in the command line. 

#### Author

Written by [Serven Maraghi](https://github.com/smaraghi/).

#### License

Pomelo is subject to the terms of the Mozilla Pubic License, v. 2.0.

For more information on the MPL, please visit [Mozilla](http://mozilla.org/MPL/2.0/).
