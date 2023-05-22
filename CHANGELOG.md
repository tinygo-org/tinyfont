0.4.0
---

- **all** 
    - avoid unneeded heap allocations, especially since embedded display value is interface anyhow
- **build**
    - improve Makefile to match other projects
    - switch to ghcr.io for docker container
    - switch to github actions
- **docs**
    - switch badge to GH actions
    - update README.md
- **make**
    - update so all task works as expected
- **modules**
    - update to latest tinygo drivers
- **tinyfontgen**
    - remove NULL character
    - Added support for generating fonts from ttf (#37)
- **misc**
    - add .gitignore


0.3.0
---

- Add Fonter and Glypher interfaces
- Refactor rotation process
- examples: use build directives for both Go1.17 and earlier versions
- Add boundingbox
- fix: tinyfontgen URL in README (#32)
- Return LineWidth for empty string as 0,0

0.2.1
---

- Update to use TinyGo drivers v0.14.0
- Correct examples use of NewSPI()

0.2.0
---

This is the first proper release!
