# Problem 4.12

I used an on-disk database in combination with a search index to implement this solution. The purpose of the database is performance: we only periodically want to build the database and index. The search index links words from each comic's title to the record IDs stored in the database. We get O(1) lookups with this approach.

To build the database and search index, use:

```
go run 4.12 build
```

To make a search with a single term, use:

```
go run 4.12 search trees
```

Future refactorings I plan on:

* Refactor the two main functions, `BuildIndex` and `SearchIndex`, by extracting functionality into smaller functions.
* Create a test suite, including mocking out the API response.
