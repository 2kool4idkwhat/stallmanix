package main

// List of GNU packages, based on https://en.wikipedia.org/wiki/List_of_GNU_packages
//
// Slashes `-` are added around the names to avoid accidentally counting non-GNU
// packages with similar names
var listOfGnuPkgs = []string{
	// "base system"
	"bash",
	"coreutils",
	"cpio",
	"diffutils",
	"findutils",
	"gnugrep",
	"groff",
	"grub2",
	"gzip",
	"inetutils",
	"linux-libre",
	"plotutils",
	"readline",
	"screen",
	"gnutar",
	"texinfo",

	// GNU toolchain (without GCC)
	"binutils",
	"bison",
	"autoconf",
	"automake",
	"libtool",
	"gdb",
	"m4",
	"gnumake",

	// libraries
	"libbfd",
	"glibc",
	"fribidi",
	"gettext",
	"gnulib",
	"libmicrohttpd",
	"lightning",
	"libosip",
	"pth",

	// "Other compilers and interpreters"
	"clisp",
	"gawk",
	"gnucobol",
	"gcl",
	"guile",
	"mdk",
	"gnu-smalltalk",
	"mitscheme",
	"gforth",

	"gcc",
	"gfortran",
	"gnat",
	"gccgo",
	"gdc",

	// "Other developer tools"
	"ddd",
	"autogen",
	"cflow",
	"cppi",
	"gperf",
	"indent",
	"complexity",

	// desktops
	"gnustep",
	"windowmaker",

	// "General system administration"
	"acct",
	"ddrescue",
	"emacs",
	"guix",
	"libextractor",
	"mc",
	"mtools",
	"nano",
	"parallel",
	"parted",
	"gnupg",
	"stow",

	// scientific software
	"archimedes",
	"gnuastro",
	"gnucap",
	"datamash",
	"gmp",
	"octave",
	"gsl",
	"units",
	"R",
	"pspp",

	// "Internet"
	"jami",
	"mailman",
	"mailutils",
	"mediagoblin",
	"wget",
	"gnunet",
	"taler-exchange",

	// "Office"
	"aspell",
	"gcal",
	"miscfiles",
	"gtypist",
	"ocrad",

	// "Multimedia"
	"gimp",
	"libredwg",
	"lilypond",

	// games
	"gnubg",
	"gnuchess",
	"gnugo",
	"gnujump",
	"freedink",
	"liquidwar",

	// capitalism
	"gnucash",

	// fonts
	"freefont_ttf",
	"unifont",
}
