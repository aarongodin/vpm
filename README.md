# vpm - Vim Package Manager

`vpm` is a CLI tool for managing packages with the default Vim package manager. Check out `:help packages` in Vim to learn more before using `vpm`.


## Overview

You can install `vpm` with:

```
go install github.com/aarongodin/vpm@latest
```

> TODO: provide install through other sources (homebrew for example)

Once you have `vpm` installed, try listing packages with:

```
vpm list
```

To learn more commands, check `vpm -h`. Otherwise, continue below for details on the most common commands.

### Adding and removing packages

As of yet, there is no central registry for vim packages. You can have `vpm` install them for you by providing a URL to clone from:

```
vpm add git@github.com:tpope/vim-fugitive.git
```

If you don't specify a group, the package is added to the `default` group. You can set a group and also a loading type (either `start` or `opt`) with options:

```
vpm add --group colors --load opt git@github.com:altercation/vim-colors-solarized.git
```

Removing a package is as simple as:

```
vpm remove vim-fugitive
```

Once a package is added, it's `name` is the remote repository name, including the username (e.g. `tpope/vim-fugitive`). However, when you refer to any added package in another command such as `remove` or `update`, you can use a partial string (such as `vim-fugitive`) as this is how we tend to the think of package names.

### Groups

Packages in the Vim default package manager are organized in groups. You can list groups with:

```
vpm groups
```

Groups can be specified when adding packages, otherwise the group `default` is used.

### Changing a package

Change a package's group or loading type with:

```
vpm change --load opt --group git vim-fugitive
```

### Managing updates

Updating a Vim package is as simple as performing a `git pull` on the directory where the package is installed. `vpm` makes this a bit easier by automating the operation across all packages.

```
vpm update
```

The `update` command pulls the latest from the origin remote in the git repository for all packages. You can specify one or many packages after, for example: `vpm update vim-fugitive nerdtree`.

If you want to check for updates first, you can run:

```
vpm outdated
```

A list is displayed showing the packages that are not on the latest git SHA from the remote.

### Sourcing a list of packages

You can add several packages at once using:

```
vpm source <filename>
```

The file is expected to be a YAML file including an array of packages to add. You can see what this file looks like for your currently installed packages with:

```
vpm ls -f yaml
```

Saving the output of the above to a file means you can transfer a list of packages to another machine using the file, or maintain the list in another location like a dotfiles repo.
