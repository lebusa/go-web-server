# Go Web Server
A basic go web server that serves static web content over https.

## Features
The server:
- handles static content from a directory (i.e. static HTML content) via an http module's Handle 
- handles basic response via an http module's handleFunc
- uses a self-signed TLS certificate
- uses a mutex to avoid race conditions for requests to one of its routes (/counter)

