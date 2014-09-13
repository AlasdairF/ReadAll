## ReadAll

`readall.ReadAll` is meant to replace `ioutil.ReadAll` when you already know the (uncompressed) length of the file. This is more efficient than `ioutil.ReadAll` because it reads directly into the correctly sized slice of bytes.

## Usage

`data, err := readall.ReadAll(ioreader, length)`
