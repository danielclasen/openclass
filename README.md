[![Build Status](https://travis-ci.org/danielclasen/openclass.svg?branch=master)](https://travis-ci.org/danielclasen/openclass)

# Stack

Backend:
- go 1.11
- gin
- godoc (for api tests)

Frontend:
- nodejs v10.2.1
- Angular 6
- Angular Material

## Run openclass


**Make sure you have go and nodejs installed in the corresponding versions.**

To start the backend just run the main.go file:

    go run main.go
    
    
Fetch the UI dependencies with npm:

    npm ci

Then start a concurrent build of the UI resources:
    
    ng build --watch

Or if you don't have ng installed globally, go ahead with the oneshot build using npm directly:

    npm run build
    
After that you should be able to see openclass in action at *http://localhost:3000*