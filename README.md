# atproto-ipfs-bkup

Util for backing up ATproto repos to IPFS.

### concept

the tool should enable people to simply backup their repo to IPFS.
the advanced functionality should be that the tool unwraps the '.car' repo export into something that IPFS natively understands.

todo: - what should be the value of being able to directly address cid's from bsky in IPFS? pull your whole graph from bsky - why?

2dos:

- todo: upload 'data' as _file_ to IPFS? or add name of the uploaded file as metadata in order to allow reconstructing the input. -
- todo: add 'tester' app to run the unit/integration tests agains an actual _kubo_ instance

### tests

##### integration

to run the execution tests, run `docker compose up --build` in the root dir.
