golang-project-template
=======================

Go Programming Language Project Template

This repository base on standards described in this [project](https://github.com/golang-standards/project-layout/blob/master/README.md).

The project directory tree:

```text
/golang_project
    ├── build
    │   └── application     # Main application executable
    ├── internal            # Internal packages
    │   └── app             # Application-specific code
    ├── pkg                 # Reusable packages
    ├── web                 # Web-related assets
    │   ├── static          # Static files (CSS, JS, images)
    │   └── templates       # HTML templates
    ├── config               # Configuration files
    ├── scripts             # Scripts for automation
    ├── tests               # Test files
    ├── README.md           # Project documentation
    └── main.go             # Main project file
```
