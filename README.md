# MessageSecure
Full OpenSource privnote clone. Messages are encrypted with PGP on the client side and are never stored or transmitted unencrypted.

![version](https://img.shields.io/badge/beta-0.1.2-blue)
![Security](https://github.com/alexohneander/MessageSecure/workflows/CodeQL/badge.svg)
![Linter](https://github.com/alexohneander/MessageSecure/workflows/Linter/badge.svg)

![Preview](https://i.imgur.com/YKcoapx.png "MessageSecure Preview")


## Development

### Start the application 

```bash
go run main.go
```

```bash
cp .env.example .env
```
After copying, the variables still need to be adjusted.

## Motivation

I wanted to build a free Privnote clone to firstly build up my Go knowledge and to build a service that can be trusted as all content is openly available on the net.
