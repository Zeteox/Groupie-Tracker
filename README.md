# Groupie-Tracker

## Description

Groupie-Tracker is a project that allows users to track music bands,
their albums, and concert locations.
#### Feature:
You can search for specific groups or members and see there concerts locations.


## Technologies

- Go
- HTML
- CSS
- JavaScript
- Leaflet (JavaScript library for interactive maps)
## Installation

1. Clone the repository:
    ```sh
    git clone <repository-url>
    ```
2. Navigate to the project directory:
    ```sh
    cd project-root
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```
2. Open your browser and navigate to `http://localhost:8080`.

## Project Structure

```
./
├── Cmd/
│   └── main.go
├── Frontend/
│   ├── Static/
│   │   ├── CSS/
│   │   │   ├── discoverpage.css
│   │   │   ├── index.css
│   │   │   └── infos.css
│   │   ├── Images/
│   │   │   ├── fond_vinyle.jpg
│   │   │   ├── vinyle_banner1.jpg
│   │   │   ├── favicon.ico
│   │   │   ├── Groupify_blanc.png
│   │   │   └── ...
│   │   └── JS/
│   │       ├── Map.js
│   │       └── discoverpage.js
│   └── Templates/
│       ├── discoverpage.gohtml
│       ├── FilterModel.gohtml
│       ├── index.gohtml
│       └── infos.gohtml
├── Internal/
│   └── Server/
│       ├── Handlers/
│       │   ├── HandlerDiscoverPage.go
│       │   ├── HandlerIndexPage.go
│       │   └── HandlerInfosPage.go
│       └── Server.go
├── Pkg/
│   ├── DataStruct/
│   │   └── GroupsDataStuct.go
│   ├── Utils/
│   │   ├── APIUtils.go
│   │   └── CommonUtils.go
├── go.mod
└── README.md
```

## Links

- [Trello](https://trello.com/invite/b/676028e853805b28510b469e/ATTI867b50d0626ea39827b69f17960fb89689CBEE8C/groupie-tracker)
- [Presentation](https://www.canva.com/design/DAGgTA-xgN4/YmjO0ELwb5h5LH69J_eEVA/edit?utm_content=DAGgTA-xgN4&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)
- [Git Repository](https://ytrack.learn.ynov.com/git/dloic/Groupie-Tracker)

## Credits

<a href="https://ytrack.learn.ynov.com/git/relisabe"><img src="https://ytrack.learn.ynov.com/git/avatars/a83fd07a806d19baa92daf4d338ea900?size=870" width="75"/></a>
<a href="https://ytrack.learn.ynov.com/git/mmelida"><img src="https://ytrack.learn.ynov.com/git/avatars/df790307056382f51371398504bc993b?size=870" width="75"/></a>
<a href="https://ytrack.learn.ynov.com/git/dloic"><img src="https://ytrack.learn.ynov.com/git/avatars/84e0a0c58261edac57bc7321d4bf0466?size=870" width="75"/></a>
- [Elizabeth Robl](https://ytrack.learn.ynov.com/git/relisabe)
- [Melisa Mckenzie](https://ytrack.learn.ynov.com/git/mmelisa)
- [Loïc Delprat](https://ytrack.learn.ynov.com/git/dloic)