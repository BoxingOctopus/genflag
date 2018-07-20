# genflag

Stupid little Golang binary which will generate Flags for CTF/Wargames

## Building

```bash
$ build.sh genflag
```

> Note: Binaries are built into the `./build` directory

## Usage

```bash
$ ./genflag -h
  -d string
    	If flag type is segmented, what should the delimiter be? (default "-")
  -f int
    	Number of Flags to generate (default 1)
  -l int
    	If Flag Type is segmented, how long should each segment be? (default 4)
  -n int
    	Flag Suffix Length if type is 'single' (default 12)
  -p string
    	Prefix for Flag (default "[FLAG]")
  -t string
    	Type of Flag ('single' or 'segmented') (default "single")
```