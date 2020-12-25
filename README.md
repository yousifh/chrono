# chrono
Static site generator in Go

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