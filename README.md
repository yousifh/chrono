# chrono
Static site generator in Go named after the Starcraft 2 ability [Chrono Boost](https://liquipedia.net/starcraft2/Chrono_Boost)

## Steps

- Read the site config
- Iterate over the templates and construct them in memory
- Iterate through the posts and build them
- Iterate through other pages and build them
- Generate everything in the `output` directory

## Questions
- Posts vs Pages or treat everything the same?
- How to generate RSS feed?
- What to do with tags?

## Directory structure
```
- config.yaml
- pages
    - index.md
    - posts
        - index.md
        - post1.md
        - post2.md
    - about
        - index.md
    - projects
        - index.md
- layouts
    - _defaults
        - base.html
        - index.html (group index page)
        - default.html (group single page)
    - _partials
        - footer.html
        - header.html
    - posts
        - index.html
        - default.html
    - index.html
- static
    - images
    - 404.html
```