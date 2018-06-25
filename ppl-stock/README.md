# ppl-stock

[![Build Status](https://travis-ci.org/ob-vss-ss18/ppl-stock.svg?branch=master)](https://travis-ci.org/ob-vss-ss18/ppl-stock)

## Description

...

## HEROKU
https://ppl-stock.herokuapp.com/

## Team

- [Michael Eggers](https://github.com/michaeleggers)
- [Sven Schatter](https://github.com/Ghouz)
- [Philipp Konopac](https://github.com/konopac/)
- [Jannes Kretschmer](https://github.com/janneskretschmer)
- [Daniel Gabl](http://github.com/DManstrator)
- [Nico Daßler](https://github.com/truenicfel/)

See the other teams [here](https://github.com/ob-vss-ss18/ob-vss-ss18/wiki/Teams).

## Getting Started

### Prerequisites

run ```go get -u github.com/kardianos/govendor``` to install govendor.
It will install the dependencies automatically when running ```go install``` See: https://devcenter.heroku.com/articles/go-dependencies-via-govendor



### Using `go get`

#### Use a Fork

First fork the main repository (github.com/ob-vss-ss18/ppl-stock). Next clone the forked repository using the `got get`-command:

```shell
go get github.com/{your-git-username}/ppl-stock
```
This will clone your fork and place it into your go-folder specified by your `GOPATH` variable (default: ~/go).

You can then go ahead and import the project into goland:

1. *File -> Open*
2. Search for the newly downloaded git repo: `$GOPATH/src/github.com/{your-git-username}/ppl-stock/`
3. *OK*

This is all you need to do.

#### Not using a Fork

```shell
go get github.com/ob-vss-ss18/ppl-stock
```
This will clone the main and place it into your go-folder specified by your `GOPATH` variable (default: ~/go).

You can then go ahead and import the project into goland:

1. *File -> Open*
2. Search for the newly downloaded git repo: `$GOPATH/src/github.com/ob-vss-ss18/ppl-stock/`
3. *OK*

This is all you need to do.


### Using `git clone`

#### Use a Fork

First fork the main repository (github.com/ob-vss-ss18/ppl-stock). Now clone:

```shell
git clone https://github.com/{your-git-username}/ppl-stock.git
```

Remember to create a branch **before** pushing your code.

```shell
cd ppl-stock
git checkout -b {your branch name}
```

#### Not using a Fork

```shell
git clone https://github.com/ob-vss-ss18/ppl-stock.git
```

Remember to create a branch **before** pushing your code.

```shell
cd ppl-stock
git checkout -b {your branch name}
```

## [Status.md](https://github.com/ob-vss-ss18/ppl-stock/blob/master/Status.md)

Every tuesday during lab we updated this with all what was done during last week.

## Guidelines

- Do **not** commit project files from your IDE! (that means do not use `git add .` or add unwanted files to `.gitignore`).
  
(only in german)
- Wöchentlich (spätestens) zu Beginn des Praktikums: Statusupdate über die in der Zeit seit dem letzten Statusupdate erledigten Arbeiten, in der Datei Status.md, die es top-level in jedem Repository gibt. Neueste Einträge immer oben anfügen!
- Commits grundsätzlich nur in einen eigenen Branch bzw in einen Feature-Branch, keine Commits direkt im master-Branch.
- Achten Sie darauf, dass Sie möglichst jeden Commit einem Issue zuordnen. Um(für den Schein) nachvollziehen zu können, wer von Ihnen was gemacht hat, schaue ich mir die Ihnen zugewiesenen, geschlossenen Issues an.
- Nutzen Sie daher Issues für alle Arbeiten, nicht nur für Code. Nutzen Sie insbesondere auch Issues, wenn Sie Fragen an mich haben oder Unterstützung benötigen. Zur besseren Übersicht nutzen Sie Labels, Milestones, Projekte, etc. Diskutieren Sie über Issues um insbesondere Entscheidungen nachvollziehbar zu machen.
- Die Commits bzw. neue Features werden in den in den master Branch nur per Pull-Requests (PR) übernommen. Dabei gilt das 4-Augen-Prinzip. Das heisst derjenige der den PR in den master-Branch merged, darf nicht selbst Committer eines Commits im PR sein. Sie/Er reviewed den PR und merged ihn.
- Nutzen Sie Synergieeffekte. Es geht nicht darum, dass Sie sich nebeneinander her alle mit den selben Problemen beschäftigen. Unterstützen Sie sich, nutzen Sie bestehenden Code oder Code anderer Gruppen, wenn es sinnvoll ist. Dokumentieren Sie dies aber entsprechend in den Commits. Wenn Sie fremden Code nutzen und anpassen, bietet es sich beispielsweise an, zunächst den fremden Code unverändert zu committen und dann erst die Anpassungen in einzelnen Commits nachvollziehbar zu machen. Achten Sie auf die Lizenz des von Ihnen verwendeten Codes

## Even more Guidelines
(only in german)
- Kommentieren Sie den Code sinnvoll.
- Schreiben Sie Tests zum Code und streben Sie eine “vernünftige” Testabdeckung an. Lassen Sie die Tests automatisiert ausführen und die Abdeckung darstellen (Travis-CI, Jenkins, Coverage.io, ...)
- Alles Services sollen in Go implementiert werden.
- Das Frontend soll in Angular implementiert werden und auf Basis des Electron laufen.
- Die Services sollen auf Heroku deployed werden und eine MongoDB als Datenbank nutzen.
- Ihr gesamter Code sollte unter der BSD-Lizenz veröffentlicht werden. Nutzen Sie keine GPL-lizensierten Bibliotheken.
- Die Konfiguration eines Moduls/Service/Clients soll mit Hilfe von yaml-Dateien möglich sein.
- Die Kommunikation erfolgt über GraphQL-Schnittstellen vom Client zu den Services und zwischen den Services.
