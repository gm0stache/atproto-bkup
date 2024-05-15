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

### setting up web2.storage

##### 1. generate a _key_ + configure _permissions_ (for the web3.storage API?):

1. `npm install -g @web3-storage/w3cli`
2. `w3 key create > private.key.full` and save the output string without `did:key:` prefix to a file _private.key_
3. add _permissions_ (aka. _capabilities_ in web3.storage speech) to your key: `w3 delegation create -c 'store/*' -c 'upload/*' [DID of private key] -o proof.ucan`
4. create a _space_ (basically a folder) where the data will be stored on web3.storage: `web3 space create atproto-bkup`, save the DID.
5. or
