# README

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
