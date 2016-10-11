# VSCode AutoComplete Issue Demonstration

I've noticed that VSCode and the Go plugin have issues with autocomplete and import of subpackages.
If you clone this simple repo, you should be able to reproduce this issue.

Basically, `> Go: Add Import` does not find subpackages (or neighbor package if you are in a subpackage) and autocomplete does not work.
However, you can mouse over the method to see its signature.

 - VSCode Version 1.5.3 (1.5.3)
 - OSx 10.11.6
 - Go 1.7.1