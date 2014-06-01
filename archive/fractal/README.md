fractal
=======

This project is dedicated to visualize the innate beauty of fractals.

Documentation
-------------

Documentation provided by GoDoc.

- [mset][]: constructs visual representations of the [Mandelbrot set][].
- [buddha][]: constructs visual representations of the Buddhabrot.

[Mandelbrot set]: https://en.wikipedia.org/wiki/Mandelbrot_set
[mset]: http://godoc.org/github.com/mewmew/playground/archive/fractal/mset
[buddha]: http://godoc.org/github.com/mewmew/playground/archive/fractal/buddha

Examples
--------

simple is an example which creates a basic visual representation of the
Mandelbrot set.

	go get github.com/mewmew/playground/archive/fractal/cmd/simple

![Simple visual representation of the Mandelbrot set](https://raw.github.com/mewmew/playground/master/archive/fractal/cmd/simple/simple.png)

pretty creates a more aesthetically looking visual representation of the
Mandelbrot set, with a gradient transition from black to blue to white.

	go get github.com/mewmew/playground/archive/fractal/cmd/pretty

![Pretty visual representation of the Mandelbrot set](https://raw.github.com/mewmew/playground/master/archive/fractal/cmd/pretty/pretty.png)

buddha creates a visual representation of the famous Buddhabrot, which requires
a more intensive calculation.

	go get github.com/mewmew/playground/archive/fractal/cmd/buddha

![Visual representation of Buddhabrot](https://raw.github.com/mewmew/playground/master/archive/fractal/cmd/buddha/buddha.png)

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
