# Open library dump parser and data reader

Intention of this repo is to provide a simple way to parse and read data from a open library dump file.

Firstly to get the data you need to download the dump from the open library website. You can find the
dumps [here](https://openlibrary.org/developers/dumps).
Once you have the dump you can use the parser to parse the data and then use the reader to read the data.
Its hard to read the data directly from the dump file as it is in a txt format so I included a parser that parses txt
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

By default the parser will look for the txt dump files in the `data` directory and will output the json files in
the `output` directory.
You can change the directories by running the binary and passing the directories as inputs while talking to the CLI.

The directories of dump files and outputs can be absolute paths on your system or relative paths to the binary.

To parse the files you can run the binary like this:

```bash
./bin/parser
```