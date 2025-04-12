# stallmanix - check if your NixOS system is actually a GNU/NixOS system

stallmanix is a CLI tool that checks how many of the packages in your NixOS closure are GNU software, so you can confidently tell Richard to shut up

## building

```
$ nix build
```

## usage

```
$ ./result/bin/stallmanix
Around 101 out of 2483 (4.07%) packages in your closure are GNU software
```

You can also set a custom profile path. For example, to check your user profile instead:

```
$ ./result/bin/stallmanix -path ~/.nix-profile
Around 67 out of 1541 (4.35%) packages in your closure are GNU software
```
