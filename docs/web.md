# Web

## Connection Handler
- Receive a TCP request and response the TCP request.
- Listen for requests for a well known port.
- To handle requests go provides `http.DefaultServerMux` which listens in port `62121`
which triggers goroutines for each request.

## Basic Request Handler
- Handler: Ability to register custom logic to respond to the different requests 
that come into the server.
- `http.Handle` 
- `http.HandleFunc`

## Built In Handlers
- NotFoundHandler: returns a *404 NotFound*
    func NotFoundHandler() Handler
- RedirectHandler: redirects to another url
    //url to redirect 
    //redirect status code
    func RedirectHandler(url string, code int) Handler
- StripPrefix
    //prefix to remove from incoming url
    //h handles request after prefix has been removed
    func StripPrefix(prefix string, h handler) Handler
- TimeoutHandler
    //h function to handle the request
    //dt amount of time that the handler(h) is allow to process
    //msg message for the response
    func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
- FileServer
    func FileServer(root FileSystem) Handler

    type FileSystem interface{
        Open(name string) (File, error)
    }

    type Dir string
    func (d Dir) Open(name string) (File, error)

## Templating

- Template: Combine and organize data on the view layer using template logic.
- Basic functions:
    - New(), generates template.
    - Parse(), create templates.
    - Execute(), compiles template with data.
- Loading templates:
    - ParseFiles(), parsed a list of files as its own template.
    - ParseGlob(), parsed a a list of files which follow a pattern filename.
    - Lookup(), finds template by its name after being loaded.
    - Must(), takes a template and an error, and panics if template does not exist.
- Subtemplates:
    - Removes redundancy in html pages
    - `{{ define "name" }} .. {{end}}`, defines a template section
    - `{{template "name"}}`, injects templates in a page.
    - Last template name defined is rendered.
- Template Composition:
    - Allows to create pages with required and/or optional templates using:
        - `{{ block "name" .}}{{end}}, optional
        - `{{template "content" .}}, required
- Data Driver templates:
    - Consume a data context in order to dynamically generate a document instead of generating one statically.
    type Context struct { 
        Title string
        ImageURL string
    }
    context := Context{Title: "title", ImageURL: "http://..."} 
    t := template.Must(template.New("").Parse("..."))
    t.Execute(w, context)

    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        <img src="{{.ImageURL}}"/>
    </body>

## Smart Templates

- Pipelines:
    - `{{command1 command2 command3}}`
    - command types: 
        - literal
        - function name: `{{template "content"}}`
        - data field: `{{.Title}}`
        - methods: `{{.SayMsg "Hello World"}}`
    - Functions and methods must return one value, or a value and an error.
    - `{{ command 1 command 2 | command 3}}`, pass result of previous command as last argument of next command.
- Built in functions:
    - `define`: define a subtemplate inside a template.
    - `template`: inject a template into the current template
    - `block`: defines a default template and injects it to the current template.
    - `html`, `js`, `urlquery`: escapes incoming data.
    - `index(collection, index)`: drill into a collection type.
    - `print`, `printf`, `println`: fmt.Sprint, fmt.Sprintf, fmt.Sprintln.
    - `len`: len of a collection.
    - `with`: narrow down the scope of you data context.
- Custom functions:
    template.Funcs(funcMap FuncMap) *Template
    type FuncMap map[string]inteface{}
    template.New("").Funcs(funcMap).Parse(...)
    - Acceptable values for FuncMap:
        - Function that returns a single value.
        - Function that returns a single value and an error type.
- Logical tests:
    {{ if pipeline }}
        T1 //prints if pipeline results is a non-empty value
    {{else if pipeline}}
        T2
    {{end}}
    - empty values:
        - false
        - zero
        - nil
        - empty collection
    - operators: _all arguments are evaluated every single time_.
        - `eq/ne`
        - `lt/gt`
        - `le/ge`
        - `and`
        - `or`
        - `not`
- Looping:
    {{range pipeline}}
    T1 //executes if pipeline is not empty 
    {{else}}
    T2
    {{end}}
    - pipeline must be array, slice, map or channel.
    - Data context of T1 is the current collection item.

## Routing Requests
- Role of a Controller:
    - main(): environment aware.
        - Setup responsability.
        - Initialize templates.

    - controller: 
        - Static resources.
        - Application logic.
- Static routing:
    - Front controller: 
        - Receives requests into the application and routes to the right handler.
        - It routes matching with the most specific pattern specified.
    - App handler: responsible to process requests.

- Parametric Routing:
    - Routes requests to pages based on url parameters.

    
