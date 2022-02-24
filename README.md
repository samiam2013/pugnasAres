# pugnasAres

A go server for dumping random bytes (currenlty 10 mb) back at a request. 

I'm using this alongside an nginx config to waste script kiddie's time. 

- [x] extra simple server that uses random data and Content-Type: text/hml header
- [ ] text mode based on [this code in my go exercises repo](https://github.com/samiam2013/learnGo/blob/91eb6101e655924af7a5df4699bd23ab2a94a4fb/intermediate/exercise05generateText.go 
- [ ] anonymous (no ip or timestamp) logging of request, user agent
- [ ] response caching (rotating list of random byte responses?)