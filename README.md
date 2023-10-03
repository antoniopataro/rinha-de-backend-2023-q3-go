<sub>Credits to [leorcvargas' version](https://github.com/leorcvargas/rinha-go/) and [akitaonrails' video](https://youtu.be/EifK2a_5K_U) that helped and inspired me to make this.</sub>

#### about
this is my late version of the [rinha-de-backend-2023-q3](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/), a Brazilian community competition of apis focused on performance that happened in aug/23. i was not aware of the competition when it happened, but i found it so interesting that i decided to make my own version with golang - strongly influenced by [leorcvargas' version](https://github.com/leorcvargas/rinha-go/) and [akitaonrails' video.](https://youtu.be/EifK2a_5K_U)

#### findings
- for this competition, it was not about what http framework you chose (overall not language also). it was about infrastructure and throughput, and i definitely learned a thing or two. it was my first time learning about nginx, db connection pools, db indexing, etc.
- i was getting about ~35k inserts and i couldn't figure what the bottleneck was. even after tweaks to the resources, removing some unnecessary json marshal/unmarshals and enhancing the database-nginx throughput, it ended up being that i was not setting the `keepalive` directive in the `upstream` block of my nginx configuration file, meaning that new connections were being created on each request. after setting it up, my results were launched to ~46k.
- also, redis was probably unnecessary due to the limited resources and since there were no duplicate reads... i think.

#### results
you can `git clone` this repo, install [gatling,](https://gatling.io/open-source/) change the executable with correct paths if needed and run the stress test in `/scritps/stress-test`, which is a copy of [the original.](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/tree/main/stress-test)

![image](https://github.com/antoniopataro/rinha-go/assets/87823281/f254f099-c767-4891-8329-a778736cf23f)

#### tech
- go v1.20.6;
- fiber;
- postgresql and pgx with pgxpool;
- redis and go-redis;
- nginx.
