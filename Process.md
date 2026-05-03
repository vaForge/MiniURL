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
  <br>
- Edit : I have implemented this working from store.go which contains a struct for in-meomry hash map of short and long url links , and then there are methods inside store to GetLinks and Save links (I've gona a bit overboard and included locks functionality too) also an instantiation of store(Newstore main.go used!)<br>
- Edit : Then we did make use of handler.go to build a handler for POST/shorten and GET/{shortURL} [also have handler struct to hold reference to store and methods to handle the routes] . Handler also contains generation of random shortURL for now only 6chars
- Edit Finally main.go to instantiate the store and handler and start the http server to listen to requests

### Phase 2 : Adding database (Persistence)

- Replace in-memory store with a database (e.g., SQLite, PostgreSQL)
- Create a table to store URL mappings with fields: id, short_url, long_url
- store.go to become database backend
- I'm still keeping the same API and handlers but swapping implementation of store to use db instead (no redesign yet, but fundamental working)
