<h1 align="center">
  <br>
  <a href=""><img src="https://user-images.githubusercontent.com/13212227/106487253-d51aca80-64f5-11eb-8c30-134ca3d4cf2b.png" alt="" width="500px;"></a>
  <br>
  <img src="https://github.com/hahwul/gitls/workflows/Build/badge.svg">
  <img src="https://github.com/hahwul/gitls/workflows/CodeQL/badge.svg">
  <img src="https://api.codacy.com/project/badge/Grade/0ebdafdc2a3b4d85b1be09033ebbd83f">
  <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat&logo=twitter"></a>
  <a href="https://github.com/hahwul"><img src="https://img.shields.io/github/stars/hahwul?style=flat&logo=github"></a>
</h1>

> Listing git repository from `URL`/`User`/`Org`

This tool is available when the repository, such as github, is included in the bugbounty scope. Sometimes specified as an org name or user name rather than a specific repository, you can use this tool to extract url from all public repositories included in the org/user.

This can be used for various actions such as scanning or cloning for multiple repositories.

> 🚧 NOTICE <br>
For unauthenticated requests in github api, the rate limit allows for up to 60 requests per hour. Unauthenticated requests are associated with the originating IP address, and not the user making requests.
[https://docs.github.com/en/rest/overview/resources-in-the-rest-api](https://docs.github.com/en/rest/overview/resources-in-the-rest-api)<br><br>
So too many tasks can be blocked by the API for a certain time from github. In this case, you can select the appropriate destination or access and use any IP using the torsocks(e.g `torsocks gitls -l user.list`) or `-tor` options.

## Installation
### From go-get
```
▶ GO111MODULE=on go get -v github.com/hahwul/gitls
```
### Using homebres
```
▶ brew tap hahwul/gitls
▶ brew install gitls
```
### Using snapcraft
```
▶ sudo snap install gitls
```

## Usage 
```
Usage of gitls:
  -include-users
    	include repo of org users(member)
  -l string
    	List of targets (e.g -l sample.lst)
  -o string
    	write output file (optional)
  -proxy string
    	using custom proxy
  -tor
    	using tor proxy / localhost:9050
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
▶ gitls -l sample.lst
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

### Get all repository in org and included users(members)
```
▶ echo https://github.com/paypal | ./gitls -include-users
```

```
....
https://github.com/paypal/tech-talks
https://github.com/paypal/TLS-update
https://github.com/paypal/yurita
https://github.com/ahunnargikar
https://github.com/ahunnargikar/docker-chronos-image
https://github.com/ahunnargikar/docker-tomcat7
https://github.com/ahunnargikar/DockerConDemo
https://github.com/ahunnargikar/elasticsearch-registry-backend
https://github.com/ahunnargikar/elasticsearchindex
https://github.com/ahunnargikar/jenkins-dind
https://github.com/ahunnargikar/jenkins-standalone
https://github.com/ahunnargikar/vagrant-mesos
https://github.com/ahunnargikar/vagrant_docker_registry
https://github.com/anandpalanisamy
https://github.com/anilgursel
https://github.com/anilgursel/squbs-sample
https://github.com/bluepnume
```

### Automated testing with [gitleaks](https://github.com/zricethezav/gitleaks)
```
▶ gitls -l sample.lst | xargs -I % gitleaks --repo-url=% -v
```

### All clone target's repo
```
▶ echo "https://github.com/paypal" | gitls | xargs -I % git clone %
```

## Contributors
![](/CONTRIBUTORS.svg)
