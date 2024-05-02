# atproto-ipfs-bkup

util for backing up ATproto repos to IPFS.

### concept

the tool should enable people to simply backup their repo to IPFS.
the advanced functionality should be that the tool unwraps the '.car' repo export into something that IPFS natively understands.

todo: what should be the value of being able to directly address cid's from bsky in IPFS? pull your whole graph from bsky - why?

### tests

##### integration

to run the execution tests, run `make test-integration` in the root dir.

### todos:

- design CLI for:

  - uploading a file (returns the _path_ or _cid_?)
  - downloading a file (based on given _path_ or _cid_?)

- unit tests:

  - allow running unit tests separately

- integrations tests:

  - run integration tests as binary (add build tags)

  - emit app binary test coverage data:

    - build app as test binary
    - run the binary in the 'tester' container

  - allow easy execution (via makefile command?)

- CI/CD:
  - run unit + integration tests on PR's/pushes
