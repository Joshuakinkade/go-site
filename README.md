# README

## Goals

## Project Design
### Architecture
- basically a web MVC with different names.
- a web layer that handles requests and responses.
- a service layer that is user interface and storage agnostic. It contains business logic and provides methods for working with data.
- a data layer that handles database access.
#### Web Server
- using fiber for routing, wanted to build this myself, but didn't want to spend the time on it right now.
  - have used another one at work, but it's no longer maintained.
  - fiber is easy to use.
#### Services
- business logic
#### Data Layer/SQL
- existing options in go
  - sqlc
    - like that it is based on hand written queries.
    - doesn't allow dynamic queries, which is good until I need one.
    - not sure it really adds much value, especially with AI code completion.
  - gorm
    - makes it harder to optimize and check queries.
    - too opinionated
  - DIY
    - wanted the learning experience
    - hand written queries
    - scan results into application datatypes directly
    - more custom code involved, but code completion can help with that.
    - I can find things that are easy to factor into reusable functions.
    - eventually, I might just end up with my own library in the end anyway.
### Testing
  - have used go's builtin testing at work, but it can be hard to read test cases and understand what's being tested.
  - have used bdd frameworks like Jest in javascript and really liked that.
  - using ginkgo and gomega in this project and am very happy so far.

## Todo
### Photos
- [ ] find a good go image manipulation library
- [ ] implement image optimization
- [ ] set up localstack s3 container

### Posts
- [ ] render post body
  - [ ] typography
  - [ ] photo gallery markdown tag
- [ ] post header
  - title
  - metadata: date, tags, reading time, etc.

## Projects Page
- [ ] database model & repository
- [ ] service
- [ ] handler
- [ ] templage

### Other
- [ ] about me page
- [ ] site footer with contact info
- [ ] site header with navigation
- [ ] hot reloading templates in dev mode

## Project Structure

`handlers` - process requests and return responses
`models` - define the data models
`services` - provide business logic
`db` - database access

## Routes

Using [fiber](https://docs.gofiber.io/) to handle routing. Go's builtin routing is too limited and I don't want to take time to build my own right now.

### Frontend

These are the routes for returning pages.

- `GET /` - display the home page
- `GET /posts` - display a paginated list of posts
- `GET /posts/<slug>` - display a post

### API

These are the routes for managing site data.

- `POST /api/posts` - create a new post
- `PUT /api/posts/<slug>` - replace a post

### Images

This is the route for loading optimized photos.

- `GET /photos/id` - load a photo with optional optimizations.

## Rendering Pages

Pages will be stored in the database as markdown text, with metadata stored in other fields. The pages will be rendered once and cached for a time.

I'll need a custom markdown tag for marking gallery photos

## Page Design

- use tailwind for css as a good starting point
- keep page design super simple, since I'm not a designer. Use default page flow as much as possible
- for frontend interactivity, use vanilla js to avoid big third party libraries/frameworks.

### Gallery Desgin

I want to be able to embed photo galleries into blog posts. I'll add a `data-gallery` and `data-photo-id` attributes to gallery photos so a gallery script can find them and load them into the fviewer.

I might want to consider 3rd party javascript for handling touch events.
