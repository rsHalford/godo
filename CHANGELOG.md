# CHANGELOG

<a name="0.16.0"></a>
## [0.16.0](https://github.com/rsHalford/godo/compare/v0.15.0...0.16.0)

> 2022-06-26

### Chores

* **docs:** added go report card badge to project README
* **docs:** improve the package header documentation
* **version:** update for v0.16.0

### Code Refactoring

* improved nomenclature with consistency throughout
* **version:** replaces versionCmd with the rootCmd Version flag

### Features

* **config:** changed to more appropriate config filetype TOML

### BREAKING CHANGE


Todo Status is now referred to as Done.


<a name="v0.15.0"></a>
## [v0.15.0](https://github.com/rsHalford/godo/compare/v0.14.0...v0.15.0)

> 2022-06-23

### Chores

* **version:** update for v0.15.0

### Features

* **list:** ability to only list items if they're a priority
* **theme:** allow users to customise the colour of output


<a name="v0.14.0"></a>
## [v0.14.0](https://github.com/rsHalford/godo/compare/v0.13.1...v0.14.0)

> 2022-06-21

### Chores

* **docs:** improved visibility to configure through environment variables
* **version:** update for v0.14.0

### Code Refactoring

* **config:** structs broken up to ease future additions

### Features

* **color:** support NO_COLOR environment variable


<a name="v0.13.1"></a>
## [v0.13.1](https://github.com/rsHalford/godo/compare/v0.13.0...v0.13.1)

> 2022-06-18

### Chores

* direnv development environment setup and documentation
* **deps:** updated and cleaned dependencies
* **docs:** updated README to include find configuration and project goals
* **gitignore:** added vendor files to .gitignore
* **version:** update for v0.13.1

### Code Refactoring

* **colour:** added go-colour-util dependency
* **find:** remove redundant break


<a name="v0.13.0"></a>
## [v0.13.0](https://github.com/rsHalford/godo/compare/v0.12.2...v0.13.0)

> 2022-05-17

### Chores

* **version:** update for v0.13.0

### Features

* **find:** case-sensitivity flag argument to filter search results

### Styling

* **config:** matched formatting of Config variables


<a name="v0.12.2"></a>
## [v0.12.2](https://github.com/rsHalford/godo/compare/v0.12.1...v0.12.2)

> 2022-05-13

### Chores

* **docs:** included installation as a nix flake
* **flake:** packaged for installation via nix flake input
* **version:** update for v0.12.2

### Performance Improvements

* **json:** parse godos.json with decoder

### Styling

* further reduced output of help description header


<a name="v0.12.1"></a>
## [v0.12.1](https://github.com/rsHalford/godo/compare/v0.12.0...v0.12.1)

> 2022-04-18

### Bug Fixes

* **config:** corrected comment describing default dataFile path

### Styling

* reduced output of help description header


<a name="v0.12.0"></a>
## [v0.12.0](https://github.com/rsHalford/godo/compare/v0.10.0...v0.12.0)

> 2022-04-18

### Bug Fixes

* **remove:** assigns var title before removal
* **version:** update for v0.12.0

### Chores

* **docs:** remove mention of remote godo-api
* **docs:** remove mention of GUI
* **docs:** correct typos for commands and config
* **docs:** updated README with latest commands and contribution details

### Features

* createdAt and updatedAt tags are assigned to todo items
* **gui:** removed GUI implementation
* **remote:** remove remote api request implementation


<a name="v0.10.0"></a>
## [v0.10.0](https://github.com/rsHalford/godo/compare/v0.9.2...v0.10.0)

> 2022-04-08

### Bug Fixes

* removed uneeded config file group env tags

### Chores

* **docs:** updated changelog for v0.10.0
* **docs:** 🎁 added CONTRIBUTING form
* **docs:** 📜 added CI section for CONTRIBUTING
* **docs:** 📜 setup CHANGELOG generator

### Features

* **config:** create config.yaml if it doesn't exist


<a name="v0.9.2"></a>
## [v0.9.2](https://github.com/rsHalford/godo/compare/v0.9.1...v0.9.2)

> 2021-11-18

### Styling

* ⬅️ more compact output


