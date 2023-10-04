<sub>Credits to [leorcvargas' version](https://github.com/leorcvargas/rinha-go/) and [akitaonrails' video](https://youtu.be/EifK2a_5K_U) that helped and inspired me to make this.</sub>

#### About
This is my late version of the [rinha-de-backend-2023-q3](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/), a Brazilian community competition of APIs focused on performance that happened in Aug/23. I was not aware of the competition when it happened, but I found it so interesting that I decided to make my own version with Go - strongly influenced by [leorcvargas' version](https://github.com/leorcvargas/rinha-go/) and [akitaonrails' video.](https://youtu.be/EifK2a_5K_U)

#### Findings
- For this competition, it was not about what HTTP framework you chose (overall not language also). It was about infrastructure and throughput, and I definitely learned a thing or two. It was my first time learning about nginx, database connection pools, database indexing, etc.
- I was getting about ~35k inserts and couldn't figure what the bottleneck was. Even after tweaks to the resources, removing some unnecessary json marshal/unmarshals and enhancing the database-nginx throughput, it ended up being that I was not setting the `keepalive` directive in the `upstream` block of my nginx configuration file, meaning that new connections were being created on each request. After setting it up, my results were launched to ~46k.
- Also, Redis was probably unnecessary due to the limited resources to the fact that there were almost no duplicate reads... I think.

#### Results
You can `git clone` this repo, install [Gatling,](https://gatling.io/open-source/) change the executable with the fixed paths if needed and run the stress test in `/scritps/stress-test`, which is a copy of [the original.](https://github.com/zanfranceschi/rinha-de-backend-2023-q3/tree/main/stress-test)

![image](https://github.com/antoniopataro/rinha-go/assets/87823281/f254f099-c767-4891-8329-a778736cf23f)

#### Tech
- go v1.20.6;
- fiber;
- postgresql and pgx with pgxpool;
- redis and go-redis;
- nginx.
