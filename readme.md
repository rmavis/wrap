# Wrap

The `wrap` program can be used to wrap lines and parts of lines with symbols. It combines functionality for splitting (when reading input) and joining (when writing output) strings.

`wrap` reads from stdin. Each line of input is treated as a record. Fields in each record are separated by some character (by default, the tab character). In the output, fields can wrapped in some character (by default, a double quote) or not, and separated from each other by some other character (by default, a space) or not, and records can be wrapped (by default, by open and close parentheses) or not, and separated from each other (by default, by newlines) or not. Also, the set of records can be wrapped (by default, by open and close parentheses) or not.

## Examples

Input:

    Jan 1	Happy new year
    Jan 3	Happy birthday
    Apr 1	April Fool's Day

Output with no options:

    (("Jan 1" "Happy new year")
    ("Jan 3" "Happy birthday")
    ("Apr 1" "April Fool's Day"))

Output when specifying set wrap characters as square braces, record wrap characters as curly braces, record separator as a comma, and field separator as a colon:

    [{"Jan 1":"Happy new year"},{"Jan 3":"Happy birthday"},{"Apr 1":"April Fool's Day"}]

Output when specifying blank set and record wrap characters, and the field separator as a comma:

    "Jan 1","Happy new year"
    "Jan 3","Happy birthday"
    "Apr 1","April Fool's Day"

Output when specifying a blank for the set wrap, record wrap, field wrap, and field separator characters, and record separator as a colon:

    Jan 1 Happy new year:Jan 1 Happy birthday:Apr 1 April Fool's Day

## Arguments

All arguments specify either characters to split records or fields or characters for wrapping the set, records, or fields.

The set itself doesn't need to be split---it is defined by the input from stdin (though it might be nice to be able to include multiple sets? TODO).

For reading input:

    -rs STR
        Split records on STR. Defaults to \n.
    -fs STR
        Split fields on STAR. Defaults to \t.

For printing output:

    -sw STR1 STR2
        Specifies the open and close set wrap characters. Defaults to ( and ).
    -swo STR
        Specifies the open set wrap as STR. Defaults to (.
    -swc STR
        Specifies the close set wrap as STR. Defaults to ).
    -rw STR1 STR2
        Specifies the open and close record wrap characters. Defaults to ( and ).
    -rwo STR
        Specifies the open record wrap as STR. Defaults to (.
    -rwc STR
        Specifies the close record wrap as STR. Defaults to ).
    -rj STR
        Specifies the record join as STR. Defaults to \n.
    -fw STR1 STR2
        Specifies the open and close field wrap characters. Defaults to " and ".
    -fwo STR
        Specifies the open field wrap as STR. Defaults to ".
    -fwc STR
        Specifies the close field wrap as STR. Defaults to ".
    -fj STR
        Specifies the field join as STR. Defaults to one space.

For all output arguments, if there is a trailing dash (e.g., `-sw-`), an empty string will be used.

Shorthand:

    --json
        Equivalent to "-sw [ ] -rw [ ] -rj , -fs ' ' -fj "
    --csv
        Equivalent to "-sw- -rw- -fs ' ' -fj ,"
