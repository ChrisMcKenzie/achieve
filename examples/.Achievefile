provider "github" {
  token = "github_token"
}

task "build" {
  action "script_run" {
    content = "go build -o do main.go"
  }
}

task "deploy" {
  action "github_create_release" {
    repo = "chrismckenzie/minesweeper"
    version = "v0.0.2-pre"
    pre_release = false
    target_commitish = "master"
  }
}

task "run" {
  action "watch" {
    script = "go run main.go"
  }
}
