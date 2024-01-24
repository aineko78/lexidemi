# lexidemi
Word generator trying to keep the phonological character of a language


For the purpose of generating realistic words in conlangs (but of course it does also generate plausible words in real languages), this program computes the conditional frequency
of 3 letters (trigrams, or what Shannon called third order approximation).

The result is acceptable, if a bit repetive. Runes of string are used internally, so all the gamut of utf-8 encoded unicode can be used.