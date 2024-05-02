# Open library dump parser and data reader

Intention of this repo is to provide a simple way to parse and read data from an open library dump file.

Firstly to get the data you need to download the dump from the open library website. You can find the
dumps [here](https://openlibrary.org/developers/dumps).
Once you have the dump you can use the parser to parse the data and then use the reader to read the data.
It's hard to read the data directly from the dump file as it is in a txt format, so I included a parser that parses txt
file and creates json files which hold structure.
Combined txt dump contains information about works, authors, editions, etc.
Parser will create json files for these categories:

- authors
- works
- editions

These files are required to read the data from the dump.
You can use the reader to read the data from the json files.
If you need to read the data separately you can use the `readers` module to read the data from the json files.

## How to parse the txt dump files

I've included a `bin/parser` file that you can use to parse the txt dump files.
Alternatively you can run the `cmd/parser/main.go` file to parse the files.

If you can look at the repo structure you will see I have 2 directories that the parser uses to parse the data.
One holds the txt dump files and the other is an output destination.

By default, the parser will look for the txt dump files in the `data` directory and will output the json files in
the `output` directory.
You can change the directories by running the binary and passing the directories as inputs while talking to the CLI.

The directories of dump file and outputs can be absolute paths on your system or relative paths to the binary.

To parse the files you can run the binary like this:

```bash
./bin/parser
```

## How to read the data from the json files

Currently, I have implemented 3 models to read the data from the json files.
You can use the `reader` module to read the data from the json files.

Example of how to read the data from the json files:

```go
    r := reader.Reader{
AuthorJsonLocation:  "./output/author.json",
WorkJsonLocation:    "./output/work.json",
EditionJsonLocation: "./output/edition.json",
}

authorChannel := make(chan models.Author)
go r.ReadAuthors(authorChannel)
for author := range authorChannel {
// do something with the author data
}
```

You can use the `ReadAuthors`, `ReadWorks`, and `ReadEditions` functions to read the data from the json files.

## Missing read features

I've implemented only the basic read features for the models.
Since data is not standardized in the dump files, I've implemented some json.RawMessage(example: `bio` in Author model)
fields in the models and skip some that I think are not important. (example is `table_of_contents` field in the Edition
model)

If you need to read some other fields you can add them to the models and implement the read feature in the reader
module.
