<p align="center">
    <img src="https://raw.githubusercontent.com/System-Glitch/goyave/develop/resources/img/logo/goyave_text.png" alt="Goyave Logo" width="550"/>
</p>

<p align="center">
    <a href="https://travis-ci.com/System-Glitch/goyave"><img src="https://api.travis-ci.com/System-Glitch/goyave.svg" alt="Build Status"/></a>
    <a href="https://github.com/System-Glitch/goyave/releases"><img src="https://img.shields.io/github/v/release/System-Glitch/goyave?include_prereleases" alt="Version"/></a>
    <a href="https://goreportcard.com/report/github.com/System-Glitch/goyave"><img src="https://goreportcard.com/badge/github.com/System-Glitch/goyave" alt="Go Report"/></a>
    <a href="https://coveralls.io/github/System-Glitch/goyave?branch=master"><img src="https://coveralls.io/repos/github/System-Glitch/goyave/badge.svg" alt="Coverage Status"/></a>
    <a href="https://github.com/System-Glitch/goyave/blob/master/LICENSE"><img src="https://img.shields.io/dub/l/vibe-d.svg" alt="License"/></a>
</p>

<h2 align="center">An Elegant Golang Web Framework</h2>

Goyave is a progressive and accessible web application framework, aimed at making development easy and enjoyable. It has a philosophy of cleanliness and conciseness to make programs more elegant, easier to maintain and more focused.

<table>
    <tr>
        <td valign="top">
            <h3>Clean Code</h3>
            <p>Goyave has an expressive, elegant syntax, a robust structure and conventions. Minimalist calls and reduced redundancy are among the Goyave's core principles.</p>
        </td>
        <td valign="top">
            <h3>Fast Development</h3>
            <p>Develop faster and concentrate on the business logic of your application thanks to the many helpers and built-in functions.</p>
        </td>
        <td valign="top">
            <h3>Powerful functionalities</h3>
            <p>Goyave is accessible, yet powerful. The framework includes routing, request parsing, validation, localization, and more!</p>
        </td>
    </tr>
</table>

Most golang frameworks for web development don't have a strong directory structure nor conventions to make applications have a uniform architecture and limit redundancy. This makes it difficult to work with them on different projects. In companies, having a well-defined and documented architecture helps new developers integrate projects faster, and reduces the time needed for maintaining them. For open source projects, it helps newcomers understanding the project and makes it easier to contribute.

## Getting Started

### Install using the template project

You can bootstrap your project using the [Goyave template project](https://github.com/System-Glitch/goyave-template). This project has a complete directory structure already set up for you.

#### Linux / MacOS

```
$ curl https://raw.githubusercontent.com/System-Glitch/goyave/master/install.sh | bash -s my-project
```

#### Windows (Powershell)

```
> & ([scriptblock]::Create((curl "https://raw.githubusercontent.com/System-Glitch/goyave/master/install.ps1").Content)) -projectName my-project
```

---

Run `go run my-project` in your project's directory to start the server, then try to request the `hello` route.
```
$ curl -H 'Accept-Language: pt-BR;q=0.9, en;q=0.8' -H "Content-Type: application/json" http://localhost:3000/user -d '{"name": "teste", "email": "teste@teste.com", "phone": "+5511999999999", "uid": "f8dbea72-1a2f-436c-ac4b-3b5665f4b8da"}'
Hi!
```


## Learning Goyave

The Goyave framework has an extensive documentation covering in-depth subjects and teaching you how to run a project using Goyave from setup to deployment.

<a href="https://system-glitch.github.io/goyave/guide/installation"><h3 align="center">Read the documentation</h3></a>

## Requirements

- Go 1.13+
- Go modules
