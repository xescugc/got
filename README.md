# GOT (Go (?:Time)? Tracker)

`got` is a command line client to track the time you invest on your projects. It's a "port" from the Perl client [App::TimeTracker][tracker].

It follows the [XDG Base Directory Spec][xdg] to store data and configuration on your computer.

## Installation

```
$> go get github.com/XescuGC/got
```

## Commands

You can see al command it supports by running:

```
$> got -h
```

### init

The `init` command is the first command you need to run for each project you want to track, it intialize the directory structure and adds a `.got.json` to the directory you run it, wich will be the root of the project.

```
$> got init
```

### start

The `start` command starts tracking the current project.

```
$> got start
```

### current 

The `current` command shows the progress of the current task

```
$> got current
```

### stop

The `stop` command marks the current task as completed

```
$> got stop
```

### report

The `report` command reports all the time you have invested on your projects

```
$> got report
```


## Roadmap

This lib is still in development, the "manin" functionality is implemented but it still need more work, this is a list of the features that I think they are missing:

* [ ] Add Test
* [ ] Add `--tags|t` to the `start` and `report` to be able to tag the projects (to filter them on the `report`)
* [ ] Add a cool feature that [App::TimeTracker][tracker] has, which is the "multiple configurations files merge", basically the configuration file is a hierarchy of all the `.got.json` files from the current direcotry/project to the $HOME. This way you can define options to the configuration file and all the children direcotries will have them (ex: tags)
* [x] Add `flags` to the `report` command to filter the data


[tracker]: https://metacpan.org/pod/distribution/App-TimeTracker/bin/tracker
[xdg]: https://specifications.freedesktop.org/basedir-spec/latest/
