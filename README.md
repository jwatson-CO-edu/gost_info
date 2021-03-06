# gost_info
[Go] [S]cripting [T]utorials and [Info]

# Dev Plan
```
[Y] Take "A Tour of Go": https://tour.golang.org/list - 2020-12-31
[ ] Concurrent TCP Server
    * https://opensource.com/article/18/5/building-concurrent-tcp-server-go
    * https://github.com/mactsouk/opensource.com/blob/master/concTCP.go
    * Additional Ref: 
    https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/08.1.html
[ ] UDP Server/Client: https://ops.tips/blog/udp-client-and-server-in-go/
[ ] Review Best Practices: https://golang.org/doc/code.html
    * https://dave.cheney.net/practical-go
[ ] Review Package Structure: 
    https://github.com/golang/go/wiki/GithubCodeLayout
[ ] Concurrency & Parallelism Deep Dive
    * Communicating Sequential Processes: Bucket Brigade
        - [ ] Goroutines + Channels
        - [ ] Select
    { } https://medium.com/@tilaklodha/concurrency-and-parallelism-in-golang-5333e9a4ba64
    { } https://hub.packtpub.com/concurrency-and-parallelism-in-golang-tutorial/
    { } https://spiralscout.com/blog/understanding-concurrency-and-parallelism-in-golang
    { } https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca
    { } https://stackoverflow.com/questions/25106526/parallel-processing-in-golang/44403016#44403016
    { } https://yourbasic.org/golang/efficient-parallel-computation/
[ ] ROM Scraper
[ ] Struct Composition: https://www.geeksforgeeks.org/composition-in-golang/
    { } https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go
    { } Compose Interfaces: https://medium.com/@_jesus_rafael/composing-interfaces-in-go-58980969e897
    { } https://journal.highlandsolutions.com/how-to-compose-an-interface-in-go-beginners-guide-26a56672f0a0
    { } Entity Component System
        [ ] Move this idea to a game project, OGL project takes precedence over a golang game
        * Interface test to see if the component is present
        * Vector of component interface structs
[ ] Review Patterns: https://refactoring.guru/design-patterns/go
```

# Resources
* Massive Package/Resource Directory: 
  https://github.com/avelino/awesome-go

# Possible Projects
```
[ ] Web Server:
    * https://golang.org/doc/articles/wiki/
    * https://medium.com/google-cloud/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud-part-1-building-a-simple-go-aee452a2e654
    * Buffalo?: https://www.gopherguides.com/articles/converting-a-static-website-to-golang-buffalo/
    * https://medium.com/better-programming/the-beginners-guide-to-setting-up-a-web-server-in-golang-9c9473cbe372
    * https://dev.to/coreyvan/from-zero-to-http-servers-with-go-and-raspberry-pi-3oi1
    * https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
    * https://aws.amazon.com/getting-started/hands-on/host-static-website/
[ ] Google API Experiments
    * https://godoc.org/google.golang.org/api
    * https://developers.google.com/youtube/v3/quickstart/go
    * https://developers.google.com/youtube/v3/code_samples/go
[ ] Task Automation / RPi Robotics
    * https://stackoverflow.com/questions/19549199/golang-implementing-a-cron-executing-tasks-at-a-specific-time
    * https://github.com/stianeikeland/go-rpio
    * https://github.com/markdaws/gohome
    * https://gobot.io/documentation/platforms/raspi/
[ ] Go Shell
    * https://golang.org/pkg/os/signal/
    * https://gobyexample.com/signals
    * https://nathanleclaire.com/blog/2014/08/24/handling-ctrl-c-interrupt-signal-in-golang-programs/
[ ] Simulation / Games
    * https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-1-5e2c11a513ed
    * https://github.com/gen2brain/raylib-go
    * https://www.reddit.com/r/golang/comments/ai1ob5/creating_ecs/
    * https://github.com/EngoEngine/ecs
[ ] Simplest APK
    * https://pkg.go.dev/golang.org/x/mobile/app
    * https://github.com/golang/go/wiki/Mobile
    * https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile
    * https://github.com/xlab/android-go

```
