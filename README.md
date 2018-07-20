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

## Sample Output

```bash
$ ./genflag -t segmented -l 5 -f 10
[FLAG]db1aN-CWLB9-gqo2S-DRW7E
[FLAG]SQtHo-3cdnn-iaFo3-WObel
[FLAG]TbrzR-Omf9z-mViVS-QG0vv
[FLAG]vKWJo-JzZaF-nVqcI-HcF3V
[FLAG]XT6sM-7Vv3a-A3idm-GNt1U
[FLAG]WYg9j-vUiAc-x5jhj-CLvrJ
[FLAG]xymqc-MpPo4-FYoNq-G1LNO
[FLAG]6K1gA-nQwrE-13Vzh-Y49JH
[FLAG]oqdmm-GL12s-hM70q-vvNsW
[FLAG]pNYtw-i4WC4-ayT3s-nbSGk
```