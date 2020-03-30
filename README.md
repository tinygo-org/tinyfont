TinyFont
=========
TinyFont is a font/text package for [TinyGo](https://tinygo.org/) displays. It's heavily based on [Adafruit's GFX library](https://github.com/adafruit/Adafruit-GFX-Library).

![example](https://raw.githubusercontent.com/conejoninja/tinyfont/master/example.png)


This package is experimental and may change in the future. It has not been optimized for speed or memory..

## Faster compilation
During compilation, tinygo will go through all the font files in a package and them discard them if not used. To improve compilation time considerably, move the files you are going to use to a new package. 

## About the fonts
The fonts compiled here were just converted or made compatible, and the original authors should be given proper credit. Each font is under its own license, and while most of them are under an _open license_, there might be differences in its usage and conditions.

## Generate your own font

You can use tinyfontgen to generate a tinyfont from a bdf font.  

https://github.com/tinygo-org/tinyfont/tree/rune/cmd/tinyfontgen

## Incompatibility warning

This package contains incompatible changes from [previous versions](https://github.com/tinygo-org/tinyfont/commit/a02e4495f8d64b671d923ec009e17c9da9e3e7f5).

* The argument has been changed from []byte to string.
  * You can simply add a string() cast.
* The Font struct has been changed.
  * If you are creating your own fonts, you need to modify them.
  * You may find [this script](https://github.com/sago35/tinyfont/tree/fontconv/cmd/tinyfontconv) helpful for font conversion.

## License

Software License Agreement (BSD License)

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

- Redistributions of source code must retain the above copyright notice,
  this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
