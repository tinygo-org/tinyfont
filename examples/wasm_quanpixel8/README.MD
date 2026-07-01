This example renders the glyph table for `quanpixel8.QuanPixel8` in a browser using TinyGo WebAssembly and an HTML canvas.

It displays every rune recorded in the generated `quanpixel8` font data. Blank glyphs are also included: even when a glyph draws no visible pixels, its Unicode code point is still shown in the table.

The example also renders a few short sample strings in Japanese, Simplified Chinese, Traditional Chinese, and Korean at the top of the page.

## Files

* `main.go` - WASM example program.
* `index.html` - Browser page that loads `main.wasm`.
* `runes.go` - Rune list extracted from `../../quanpixel8/quanpixel8.go`.

Generated files such as `main.wasm` and `wasm_exec.js` are not committed.

## Build

From the repository root:

`tinygo build -target wasm -o examples/wasm_quanpixel8/main.wasm ./examples/wasm_quanpixel8`

Copy TinyGo's WASM support script:

`cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" examples/wasm_quanpixel8/`

## Run

Serve the example locally:

`python -m http.server 8080 -d examples/wasm_quanpixel8`

Then open:

`http://localhost:8080/`

Alternatively, upload `index.html`, `main.wasm`, and `wasm_exec.js` to a website and open the page over HTTP or HTTPS.

