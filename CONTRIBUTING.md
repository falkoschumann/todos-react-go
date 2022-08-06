# Contributing

## Style Guide

### Go

Differences between pointer and reference in constructor function with prefix `new` and `make`.

### JavaScript/TypeScript

Code is formatted with [Prettier](https://prettier.io). Prettier's configuration has only one rule added. Because JavaScript code can contain HTML code and vice versa, double quotes are used in HTML and single quotes are used in JavaScript to simplify this language mix.

The JavaScript and TypeScript code is analyzed by [ESLint](https://eslint.org). We use recommended rules. Only one rule is added. For stable SCM commits, imports must be case-insensitively sorted. Imports must be grouped in the following order:

1. Imports from other NPM packages.
2. Imports from other folders in the same NPM package.
3. Imports from the same folder in the same NPM package.
