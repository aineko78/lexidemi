# lexidemi
Word generator trying to keep the phonological character of a language

For the purpose of generating realistic words in conlangs (but of course it does also generate plausible words in real languages), this program computes the conditional frequency of 3 letters (trigrams, or what Shannon called third order approximation).
As first letter we choose only among those that are initials, so we are a bit less non-stationary (to be fully stationary we would need to consider the last letter of the previous word). 

The result is acceptable, if a bit repetive. Slices of runes are used internally, so all the gamut of utf-8 encoded unicode can be used.

USAGE

go build .\main.go

.\main.exe -f .\quenya.txt -n 10


TODO
It is possible to trick the program into loops, if an inadeguate statistics is provided.

We could add a sprinkle of creativity (a random kick off frequency) and tune this

Ability to handle puctuaction and diacritics, so that the text does not have to be polished
