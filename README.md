# Cocoa application written in Go

Just launch Xcode and you can run this application on Xcode!

## Go version

1.2

## License

MIT

## What did I do?

1. Build a dynamic library instead of an executable (see the build setting).
2. Execute post_build.sh after building to build go files with the dynaic library.
3. Use the generated binary as the executable file (see the plist setting).
