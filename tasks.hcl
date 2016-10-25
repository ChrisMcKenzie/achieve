provider "github" {
  token = "github_token"
}

task "default" {
  action "script_run" {
    content = "go install"
  }
}

task "build" {
  action "script_run" {
    content = "go build -o achieve main.go"
  }
}

task "release" {
  action "github_create_release" {
    repo = "chrismckenzie/achieve"
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
