### Perlin Noise
Library for generating Perlin Noise in Go.

    go get github.com/jasdel/perlinnoise

* Based on [Java](http://cs.nyu.edu/~perlin/noise/) implementation of the Improved Noise.
* Also based on https://github.com/sol-prog/Perlin_Noise for random permutation genration.


### Usage
    ./perlinnoise/main.go provides sample usage.

    // Create with seed which will generate permutations randomly from the seed
    p := perlinnoise.New(*seed)

    // Create with default permutation
    p := perlinnoise.NewDefault()

    // Noise returns the noise value in three dementional space
    pn.Noise(x, y, z)


#### Example Executable
In the ./perlinnoise/ folder an example executable will print out the result of the noise.

The example takes thre options, h (height in rows), w (width in columns), and s seed for the noise.

    perlinnoise -h 20 -w 40 -s 10

    ~~~~~~~~..........~~~~~...........~~~~~.
    ~~~~~~~..,...~~~~~~~~~~...~~~~~...~~~~~.
    ~~~~~~~.,#,...~~~~~~.....~~~~~~..~~.,#,.
    ~~~~~~~~..~~~~~~~~~~..~~~~~~~~~~~~~,,,.~
    ,.~~~~~~.~~~~~~~~~~...~~~~~~~~.#,..,...,
    ..~~~~~,#,~~~~~~~~~~~~~~~~~~~~~..,,,,...
    ..~~~~~...~~~~.,.~~~~~~.......~.,#,...~.
    ..~~~...~~~~~~~.~~~~~~~.~~~.....,,.~~~~~
    .~~.,##,.~~~~~~~~~~~~~~~~~~~.,..,,.~~~~.
    .~~..,..~~~~~~~~~~~~~~~~~~~...~~..~~~~~~
    ~~~.....,#,.......~~~~~~~~.,,~~~~~~~~~~~
    ~~~~~~~~.,,,..~~~~~~~~~~~~~...~~~~~~~~~~
    ~~~~~~~~~~~~~~~~~~~......~~~..~~~~~~~~~~
    ~~~~~~~~~~~~~~~~~~~...~~~~~~..~~~~~~~~~~
    ~~.,.~~~.###..~...~~~~~..~~~~~~~~~~...~~
    ~.###,,,###,..~~~~~~~~~~~~~~~~~~~~~~~~~~
    .,,......~~~.,,...,,..~~~~~~..~~~~~~,#.~
    ..~~~~~~~~~~....~~~....~~~~~~~~~~~~,##,~
    .~~~.~~~~~~~~~~~~~~~.,.~~~~~~~~~~~~~..~~
    ,.~~~~~~~~~~~~~~~~~~...~~~~....~~~~~~~~~


### License
BSD 2 clause
