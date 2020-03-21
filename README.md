TinyFont
=========
TinyFont is a font/text package for [TinyGo](https://tinygo.org/) displays. It's heavily based on [Adafruit's GFX library](https://github.com/adafruit/Adafruit-GFX-Library).

![example](https://raw.githubusercontent.com/conejoninja/tinyfont/master/example.png)


This package is experimental and may change in the future. It has not been optimized for speed or memory..

## Faster compilation
During compilation, tinygo will go through all the font files in a package and them discard them if not used. To improve compilation time considerably, move the files you are going to use to a new package. 

## About the fonts
The fonts compiled here were just converted or made compatible, and the original authors should be given proper credit. Each font is under its own license, and while most of them are under an _open license_, there might be differences in its usage and conditions.

## License

[This project is licensed](./LICENSE) under the BSD 3-clause license, just like the [Go project](https://golang.org/LICENSE) itself.
