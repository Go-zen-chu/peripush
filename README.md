# peripush

When you want to send metrics to prometheus, it's a bit bothersome to configure scrape config every time.

peripush is very simple project to gather metrics from localhosted server and push it to push gateway, that's it.

## Feature
- gather metrics, send it to push gateway  
- single binary

## Development

### architecture

Layered architecture is adopted to this system.

### project layout

Project layout is under [golang-standards/project-layout: Standard Go Project Layout](https://github.com/golang-standards/project-layout)