WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

openrisc
========

This project contains code related to the [OpenRISC][] architecture. The initial
aim is to provide an interface for encoding and decoding instructions, based on
the [OpenRISC 1000][] architecture specification.

It would be interesting to implement an emulator capable of running these
instructions in the future.

[OpenRISC]: http://opencores.org/or1k/Main_Page
[OpenRISC 1000]: http://opencores.org/websvn,filedetails?repname=openrisc&path=%2Fopenrisc%2Ftrunk%2Fdocs%2Fopenrisc-arch-1.0-rev0.pdf

Documentation
-------------

Documentation provided by GoDoc.

- [or1k-32][]: provides access to the 32-bit version of the Open RISC 1000 instruction sets.
   - [orbis][]: provides access to the OpenRISC Basic Instruction Set (ORBIS32).

[or1k-32]: http://godoc.org/github.com/mewmew/playground/openrisc/or1k-32
[orbis]: http://godoc.org/github.com/mewmew/playground/openrisc/or1k-32/orbis

public domain
-------------

This code is hereby released into the *[public domain][]*.

Comments derived from the [OpenRISC 1000][] architecture specification are
covered by the [GPL license][].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
[GPL license]: https://www.gnu.org/licenses/gpl.html
