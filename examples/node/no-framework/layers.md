# Layers

## N-Layers

- src: All source code

  - entities: Object mappings
  - factories: Instance generators
  - repositories: Data access
  - routes: Endpoints mappings
  - services: Communication between the routes and repositories (business logic)
  - util: Shared codes
  - handler: Communication between routes and server
  - index: Server instance

- tests: All automated test suites
  - integration tests: Testing on the user point of view. It is also an E2E test because there is no app consuming it.
  - unit tests: All tests that must run without any external connections such as database, external APIs and on our case, the fileSystem.
