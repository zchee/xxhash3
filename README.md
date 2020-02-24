# xxhash3

[![CircleCI][circleci-badge]][circleci] [![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev] [![codecov.io][codecov-badge]][codecov] [![releases][release-badge]][release] [![ga][ga-badge]][ga]

Go implementation of the xxHash's XXH3 using SIMD instructions.

## XXH3

> XXH3 is based by xxHash (extremely fast non-cryptographic hash algorithm), New generation hash designed for speed on small keys and vectorization algorithm.

See more details: [New experimental hash algorithm - github.com/Cyan4973/xxHash](https://github.com/Cyan4973/xxHash#new-experimental-hash-algorithm)


## testing

- [SMHasher](https://github.com/aappleby/smhasher)


## Implementations

### C original implementation

- [Cyan4973/xxHash: Extremely fast non-cryptographic hash algorithm](https://github.com/Cyan4973/xxHash)

### Go packages

- Pure Go
  - XXH32
    - [StephaneBunel/xxhash-go: xxhash-go is a go (golang) wrapper for xxhash](https://bitbucket.org/StephaneBunel/xxhash-go/src/default/)
  - XXH64:
    - [OneOfOne/xxhash: A native implementation of the excellent XXHash hashing algorithm.](https://github.com/OneOfOne/xxhash)
    - [dgryski/go-xxh3: xxh3 fast hash function](https://github.com/dgryski/go-xxh3)

- Go + Go Plan9 Assembly
  - [cespare/xxhash: A Go implementation of the 64-bit xxHash algorithm (XXH64)](https://github.com/cespare/xxhash)
  - [zeebo/xxh3: XXH3 algorithm in Go](https://github.com/zeebo/xxh3)


<!-- badge links -->
[circleci]: https://app.circleci.com/github/zchee/xxhash3/pipelines
[codecov]: https://codecov.io/gh/zchee/xxhash3
[pkg.go.dev]: https://pkg.go.dev/github.com/zchee/xxhash3
[release]: https://github.com/zchee/xxhash3/releases
[ga]: https://github.com/zchee/xxhash3

[circleci-badge]: https://img.shields.io/circleci/build/github/zchee/xxhash3/master?logo=circleci&style=for-the-badge
[pkg.go.dev-badge]: https://bit.ly/2P3GHq1
[codecov-badge]: https://img.shields.io/codecov/c/github/zchee/xxhash3/master?logo=codecov&style=for-the-badge
[release-badge]: https://img.shields.io/github/release/zchee/xxhash3.svg?logo=github&style=for-the-badge
[ga-badge]: https://gh-ga-beacon.appspot.com/UA-89201129-1/zchee/xxhash3?useReferer&pixel
