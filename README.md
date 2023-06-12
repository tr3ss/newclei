# newclei

A Go tool that gets the newest PRs from `projectdiscovery/nuclei-templates`. It will print the raw file URL.

https://github.com/projectdiscovery/nuclei-templates/pulls

# Install

```zsh
go install -v github.com/tr3ss/newclei@latest
```

# Usage

```zsh
❯ newclei
Usage of newclei:
  -cves
    	- Show only CVEs
  -token string
    	- GitHub token to use.
```

By default, you can do 60 unauthenticated requests/hour to the GitHub API. If needed, use a token to increase that limit.

```zsh
❯ newclei -token ghp_**************
https://github.com/projectdiscovery/nuclei-templates/raw/1d0d1edc658be15cb75d3e37fae86a97684d9b6b/http%2Fcves%2F2023%2FCVE-2023-1434.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/3a207683305cc5d34ddd709061f050dffca8f490/http%2Fcves%2F2022%2FCVE-2022-24627.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/7820f9bab108d173ba4733df838403f8fafdd0e5/http%2Fcves%2F2008%2FCVE-2008-1547.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/1a5a1cecd148164a222cf5c9c196c673f077c5d6/http%2Fcves%2F2008%2FCVE-2008-7269.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/5f8d2953472e7187e3a9be97d8988d3eae78600e/http%2Fcves%2F2010%2FCVE-2010-1586.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/4b03a558db967dbee1f3dd1d6f96f6c944c8b0cb/http%2Fcves%2F2013%2FCVE-2013-2621.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/171082cf0ebe9bf6b68601999ff41a167fea8dee/http%2Fcves%2F2011%2FCVE-2011-5252.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/b5726fac4abad0b3929957bd51df4d4536492574/http%2Fcves%2F2012%2FCVE-2012-4982.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/b45e6b9e49d6a608a453748c79387ac25e8ca384/http%2Ftakeovers%2Fwebflow-takeover.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/9fff393178d0c4fb17ed69ed2bb18a1cc05ce92c/http%2Ftechnologies%2Fnacos-version.yaml
https://github.com/projectdiscovery/nuclei-templates/raw/5f31b7cbc01ac4b0a64dcbf247cc93c5e370dace/http%2Fosint%2Fdotcards.yaml
[...]
```

# Note

Always make sure the templates are legit, as these are not validated templates.