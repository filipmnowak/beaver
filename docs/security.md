# Security

## SQLITE3

### Input Sanitization

Input is not being sanitized, and this is - in general - not a problem:
- App doesn't accept user input.
- DB is used as an internal key-value store with inputs used to render HTML.
- Most of the input comes or will come from test scripts bundled with the app, and those will include tools installed at OS level.
- Communication with web pages is one directional.

So, input sanitization would be good to have, it's not a high prio.

What might help (or might not) is [parameter](https://www.sqlite.org/cli.html#sql_parameters) [binding](https://www.sqlite.org/lang_expr.html#varparam).
