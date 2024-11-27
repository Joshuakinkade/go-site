# README

## Goals
- build a website that allows me to write blog posts and share photos.
- experiment go application development and architecture.
  - handling database stuff: sql queries, migrations, etc.
  - testing methodologies
- get some experience with using AI code completion in go, specifically GitHub Copilot.

## Features
- frontend website with blog posts and photo galleries
- backend API for managing posts and photos
- support for markdown in blog posts
- optimized photo loading

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

### Post List Page
- [ ] intelligently create snippet: don't cut off markdown tags
- [x] paginate post list

### Post Page
- [ ] render post body
  - [x] typography -- just using tailwinds prose plugin for now
  - [ ] photo gallery markdown tag
- [x] post header -- in other project
  - title
  - metadata: date, tags, reading time, etc.

## Projects Page
Save this for later. Work on the blog and photos first.
- [ ] database model & repository
- [ ] service
- [ ] handler
- [ ] templage

### Other
- [ ] about me page - other project
- [x] site footer - other project
- [ ] site header with navigation - other project
- [ ] hot reloading templates in dev mode
- [ ] standardize JSON reponses for API
- [ ] write architecture documentation and thought processes behind decisions
- [ ] move template helpers to a separate file
- [ ] figure out request validation

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

#### Response Format
The API always responds with valid JSON. If the request was successful, the response will contain the data requested in an object with the type and reponse. The http status will be in the 2xx range. If the request resulted in an error, an error message will be returned in the error field. The API will never return a data field if there's an error. The http status will be in the 4xx or 5xx range.

##### Successful Response
```json
{
  "data": {
    type: "postList":
    content: [
      {
        "id": 1,
        "title": "My First Post",
        "slug": "my-first-post",
        "date": "2021-09-01",
        "tags": ["go", "webdev"],
        "readingTime": 5
      }
    ]
  }
}
```

##### Error Response
```json
{
  "error": "Post not found"
}
```

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

## Request Validation

I need a way to easily validate requests. I could use a third party library, I'm trying to avoid that and I think this is a simple enough problem for me to tackle myself. Here are my requirements:
- simple way of defining validation rules. I'd prefer to avoid parsing a micro-syntax in a string.
- need to support a variety of data types each with their own validation rules.
  - strings: min length, max length, regex, reguired vs optional
  - numbers: min, max, required vs optional
  - dates: min, max, required vs optional
- want to avoid doing reflection on go types.
- ideally, return the validated data in the expected type.

### Questions
- What can I reuse from my current validation code?
- Is there a way to make MapValidator return typed fields? Passing in a struct?
- Would using reflection make this easier?
  - Use struct tags on input struct to define validation rules?

### Steps to Validate a Field
1. Check that the field exists in the data.
2. If it is missing, and the field is required, return an error.
3. If the field exists, try to convert it to the expected type.
4. If the conversion fails, return an error.
5. If the conversion succeeds, check the value against the validation rules.
6. As soon as one of the rules fails, return an error.
7. If all of the rules pass, return the type converted value.
8. If the field does not exist, and is optional, return an empty value of the expected type.

```go
type StringChecker struct {
	fieldName string
	required bool
	maxLength int
	minLength int
	pattern *Regexp
}

func (c StringChecker) CheckString(data map[string]interface{}) {

}
```

```go
type CreatePostInput struct {
	Title string `json:"title"; checker:"minLength:5,maxLength:10,regexp:"`
	Body string `json:"body"; checker:""`
}
```
