# Pugnas Ares
The words for "Fighting" in Latin and Ares, the Greek god of war because that sounded better than Mars.

![from github.com/egonelbre](https://github.com/egonelbre/gophers/blob/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/fairy-tale/knight.svg)

<sub>[credit Egon Elbre](https://github.com/egonelbre/gophers/)</sub>

A go server for dumping random bytes (currenlty 10 mb) back at a request. 

I'm using this alongside an nginx config to waste script kiddie's time. 

you can access this at [ares.myres.dev](https://ares.myres.dev) (warning: 15mb page load)

- [x] extra simple server that uses random data and Content-Type: text/hml header
- [ ] text mode based on [this code in my go exercises repo](https://github.com/samiam2013/learnGo/blob/91eb6101e655924af7a5df4699bd23ab2a94a4fb/intermediate/exercise05generateText.go)
- [ ] text mode with descending encapsulated html tags (maybe templates to wrap it?)
- [ ] response caching (rotating list of random byte responses?)
