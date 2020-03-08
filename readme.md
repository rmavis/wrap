# Wrap

This program combines functionality for splitting input strings on various characters and printing that split input wrapped and joined by other strings.

`wrap` works with the concepts of sets, records, and fields. By default, the entirety of the input stream is treated as a set, and each line of input is treated as a record in that set, and each tab-separated part of each line is treated as a field in that record. In the output, fields can wrapped in some character (by default, a double quote) or not, and separated from each other by some other character (by default, a space) or not, and records can be wrapped (by default, by open and close parentheses) or not, and separated from each other (by default, by newlines) or not. Also, the set of records can be wrapped (by default, by open and close parentheses) or not.


## Arguments

All arguments specify either strings to split records or fields, or strings for wrapping or joining the set, records, or fields.

The set itself doesn't need to be split---it is defined by the input from stdin (though it might be nice to be able to include multiple sets? TODO).

For reading input:

    -rs STR
        Split records on STR. Defaults to \n.
    -fs STR
        Split fields on STR. Defaults to \t.

For printing output:

    -sw STR1 STR2
        Specifies the set wrap open and close strings. Defaults to ( and ).
    -swo STR
        Specifies the set wrap open string as STR. Defaults to (.
    -swc STR
        Specifies the set wrap close string as STR. Defaults to ).
    -rw STR1 STR2
        Specifies the record wrap open and close strings. Defaults to ( and ).
    -rwo STR
        Specifies the record wrap open string as STR. Defaults to (.
    -rwc STR
        Specifies the record wrap close string as STR. Defaults to ).
    -rj STR
        Specifies the record join string as STR. Defaults to \n.
    -fw STR1 STR2
        Specifies the field wrap open and close strings. Defaults to " and ".
    -fwo STR
        Specifies the field wrap open string as STR. Defaults to ".
    -fwc STR
        Specifies the field wrap close string as STR. Defaults to ".
    -fj STR
        Specifies the field join string as STR. Defaults to one space.

For the output arguments above, if there is a trailing dash (e.g., `-sw-`), an empty string will be used.

Output shorthand arguments:

    --json
        Equivalent to "-sw [ ] -rw [ ] -rj , -fj ,"
    --csv
        Equivalent to "-sw- -rw- -fj ,"


## Examples

Input:

    $ cat test.data
    Jan 1	Happy new year
    Jan 3	Happy birthday
    Apr 1	April Fool's Day

All examples split records on newlines and fields on tabs.

Output with no options:

    $ wrap < test.data
    (("Jan 1" "Happy new year")
    ("Jan 3" "Happy birthday")
    ("Apr 1" "April Fool's Day"))

Output when specifying blank set and record wrap characters, and the field separator as a comma:

    $ wrap -sw- -rw- -fj , < test.data
    "Jan 1","Happy new year"
    "Jan 3","Happy birthday"
    "Apr 1","April Fool's Day"

Output when specifying set wrap characters as square braces, record wrap characters as curly braces, the record separator as a comma, and the field separator as a colon:

    $ wrap -sw [ ] -rw { } -rj , -fj : < test.data
    [{"Jan 1":"Happy new year"},{"Jan 3":"Happy birthday"},{"Apr 1":"April Fool's Day"}]

Output when specifying a blank for the set, record, and field wraps, two colons for the record separator, and one colon for the field separator:

    $ wrap -sw- -rw- -fw- -rj :: -fj : < test.data
    Jan 1:Happy new year::Jan 1:Happy birthday::Apr 1:April Fool's Day


Convert from `calendar`-formatted output to CSV and from that to JSON:

    $ wrap --csv < test.data | wrap --json -fw-
    [["Jan 1","Happy new year"],["Jan 3","Happy birthday"],["Apr 1","April Fool's Day"]]
