pacdump
=======

This tool creates an archive containing the files of the specified Arch Linux
packages.

Installation
------------

	go get github.com/mewmew/playground/cmd/pacdump

Usage
-----

	pacdump PKG...

Examples
--------

1. Create an archive ("boll.tar.gz") containing the files of the mesa package.

		pacdump mesa

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