<a name="v0.9.1"></a>
## [v0.9.1](https://github.com/rsHalford/godo/compare/v0.9.0...v0.9.1)

> 2021-11-17

### Styling

* swap from escaped quotations around %v to %q


<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/rsHalford/godo/compare/v0.8.5...v0.9.0)

> 2021-11-17

### Features

* 🔖 add tags to todos


<a name="v0.8.5"></a>
## [v0.8.5](https://github.com/rsHalford/godo/compare/v0.8.4...v0.8.5)

> 2021-11-16


<a name="v0.8.4"></a>
## [v0.8.4](https://github.com/rsHalford/godo/compare/v0.8.3...v0.8.4)

> 2021-10-27


<a name="v0.8.3"></a>
## [v0.8.3](https://github.com/rsHalford/godo/compare/v0.8.1...v0.8.3)

> 2021-10-25


<a name="v0.8.1"></a>
## [v0.8.1](https://github.com/rsHalford/godo/compare/v0.8.0...v0.8.1)

> 2021-10-19


<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/rsHalford/godo/compare/v0.7.0...v0.8.0)

> 2021-10-18


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/rsHalford/godo/compare/v0.6.1...v0.7.0)

> 2021-10-17


<a name="v0.6.1"></a>
## [v0.6.1](https://github.com/rsHalford/godo/compare/v0.6.0...v0.6.1)

> 2021-08-21


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/rsHalford/godo/compare/v0.5.1...v0.6.0)

> 2021-08-21


<a name="v0.5.1"></a>
## [v0.5.1](https://github.com/rsHalford/godo/compare/v0.5.0...v0.5.1)

> 2021-08-21


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/rsHalford/godo/compare/v0.4.1...v0.5.0)

> 2021-08-18


<a name="v0.4.1"></a>
## [v0.4.1](https://github.com/rsHalford/godo/compare/v0.4.0...v0.4.1)

> 2021-07-29


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/rsHalford/godo/compare/v0.3.2...v0.4.0)

> 2021-07-28


<a name="v0.3.2"></a>
## [v0.3.2](https://github.com/rsHalford/godo/compare/v0.3.1...v0.3.2)

> 2021-07-22


<a name="v0.3.1"></a>
## [v0.3.1](https://github.com/rsHalford/godo/compare/v0.3.0...v0.3.1)

> 2021-06-14


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/rsHalford/godo/compare/v0.2.5...v0.3.0)

> 2021-06-12


<a name="v0.2.5"></a>
## [v0.2.5](https://github.com/rsHalford/godo/compare/v0.2.4...v0.2.5)

> 2021-06-11


<a name="v0.2.4"></a>
## [v0.2.4](https://github.com/rsHalford/godo/compare/v0.2.3...v0.2.4)

> 2021-06-11


<a name="v0.2.3"></a>
## [v0.2.3](https://github.com/rsHalford/godo/compare/v0.2.2...v0.2.3)

> 2021-05-26


<a name="v0.2.2"></a>
## [v0.2.2](https://github.com/rsHalford/godo/compare/v0.2.1...v0.2.2)

> 2021-05-26


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/rsHalford/godo/compare/v0.2.0...v0.2.1)

> 2021-05-26


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/rsHalford/godo/compare/v0.1.1...v0.2.0)

> 2021-05-21


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/rsHalford/godo/compare/v0.1.0...v0.1.1)

> 2021-05-17


<a name="v0.1.0"></a>
## [v0.1.0](https://github.com/rsHalford/godo/compare/v0.0.4...v0.1.0)

> 2021-05-16


<a name="v0.0.4"></a>
## [v0.0.4](https://github.com/rsHalford/godo/compare/v0.0.3...v0.0.4)

> 2021-05-07


<a name="v0.0.3"></a>
## [v0.0.3](https://github.com/rsHalford/godo/compare/v0.0.2...v0.0.3)

> 2021-05-07


<a name="v0.0.2"></a>
## [v0.0.2](https://github.com/rsHalford/godo/compare/v0.0.1...v0.0.2)

> 2021-05-07


<a name="v0.0.1"></a>
## v0.0.1

> 2021-05-06

