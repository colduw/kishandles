# KISHandles
Kis <sub>(kiʃ)</sub> (Small) Vanity Handles for Bluesky

## So, what is it?
I saw [Simple Atproto Handles by furSUDO](https://github.com/furSUDO/simple-atproto-handles), and decided to re-write it in Go with some changes, ie: not using Qwik, Vite, Cloudflare Pages, and Cloudflare D1

**There are some caveats**
- There isn't a frontend, as of now, it's manual only (you insert the records, helpers for them can be found in the /cmd directory)
- PostgreSQL database required (use the .env.example as an example, you also need to set the DOMAIN variable)

## Usage
1. Rename the .env.example file to .env, then fill it out by changing the <code>DB_DSN</code> and the <code>DOMAIN</code> variables
2. Build with <code>go build main.go</code>, make sure you have [Go](https://go.dev) installed
3. Then, run ./main