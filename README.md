# Multithreaded Text Justification

Program that reads an array of words and generates an array of strings made up
of the words fully justified such that each array element's length is set to the
specified width.

The program works by reading words from the input array until the total width
with spaces would exceed the specified line width. A thread is then created that
fully justifies the words and creates a string. The process repeats for all
lines excluding the last line. The last line is left justified in in a separate
thread as well.

In addition to the words provided as input to the justification threads, the
line number is also provided.

When the threads are finished justifying the line, the data is written to a
channel.

Another thread waiting for input from the channel writes the justified lines to
an array preserving the order by the provided line number. The array is
pre-sized to the maximum number in order to handle asynchronous arrival of lines
and to avoid the need for a synchronization primitive (i.e. mutex). This is an
example of a lock free algorithm.

```text
+------------------------------------------+
| ["foo", "bar", "baz", "foobar", ...]   G |
|                                          |
|                                          |
| FullJustifiy()                           |
+---------------------+--------------------+
                      |
                      +------------------------------------+
                                                           |
                                                           v
  +----------------------------------+       +--------------------------+
  | ["foo  bar", "baz     ", ...]  G |       | ["foo", "bar"]         G |-+
  |                                  |       |                          | |-+
  |                                  |<------+                          | | |
  | processIncomingLines()           |       | *JustifyLine()           | | |
  +----------------------------------+       +--------------------------+ | |
                                               +--------------------------+ |
                                                 +--------------------------+

G: Go Routine
```

## Examples

```Go
words := []string{
    "Science", "is", "what", "we", "understand", "well", "enough", "to",
    "explain", "to", "a", "computer.", "Art", "is", "everything", "else",
    "we", "do",
}
width := 20
result := []string{
    "Science  is  what we",
    "understand      well",
    "enough to explain to",
    "a  computer.  Art is",
    "everything  else  we",
    "do                  ",
},
```
