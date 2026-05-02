## Journey of Builing MiniURL

### Phase 1 : Builinng basic features ( MVP )

- Minimal working backend with basic features :
  - Input a URL
  - Generate a short URL
  - Store it
  - Redirect from short to long
    No auth, analytics, UI etc

- Just routing for POST/shorten and GET/{shortUrl}

- We need a store to save mapping of short to long URL, for now using in-memory store(hashmap) later we can use db
- We also need a handler to redirect the routes and http server to listen to requests
- simple struct `{short : string, long : string}` to store the mapping
