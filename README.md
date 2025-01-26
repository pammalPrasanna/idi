<h2>idi</h2> 
<div style="text-align: center;">
  <h4>
	Idiomatic* Golang web project scaffolder
  </h4>
</div>

<h4>Objectives</h4>

- A CLI tool that allows me to setup a idomatic* and modular* Go project with common web helpers.
- Tangible code with CLEAN* architecture and no fights with frameworks.
  
<h4> What is idi?</h4>

- The name "idi" is from the word "idiomatic". I prefered a short name for CLI since it requires me to type it often.


<h4>Why Would I use this?</h4>

- Well you don't have to, since it is highly opinionated and might be a overkill for some projects.
- For zero or less opinionated scaffolders checkout the project below:
      - [autostrada](autostrada.dev) 
      - [go-blueprint](https://github.com/Melkeydev/go-blueprint) 
  

<h4>Table of Contents</h4>

- [Install](#install)
- [Database](#database-support)
- [Usage Example](#usage-example)
- [License](#license)

<a id="install"></a>
<h4>
  Install
</h4>

```bash
go install github.com/pammalPrasanna/idi@v0.1.6
```

This installs a go binary that will automatically bind to your $GOPATH

> if you're using Zsh, you’ll need to add it manually to `~/.zshrc`.

```bash
GOPATH=$HOME/go  PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

don't forget to update

```bash
source ~/.zshrc
```

Then in a new terminal run:

```bash
idi -cp <project_name> -cdb <db_name> -auth -paseto
```

See `idi -h` for all the flags available.

-  -cp : create project with flag: idi -cp [project name]
-  -ca : add one or more apps with flag: idi -ca [appname1,appname2]
-  -cdb : add db with flag: idi -cdb [mysql/postgres/sqlite3]
-  -auth : add JWT authentication with flag: idi -cp -auth
-  -paseto : add Paseto instead of JWT with flag: idi -cp -auth -paseto
-  -cr : add router with flag: idi -cr [chi/httprouter/mux] (currently 'httprouter' only)
-  -v :   display version and exit

<a id="database-support"></a>

<h4>
  Database 
</h4>

- Use the `-cdb` create db flag to specify the database driver me want to integrate into mer project.
- Currently below three SQL database drivers and helpers can be scaffolded.

  - [mysql](https://github.com/go-sql-driver/mysql)
  - [postgres](https://github.com/jackc/pgx/)
  - [sqlite3](https://github.com/mattn/go-sqlite3)


<a id="usage-example"></a>
<h4>
  Usage Example: 
</h4>

To create a project with single app, JWT, sqlite3 (later you can add more apps)
```bash
idi -cp projectname -ca app1 -cdb sqlite3 -auth
```
To create a project with multiple apps, Paseto, postgres (to create all the apps at the same time)
```bash
idi -cp projectname -ca app1,app2 -cdb postgres -auth -paseto
```

```* my perception may not reflect reality```


<a id="license"></a>
<h4>
  Licence
</h4>

This project is licensed under [GNU AGPL License](./LICENSE)
Scaffolded / Generated code licensed under [MIT License](https://opensource.org/license/mit)