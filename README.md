# billable
A golang web application, accepts an csv file, processes it and sends back a result as json

How To : 
1. Ensure you have Golang installed and your GOPATH / GOROOT are set correctly on your computer (see https://golang.org/doc/install for installations)
2. clone repository
3. To get a quick feel and flow of the application, run the executables provided in the build directory, the golang web server will start and listen for connections on port : 8100. Go to localhost:8100/ to see web api documentation
4. To continue with the project build process, open the cloned repo on the command line and run "go get" this will get all project dependencies, run "go test ./.." to run test on the multiple project packages, then run "./build/go-build-linux.sh" for linux os || "./build/go-build-mac.sh" for mac os || "./build/go-build-win.sh" for windows os, to build the project, this will output an executable "billable.elf||billable.app||billable.exec" at the root of the project folder, double click on the executable to start the application
5.  The golang web server will start and listen for connections on port : 8100. Go to localhost:8100/ to see web api documentation
