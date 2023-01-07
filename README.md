# BookStack2Site

CLI tool which generates static sites from Bookstack Wikis.

**Usecases:**

- Sometimes you want a BookStack wiki for personal/team use and a public facing high traffic site for everyone else.
- Offline backup of your wiki which is good looking and easy to navigate.
- You want a markdown version of your wiki synced to a Git repo.

## Screenshots

TODO

## Usage

One day I'll get time to add the _automatically trigger SSG_ feature. Until then, this generates MdBook format markdown and you have to run the `mdbook build` command yourself.

If you just want markdown without an HTML site then don't run the `mdbook` command.

While setting up the first time:

```bash
mdboook init ./book-test
```

and edit the `book.toml` config to your liking.

Then every time you wanna download/update your wiki:

```bash
bookstack2site
    --bookstack-url=${BookStackEndpoint} \
    --token-id=${BookStackAPITokenID} \
    --token-secret=${BookStackAPITokenSecret} \
    --download-location="./book-test/src"

# to preview
cd ./book-test && mdbook serve -n 0.0.0.0

# to build
cd ./book-test && mdbook build
```

## Thanks

- The BookStack Project
- MdBook
