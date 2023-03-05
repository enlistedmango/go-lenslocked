# Code Organisation

- Some codebases are easy to navigate
- Others are hard to follow
- Code organisation is often the difference

A good structure will:

- Make it easier to guess where bugs are
- Make it easier to add new features
- And more

## Flat Structure

All code is in a single package

Files can still separate code:

```
myapp/
    gallery_store.go
    gallery_handler.go
    gallery_templates.go
    user_store.go
    user_handler.go
    user_templates.go
    router.go
    ...

```

## Seperation of concerns

Separate code based on duties.

HTML & CSS is a traditional example:

- HTML focuses on overall structure
- CSS focuses on styling it

Model-View-Controller (MVC) is a popular structure following this strategy. In a web app:

- `models` => data, logic, rules; usually database
- `views` => rednering things; usually html
- `controller` => connects it all. Accepts user input, passes that to the models to do stuff, then passes data to views to render things; usually handlers

```bash
myapp/
    controllers/
        user_handler.go
        gallery_handler.go
        ...
    views/
        user_templates.go
        ...
    models/
        user_store.go
        ...
```

Doesn't need to be named `models` `views` and `controllers`.
gobuffalo.io/en uses a variation of MVC, but doesn't name directories exactly.

Very common in framework b/c it is predictable. Makes it easier for a framework to guess were certain files will be.

Ruby on Railes, Django, etc use MVC

## Dependency-based structure

Ben Jognson writes about this at [Go Beyond](https://gobeyond.dev/standard-package-layout)

Structured based on dependencies, but with a common set of interfaces and types they all work with.

```bash
myapp/
    user.go # defines a User type that isn't locked into any specific data store
    user_store.go # defines a UserStore interface

    psql/
        user_store.go # implements the UserStore interface with a Postgresql specific implementation
```

## Many More...

- Domain Driven Design
- Onion architecture
- ...

## Which should you use?

No such thing as one-size fits all.
Each has pros & cons.

Mat Ryer and David Hernandez used flat at <pace.dev> early on

Flat can work well with microservices since each is small
Flat won't work well at Google with a massive monorepo

Experience also matters:

- Globals are bad, but why?
- Testable code only makes sense after you try to test untestable code.

The goal is to make this course accessible!

_Remember: You can always refactor._

Don't get decision paralysis.
Most people spend way longer deciding on structure, especially beginners to the point that it discourages them from doing anything.

## We're are going to use MVC

It isn't perfect, but it:

- Is predictable
- Will provide us with some clear rules to follow when organising out code
- It will prevent us from getting stuch trying to decide how to organise our code
- Is easy to pickup and utilise when learning. Many others are not and are better suited to experienced developers and teams.

Later we can explore how we might refactor to something different from MVC

_when we understand out needs better_
