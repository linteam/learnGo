=== onemli =====
- you would need to make your function exportable with an uppercase for its name



==== go notlari ====
go version
go help     =    godoc.org

==== ENVIRONMENT =========
go env      
    GOPATH="/workspace/go"

--- go (workspace) ---
-bin
-src
    -github.com
        githubUserName
            project1
            project2
-pkg

go mod init oes // projeyi ilklendirmek icin

go mod vendor //importlari indirmek icin .. npm i //qor ile calisirken bunu kullanma
go get https://github.com/nomad-software/vend
vend

go run main.go  //to run 
go fmt fileName.go //yazimi duzeltiyor sekil acisindan

go install    //build icin
