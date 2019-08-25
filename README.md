# billable
A golang web application, accepts an csv file, processes it and sends back a result as json

A web application that accepts a timesheet (in csv format) as input and automatically generates invoices for each company in the following format.

How To : 
1. Ensure you have Golang installed on your computer (see https://golang.org/doc/install for installations)
2. clone repository
3. To get a quick feel and flow of the application, run the executables provided in the build directory, the golang web server will start and listen for connections on port : 8100. Go to localhost:8100/ to see web api documentation
4. To continue with the project build process, cd to the project directory and run "go get" on command line to get all project dependencies, run "go test" to run test on project, run "go test --cover" to get code coverage details, then run "go build" to build the project, after running go build an executable will be created at the root of the project folder, double click on the executable to start the application
