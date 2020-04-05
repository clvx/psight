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

## Working with HTTP Requests 
- Query parameters: 

    https://localhost:8000/search?q=products&page=1

    func (w http.ResponseWriter, r *http.Request) {
        url := r.URL                // net/url.URL - access to the url itself
        query := url.Query()        // net/url.Values (map[string][]string)
                                    // map of string of slice of strings,
                                    // a key parameter could appear multiple times in an url.
        q := query["q"]             // []string{"producs"}
        page := query.Get("page")   // "1" , returns the first instance object of that key.
    }

- Form data:

    <form action="." method="post">
        Username: <input type="text" name="username"/>
        Password: <input type="password" name="password"/>
        <button type="submit"</button>
    </form>


    func (w http.ResponseWriter, r *http.Request){
        err := r.ParseForm          // Populate Form and PostForm fields on the request object

        f := r.Form                 // net/url.Values (map[string][]string)
                                    // map of string of slice of strings,
                                    // a key parameter could appear multiple times in an url.
        username := f["username"]   // []string{"Michael"}

        pass := f.Get("password") // "password", returns the first instance object of that key.
    }

    - Parsing form data: Letting the request object know you are recieving a form.
        - ParseForm
        - ParseMultipleForm
    - Reading form data 
        - Form, populate itself wherever it can get the data(request body or url parameters).
        - PostForm, populate itself using request body data but it won't look into
        the url parameters.
        - FormValues(), returns the first value for the named component of the query. 
        POST and PUT body parameters take precedence over URL query string values.
        - PostFormValue(), returns the first value for the named component of the 
        POST, PATCH, or PUT request body. URL query parameters are ignored.
        - FormFile(), returns the first file for the provided form key.
        - MultipartReader(),  returns a MIME multipart reader if this is a multipart/form-data 
        or a multipart/mixed POST request, else returns nil and an error. Use this 
        function instead of ParseMultipartForm to process the request body as a stream. 

- Working with JSON

    { 
        "term":  "products",
        "page": 1,
        "pageSize": 25
    }

    //Using tags to map struct field names to json keys
    type Query struct {
        Term        string  `json:"term"`
        Page        int     `json:"page"`
        PageSize    int     `json:"pageSize"`
    }

    - Reading JSON

    func (w http.ResponseWriter, r *http.Request) {
        dec := json.NewDecoder(r.Body)      // NewDecoder returns a new decoder that reads from r. 
                                            //r.Body implements the reader interface.
        var query Query
        err := dec.Decode(&query)           //stores data after it's being decoded in query.
                                            //Decode reads the next JSON-encoded 
                                            //value from its input and stores it 
                                            //in the value pointed to by v. 
    }

    - Writing JSON

    type Result struct {
        ID          int     `json:"id" 
        Name        string  `json:"name" 
        Description string  `json:"string" 
    }

    func (w http.ResponseWriter, r *http.Request) {
        var results []Result = model.GetResults()
        enc := json.NewEnconder(w) //w implements io.Writer
        err := enc.Encode(results) //Encode writes the JSON encoding of v to the 
                                   // stream, followed by a newline character. 
    }

## Middleware

A request would be processed before the request handler. In the same way, a response
would be processed after the request handler.

- Creating a middleware: To implement a middleware we must implement the handler
interface. 

    //if handler is nil, it will use http.DefaultServerMux
    http.ListenAndServer(addr string, handler Handler) error

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

    type Middleware struct {
        Next http.Handler
    }

    func (m MyMiddleware) ServeHTTP(w htt.ResponseWriter, r *http.Request) {
        // do things before next handler(managing request)

        m.Next.ServerHTTP(w, r) //forwards request to whichever next handler is.
                                //It could be another middleware or http.DefaultServerMux

        // do things after next handler(managing response)
    }

- Common uses:
    - Logging.
    - Security.
    - Request timeouts.
    - Response compression.

## Context
A way to think about context package in go is that it allows you to pass in a “context” 
to your program. Context like a timeout or deadline or a channel to indicate stop 
working and return.

- Using request contexts

    //Returns the current context the request is operating whitin
    //Context() does not allow to manipulate the context, only interrogate it
    func (*Request) Context() context.Context

    //WithContext() allows you to modify the context
    func (*Request) WithContext(ctx context.Context) context.Context

- Context API

    type Context interface {
        //Deadline() lets you know when the context becomes invalid
        Deadline() (deadline time.Time, ok bool)

        //Gets a signal on a channel when the context is aborted(like receiving a singnal)
        Done() <-chan struct{}

        //Error when context is aborted
        Err() error

        //Pull information from the context to pass them between layers
        Value(key interface{}) interface{}
    }

- Modify Context

    //Return a new context with a Context.Cancel() which allows sending a signal 
    // to the Done() channel to cancel the context
    WithCancel()

    //Like Context.Cancel() but it adds a timestamp to the Done() channel
    WithDeadline()

    //Like WithDeadline() but instead of a time it adds a duration to the Done() channel
    WithTimeOut()

    //Create a new context which adds a new value onto the existing context
    WithValue()
