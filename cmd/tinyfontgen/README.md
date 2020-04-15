# tinyfontgen

Generate a font file for tinyfont project.  

* https://tinygo.org/x/tinyfont

## Description

By default, all ASCII characters are included.  
To save font flash size, you can delete the output Glyph on a line-by-line basis.  
If you want to include non-ASCII characters, you can use one of the following methods

* Specify the -all option to output all characters contained in the bdf font.
* Outputs all characters contained in the string specified by -string
* Output all characters in the file specified by -string-file.

```
usage: tinyfontgen [flags] [path ...]
  -all
        include all ascii glyphs (codepoint <= 255) in the font
  -ascii
        include all glyphs in the font (default true)
  -fontname string
        font name (default "TinyFont")
  -output string
        output path
  -package string
        package name (default "main")
  -string string
        strings for font
  -string-file value
        strings file for font (can be set multiple times)
  -verbose
        run verbosely
  -yadvance int
        new line distance
```

## Examples


### Noto Emoji

Google's font

* note: [Noto Color Emoji does not work](https://github.com/golang/freetype/issues/45#issuecomment-299074585)
* https://www.google.com/get/noto/#emoji-zsye
* http://sofia.nmsu.edu/~mleisher/Software/otf2bdf/

```
// First, convert fonts with otf2bdf
$ ./otf2bdf -o NotoEmoji-Regular-12pt.bdf NotoEmoji-Regular.ttf -p 12pt -r 75dpi
$ ./otf2bdf -o NotoEmoji-Regular-16pt.bdf NotoEmoji-Regular.ttf -p 16pt -r 75dpi
$ ./otf2bdf -o NotoEmoji-Regular-20pt.bdf NotoEmoji-Regular.ttf -p 20pt -r 75dpi

$ tinyfontgen --package notoemoji --fontname NotoEmojiRegular12pt NotoEmoji-Regular-12pt.bdf --output NotoEmoji-Regular-12pt.go --all --verbose
Approx. 3131 bytes

$ tinyfontgen --package notoemoji --fontname NotoEmojiRegular16pt NotoEmoji-Regular-16pt.bdf --output NotoEmoji-Regular-16pt.go --all --verbose
Approx. 4551 bytes

$ tinyfontgen --package notoemoji --fontname NotoEmojiRegular20pt NotoEmoji-Regular-20pt.bdf --output NotoEmoji-Regular-20pt.go --all --verbose
Approx. 6453 bytes
```

### M+ BITMAP FONTS

M+ BITMAP FONTS were designed to be simple and highly readable.  
It incorporates all Kanji (Chinese character in Japanese) until  
level 2 (the one described in JISX0208).  

* https://mplus-fonts.osdn.jp/
  * https://mplus-fonts.osdn.jp/mplus-bitmap-fonts/


```
// ascii
$ tinyfontgen.exe --package mplus --fontname mplus_j10b_iso mplus_j10r-iso-W4 --output mplus_j10b_iso.go --verbose
Approx. 3569 bytes

// ascii
$ tinyfontgen.exe --package mplus --fontname mplus_j10r mplus_j10r-iso-W4 mplus_j10r.bdf --output mplus_j10r.go --verbose
Approx. 3569 bytes

// ascii + all jisx0208
$ tinyfontgen.exe --package mplus --fontname mplus_j10r mplus_j10r-iso-W4 mplus_j10r.bdf --output mplus_j10r.go --verbose --all
Approx. 161526 bytes

// ascii + two char (\u4e16 + \u754c)
$ tinyfontgen.exe --package mplus --fontname mplus_j10r mplus_j10r-iso-W4 mplus_j10r.bdf --output mplus_j10r.go --verbose --string "世界"
Approx. 3615 bytes

// ' ' + 'e' + 'h' + 'l' + 'o' + \u4e16 + \u754c
$ tinyfontgen.exe --package mplus --fontname mplus_j10r mplus_j10r-iso-W4 mplus_j10r.bdf --output mplus_j10r.go --verbose --string "hello 世界" --ascii=false
Approx. 127 bytes
```

## Installation

    go get tinygo.org/x/tinyfont/cmd/tinyfontgen

## License

[This project is licensed](../../LICENSE) under the BSD 3-clause license, just like the [Go project](https://golang.org/LICENSE) itself.
