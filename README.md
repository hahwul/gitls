<h1 align="center">
  <br>
  <a href=""><img src="" alt="" width="500px;"></a>
  <br>
  <a href="https://app.codacy.com/gh/hahwul/gitls?utm_source=github.com&utm_medium=referral&utm_content=hahwul/gitls&utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/0ebdafdc2a3b4d85b1be09033ebbd83f"></a>
  <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat&logo=twitter"></a>
  <a href="https://github.com/hahwul"><img src="https://img.shields.io/github/stars/hahwul?style=flat&logo=github"></a>
</h1>

> abcd

## Installation
### From go-get
```
$ go get -u github.com/hahwul/gitls
```

## Usage 
```
  -l string
    	List of targets (e.g -l sample.lst)
  -o string
    	write output file (optional)
  -version
    	version of gitls
```

## Case Study
### Make all repo urls from repo/org/user urls
sample.lst
```
https://github.com/hahwul
https://github.com/tomnomnom/gron
https://github.com/tomnomnom/httprobe
https://github.com/s0md3v
```

make repo url list from sample file
```
$ gitls -l sample.lst
https://github.com/hahwul/a2sv
https://github.com/hahwul/action-dalfox
https://github.com/hahwul/asset-of-hahwul.com
https://github.com/hahwul/awesome-zap-extensions
https://github.com/hahwul/backbomb
https://github.com/hahwul/booungJS
https://github.com/hahwul/buildpack-nmap
https://github.com/hahwul/buildpack-zap-daemon
https://github.com/hahwul/can-i-protect-xss
https://github.com/hahwul/cyan-snake
https://github.com/hahwul/dalfox
https://github.com/hahwul/DevSecOps
https://github.com/hahwul/droid-hunter
https://github.com/hahwul/exploit-db_to_dokuwiki
https://github.com/hahwul/ftc
https://github.com/hahwul/gitls
https://github.com/hahwul/go-github-selfupdate-patched
https://github.com/hahwul/hack-pet
...snip...
https://github.com/hahwul/zap-cloud-scan
https://github.com/tomnomnom/gron
https://github.com/tomnomnom/httprobe
https://github.com/s0md3v/Arjun
https://github.com/s0md3v/AwesomeXSS
https://github.com/s0md3v/Blazy
https://github.com/s0md3v/Bolt
...snip...
https://github.com/s0md3v/velocity
https://github.com/s0md3v/XSStrike
https://github.com/s0md3v/Zen
https://github.com/s0md3v/zetanize
```

### Automated testing with [gitleaks](https://github.com/zricethezav/gitleaks)
```
$ gitls -l sample.lst | xargs -I % gitleaks --repo-url=% -v
```

### All clone target's repo
```
$ echo "https://github.com/paypal" | gitls | xargs -I % git clone %
```

## Contributors
![](/CONTRIBUTORS.svg)
