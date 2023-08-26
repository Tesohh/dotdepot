# 📦 Dotdepot

Dotdepot is a simple utility, written in Go, to manage your dotfiles across multiple computers and OSes (crossplatform) and take the pain out of keeping your config up to date on all your machines.

## Sharing configs

It's easy to share a config with someone.
The link format to share a config is `dotdepot.pyros.dev/<username>`

So if you want to share your config, or add it to your link tree or something similar, here you go

## Can't i just use git?

Yes, you can, but it only works well if you have all machines running the same os.
For example, VSCode saves it's config in a different folder for each OS.
Plus, as a bonus, you can configure dotdepot to run a installation script after pulling all your files.

## Installing

- With go: `go install github.com/Tesohh/dotdepot`

## After installing

1. Create a `~/.config/dotdepot` folder
2. Add the two files:
   - `login.yml`
   - `config.yml`
3. Fill out login.yml:

```yml
username: <yourUsername>
password: <yourPassword> # if you want to just pull, you don't need the password
```

4. run `dotdepot signup`
5. Fill out config.yml (see [configuration](#configuration))

## Configuration

> Note: if you want to just pull, you can leave `config.yml` empty.
> Note: you only need to write this on one machine. It is synced automatically with every push and pull.

Here's an example config.yml:

```yml
files:
  - macos: ~/.vimrc
    linux: ~/.vimrc
    # windows is ignored because i didn't specify it

  - macos: ~/Library/Application Support/Code/User/settings.json
    linux: ~/.config/Code/settings.json
    windows: "%APPDATA%/Code/User/settings.json"
    # if it has special characters like %, wrap in quotes
    # also notice how i didn't use backslash.

directories: # directories are read/written to recursively
  - macos: ~/.config/nvim/
    windows: ~/.config/nvim/
```

Please note that indentation is important (make sure to follow yaml rules)

## Notes

I am NOT responsible for any files uploaded/downloaded/run through the service

## Feature wishlist

- [ ] Sync entire directories instead of just files
- [ ] Add option to run a pre-push and post-pull script (with previous confirmation)
