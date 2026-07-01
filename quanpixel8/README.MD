# quanpixel8

Package quanpixel8 provides an 8x8 fixed-cell bitmap font derived from
QuanPixel for use with tinyfont.

The font is exposed as QuanPixel8.

```
import "tinygo.org/x/tinyfont/quanpixel8"

tinyfont.WriteLine(display, quanpixel8.QuanPixel8, 0, 0, "文字", white)
```

## About the font

QuanPixel is an 8x8 bitmap font by Poxiao Fonts.

Project page:

```
https://diaowinner.itch.io/galmuri-extended
```

The font data was generated from QuanPixel BDF data using tinyfontgen.

## Character coverage

The generated font was checked against several CJK legacy character sets by
comparing Unicode code points extracted from the generated QuanPixel8 font data
with Unicode code point sets decoded from each character set.

| Language / region                   | Character set / subset          |        Covered | Missing |
| ----------------------------------- | ------------------------------- | -------------: | ------: |
| Japanese                            | JIS X 0208 level 1 kanji        |  2,965 / 2,965 |       0 |
| Japanese                            | JIS X 0208 level 2 kanji        |  3,390 / 3,390 |       0 |
| Japanese                            | JIS X 0208 non-kanji characters |      503 / 524 |      21 |
| Japanese                            | JIS X 0208 total                |  6,858 / 6,879 |      21 |
| Japanese                            | JIS X 0213 level 3 kanji        |    628 / 1,128 |     500 |
| Japanese                            | JIS X 0213 level 4 kanji        |    802 / 2,436 |   1,634 |
| Simplified Chinese / Mainland China | GB 2312 level 1 hanzi           |  3,755 / 3,755 |       0 |
| Simplified Chinese / Mainland China | GB 2312 level 2 hanzi           |  3,008 / 3,008 |       0 |
| Simplified Chinese / Mainland China | GB 2312 non-hanzi characters    |      613 / 682 |      69 |
| Simplified Chinese / Mainland China | GB 2312 total                   |  7,376 / 7,445 |      69 |
| Traditional Chinese / Taiwan        | Big5 symbols                    |      396 / 406 |      10 |
| Traditional Chinese / Taiwan        | Big5 level 1 hanzi              |  5,401 / 5,401 |       0 |
| Traditional Chinese / Taiwan        | Big5 level 2 hanzi              |  2,444 / 7,652 |   5,208 |
| Traditional Chinese / Taiwan        | Big5 other decoded characters   |      249 / 249 |       0 |
| Traditional Chinese / Taiwan        | Big5 total                      | 8,488 / 13,706 |   5,218 |
| Korean                              | KS X 1001 Hangul                |  2,350 / 2,350 |       0 |
| Korean                              | KS X 1001 Hanja                 |  4,418 / 4,888 |     470 |
| Korean                              | KS X 1001 other characters      |      715 / 987 |     272 |
| Korean                              | KS X 1001 total                 |  7,483 / 8,225 |     742 |
| Korean                              | Unicode Hangul Syllables        | 4,363 / 11,172 |   6,809 |

The font fully covers the following checked subsets:

* JIS X 0208 level 1 and level 2 kanji
* GB 2312 level 1 and level 2 hanzi
* Big5 level 1 hanzi
* Big5 other decoded characters
* KS X 1001 Hangul

The font does not fully cover the following checked sets or subsets:

* JIS X 0208 non-kanji characters
* JIS X 0208 as a whole
* JIS X 0213 level 3 and level 4 kanji
* GB 2312 non-hanzi characters
* GB 2312 as a whole
* Big5 symbols
* Big5 level 2 hanzi
* Big5 as a whole
* KS X 1001 Hanja
* KS X 1001 other characters
* KS X 1001 as a whole
* Unicode Hangul Syllables

Missing-character lists are included in MISSING.MD. The MISSING.MD file lists
the specific non-zero subsets used by the coverage check. Total rows in the table
above may be represented by their component subset lists to avoid duplicating
the same missing characters.

## Coverage check basis

The coverage checker used Python standard codecs and the following ranges:

| Character set            | Basis                                                                                                                                           |
| ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| JIS X 0208               | `iso2022_jp`; rows 16-47 as level 1 kanji, rows 48-84 as level 2 kanji; all other decoded code points as non-kanji characters                   |
| JIS X 0213               | `euc_jis_2004`; plane 1 non-JIS-X-0208 CJK ideographs as level 3 kanji; plane 2 rows 1, 3, 4, 5, 8, 12-15, and 78-94 as level 4 kanji           |
| GB 2312                  | `gb2312`; rows 1-9 as non-hanzi characters, rows 16-55 as level 1 hanzi, rows 56-87 as level 2 hanzi                                            |
| Big5                     | `big5`; A140-A3FE as symbols, A440-C67E as level 1 hanzi, C940-F9D5 as level 2 hanzi; remaining decoded code points as other decoded characters |
| KS X 1001                | `euc_kr`; rows 16-40 as Hangul, rows 42-93 as Hanja; all other decoded code points as other characters                                          |
| Unicode Hangul Syllables | U+AC00 through U+D7A3                                                                                                                           |

For Big5, the total row includes all unique Unicode code points decoded in the
checked Big5 range. The named rows split out the symbols, level 1 hanzi, and
level 2 hanzi ranges. Remaining decoded code points are counted as "other
decoded characters".

The range choices above are based on the following public references:

* JIS X 0208:
  https://en.wikipedia.org/wiki/JIS_X_0208

* JIS X 0213:
  https://en.wikipedia.org/wiki/JIS_X_0213

* GB 2312:
  https://en.wikipedia.org/wiki/GB_2312

* Big5:
  https://www.cns11643.gov.tw/pageView.jsp?ID=9&la=1

* KS X 1001:
  https://en.wikipedia.org/wiki/KS_X_1001

* Unicode Hangul Syllables:
  https://www.unicode.org/charts/PDF/UAC00.pdf

## License

The tinyfont project is licensed under the BSD 3-Clause License.

The Go implementation code added in this package follows the tinyfont project
license.

The bitmap font data included in this package is derived from QuanPixel. That
font data is licensed under the SIL Open Font License 1.1. See LICENSE_OFL.txt.

## Attribution

The QuanPixel project states that it uses glyph data from Galmuri and Chill
Bitmap.

